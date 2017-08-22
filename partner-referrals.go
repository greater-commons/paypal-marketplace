package market

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/greater-commons/paypal-marketplace/partner"
)

const createPartnerReferralRoute = "/v1/customer/partner-referrals"

// CreatePartnerReferral is used to connect a user's Paypal account with your platform.
// It is used in both the connected and the managed paths.
func (c *Client) CreatePartnerReferral(ctx context.Context, params *partner.CreatePartnerReferralParams) (*partner.CreatePartnerReferralResponse, error) {
	d, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	r := &request{
		client:   c,
		method:   http.MethodPost,
		endpoint: createPartnerReferralRoute,
		body:     bytes.NewReader(d),
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status == 201 {
		r := &partner.CreatePartnerReferralResponse{}
		err = json.NewDecoder(res.body).Decode(r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	errorData, err := ioutil.ReadAll(res.body)
	if err != nil {
		return nil, err
	}
	return nil, &BadResponse{
		Status: res.status,
		Body:   string(errorData),
	}
}

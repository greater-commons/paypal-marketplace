package market

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/greater-commons/paypal-marketplace/partner"
)

const showAccountTrackngRoute = "/v1/customer/partners/%s/merchant-integrations"

func (c *Client) ShowAccountTracking(ctx context.Context, partnerID, trackingID string) (*partner.GetAccountTrackingResponse, error) {
	r := &request{
		client:   c,
		method:   http.MethodGet,
		endpoint: fmt.Sprintf(showAccountTrackngRoute, url.PathEscape(partnerID)),
	}
	if trackingID != "" {
		r.endpoint += "?tracking_id=" + url.QueryEscape(trackingID)
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status == http.StatusOK {
		r := &partner.GetAccountTrackingResponse{}
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

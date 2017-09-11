package market

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/greater-commons/paypal-marketplace/partner"
)

const showAccountTrackngRoute = "/v1/customer/partners/%s/merchant-integrations"

func (c *Client) ShowAccountTracking(ctx context.Context, partnerID, trackingID string) (*partner.MerchantDetailsData, error) {
	endpoint := fmt.Sprintf(showAccountTrackngRoute, url.PathEscape(partnerID))
	if trackingID != "" {
		endpoint += "?tracking_id=" + url.QueryEscape(trackingID)
	}
	return c.getMerchantData(ctx, endpoint)
}

func (c *Client) ShowMerchantStatus(ctx context.Context, partnerID, merchantID string, fields []string) (*partner.MerchantDetailsData, error) {
	endpoint := fmt.Sprintf(showAccountTrackngRoute+"/%s", url.PathEscape(partnerID), url.PathEscape(merchantID))
	if len(fields) > 0 {
		endpoint += "?fields=" + url.QueryEscape(strings.Join(fields, ","))
	}
	return c.getMerchantData(ctx, endpoint)
}

func (c *Client) getMerchantData(ctx context.Context, endpoint string) (*partner.MerchantDetailsData, error) {
	r := &request{
		client:   c,
		method:   http.MethodGet,
		endpoint: endpoint,
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status == http.StatusOK {
		r := &partner.MerchantDetailsData{}
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

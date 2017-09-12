package market

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/greater-commons/paypal-marketplace/orders"
)

const (
	createOrderRoute           = "/v1/checkout/orders"
	getTransactionContextRoute = "/v1/risk/transaction-contexts"
)

func (c *Client) CreateOrder(ctx context.Context, trackingID string, params *orders.CreateOrderParams) (*orders.CreateOrderResponse, error) {
	d, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	r := &request{
		client:   c,
		method:   http.MethodPost,
		endpoint: createOrderRoute,
		body:     bytes.NewReader(d),
		headers: map[string][]string{
			"Paypal-Client-Metadata-Id": []string{trackingID},
		},
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status == http.StatusCreated {
		r := &orders.CreateOrderResponse{}
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

func (c *Client) CancelOrder(ctx context.Context, orderID string) error {
	r := &request{
		client:   c,
		method:   http.MethodDelete,
		endpoint: createOrderRoute + "/" + url.PathEscape(orderID),
	}
	res, err := r.do(ctx)
	if err != nil {
		return err
	}
	if res.status == http.StatusNoContent {
		return nil
	}
	errorData, err := ioutil.ReadAll(res.body)
	if err != nil {
		return err
	}
	return &BadResponse{
		Status: res.status,
		Body:   string(errorData),
	}
}

func (c *Client) SaveTransactionContext(ctx context.Context, merchantID, trackingID string, additionalData []orders.KeyValuePair) error {
	body := struct {
		AdditionalData []orders.KeyValuePair `json:"additional_data,omitempty"`
	}{
		AdditionalData: additionalData,
	}
	d, err := json.Marshal(&body)
	if err != nil {
		return err
	}
	r := &request{
		client:   c,
		method:   http.MethodPut,
		endpoint: getTransactionContextRoute + "/" + url.PathEscape(merchantID) + "/" + url.PathEscape(trackingID),
		body:     bytes.NewReader(d),
	}
	res, err := r.do(ctx)
	if err != nil {
		return err
	}
	if res.status == http.StatusOK {
		return nil
	}
	errorData, err := ioutil.ReadAll(res.body)
	if err != nil {
		return err
	}
	return &BadResponse{
		Status: res.status,
		Body:   string(errorData),
	}
}

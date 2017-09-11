package market

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/greater-commons/paypal-marketplace/orders"
)

const createOrderRoute = "/v1/checkout/orders"

func (c *Client) CreateOrder(ctx context.Context, params *orders.CreateOrderParams) (*orders.CreateOrderResponse, error) {
	d, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	r := &request{
		client:   c,
		method:   http.MethodPost,
		endpoint: createOrderRoute,
		body:     bytes.NewReader(d),
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status == http.StatusOK {
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

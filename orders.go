package market

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/greater-commons/paypal-marketplace/orders"
)

const (
	createOrderRoute           = "/v1/checkout/orders"
	getTransactionContextRoute = "/v1/risk/transaction-contexts"
	disbursePaymentsRoute      = "/v1/payments/referenced-payouts-items"
	requestRefundRoute         = "/v1/payments/capture/"
)

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

func (c *Client) GetOrderDetails(ctx context.Context, orderID string) (*orders.CreateOrderResponse, error) {
	r := &request{
		client:   c,
		method:   http.MethodGet,
		endpoint: createOrderRoute + "/" + url.PathEscape(orderID),
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status == http.StatusOK {
		r := &orders.CreateOrderResponse{}
		err := json.NewDecoder(res.body).Decode(r)
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

func (c *Client) PayOrder(ctx context.Context, orderID string, disbursementMode orders.DisbursementModeData) (*orders.PayOrderResponse, error) {
	data := struct {
		DisbursementMode orders.DisbursementModeData `json:"disbursement_mode"`
	}{
		DisbursementMode: disbursementMode,
	}
	d, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	r := &request{
		client:   c,
		method:   http.MethodPost,
		endpoint: createOrderRoute + "/" + url.PathEscape(orderID) + "/pay",
		body:     bytes.NewReader(d),
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status >= 200 && res.status < 300 {
		r := &orders.PayOrderResponse{}
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

func (c *Client) FinalizeDisbursement(ctx context.Context, responsePreference orders.ResponsePreferenceData, transactionID string) (*orders.FinalizeDisbursementResponse, error) {
	data := struct {
		ReferenceID   string `json:"reference_id"`
		ReferenceType string `json:"reference_type"`
	}{
		ReferenceID:   transactionID,
		ReferenceType: "TRANSACTION_ID",
	}
	d, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	r := &request{
		client:   c,
		method:   http.MethodPost,
		endpoint: disbursePaymentsRoute,
		body:     bytes.NewReader(d),
		headers: map[string][]string{
			"Prefer": []string{string(responsePreference)},
		},
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status == http.StatusOK {
		r := &orders.FinalizeDisbursementResponse{}
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

func (c *Client) RequestRefund(ctx context.Context, captureID, clientID, payerID string, params *orders.RequestRefundParams) (*orders.RequestRefundResponse, error) {
	d, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	authHeader := base64.StdEncoding.EncodeToString([]byte(`{"alg":"none"}`)) + "." + base64.StdEncoding.EncodeToString([]byte(`{"iss":"`+clientID+`","payer_id":"`+payerID+`"}`)) + "."
	r := &request{
		client:   c,
		method:   http.MethodPost,
		endpoint: requestRefundRoute + captureID + "/refund",
		body:     bytes.NewReader(d),
		headers: map[string][]string{
			"PayPal-Auth-Assertion": []string{authHeader},
		},
	}
	res, err := r.do(ctx)
	if err != nil {
		return nil, err
	}
	if res.status == http.StatusCreated {
		r := &orders.RequestRefundResponse{}
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

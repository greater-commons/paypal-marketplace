package market

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	Sandbox = "https://api.sandbox.paypal.com"
	Live    = "https://api.paypal.com"
)

const (
	tokenRoute = "/v1/oauth2/token"
)

type Client struct {
	client  *http.Client
	apiBase string
	BNCode  string
}

type BadResponse struct {
	Status int
	Body   string
}

func (b *BadResponse) Error() string {
	return "Bad response from Paypal, status code: " + strconv.Itoa(b.Status) + ", body is:\n" + b.Body
}

// NewClient creates a new Paypal marketplace client.
// clientID and clientSecret are provided by Paypal.
// apiBase should be either market.Sandbox or market.Live
func NewClient(ctx context.Context, clientID, clientSecret, apiBase string) *Client {
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     apiBase + tokenRoute,
	}
	c := &Client{
		client:  conf.Client(ctx),
		apiBase: apiBase,
	}
	return c
}

type request struct {
	client   *Client
	method   string
	endpoint string
	body     io.Reader
	headers  http.Header
}

type response struct {
	status  int
	headers http.Header
	body    io.ReadSeeker
}

func (r *request) do(ctx context.Context) (*response, error) {
	req, err := http.NewRequest(r.method, r.client.apiBase+r.endpoint, r.body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, vs := range r.headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	if r.client.BNCode != "" {
		req.Header.Set("PayPal-Partner-Attribution-Id", r.client.BNCode)
	}
	res, err := r.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(io.LimitReader(res.Body, 5*1024*1024)) // Read at most 5 MB
	if err != nil {
		return nil, err
	}
	return &response{
		status:  res.StatusCode,
		headers: res.Header,
		body:    bytes.NewReader(resData),
	}, nil
}

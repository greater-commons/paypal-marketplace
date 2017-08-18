package market

import (
	"context"
	"io"
	"net/http"
	"path"

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

// NewClient creates a new Paypal marketplace client.
// clientID and clientSecret are provided by Paypal.
// apiBase should be either market.Sandbox or market.Live
func NewClient(ctx context.Context, clientID, clientSecret, apiBase string) *Client {
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     path.Join(apiBase, tokenRoute),
	}
	return &Client{
		client:  conf.Client(ctx),
		apiBase: apiBase,
	}
}

type request struct {
	client   *Client
	method   string
	endpoint string
	body     io.Reader
}

func (r *request) do(ctx context.Context, res interface{}) error {
	req := http.NewRequest(r.method, r.client.apiBase+r.endpoint, r.body)
	response, err := r.client.client.Do(req)
	if err != nil {
		return err
	}
}

package market

import "context"

// CreatePartnerReferral is used to connect a user's Paypal account with your platform.
// It is used in both the connected and the managed paths.
func (c *Client) CreatePartnerReferral(ctx context.Context, params *partner.CreatePartnerReferralParams) (*partner.CreatePartnerReferralResponse, error) {
}

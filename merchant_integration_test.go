package market

import (
	"context"
	"testing"
)

func TestShowAccountTracking(t *testing.T) {
	ctx := context.Background()
	c := NewClient(ctx, GetTestClientID(), GetTestSecret(), Sandbox)
	c.BNCode = GetTestBNCode()

	tRes, err := c.ShowAccountTracking(ctx, GetTestPayerID(), GetTrackingID(t))
	if err != nil {
		t.Fatal("Error attempting to create a partner referral:", err)
	}
	t.Logf("Account Tracking: %+v\n", tRes)
}

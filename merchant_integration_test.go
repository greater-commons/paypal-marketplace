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
		t.Fatal("Error attempting to get account tracking:", err)
	}
	t.Logf("Account Tracking: %+v\n", tRes)
}

func TestShowMerchantStatus(t *testing.T) {
	ctx := context.Background()
	c := NewClient(ctx, GetTestClientID(), GetTestSecret(), Sandbox)
	c.BNCode = GetTestBNCode()

	tRes, err := c.ShowMerchantStatus(ctx, GetTestPayerID(), GetMerchantID(t), []string{"payments_receivable"})
	if err != nil {
		t.Fatal("Error attempting to get merchant status:", err)
	}
	t.Logf("Merchant Status: %+v\n", tRes)
}

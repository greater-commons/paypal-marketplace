package market

import (
	"os"
	"testing"
)

func GetTestClientID() string {
	cid := os.Getenv("PAYPAL_CLIENT_ID")
	if len(cid) == 0 {
		panic("PAYPAL_CLIENT_ID environment variable is not set, but is needed to run tests!\n")
	}
	return cid
}

func GetTestSecret() string {
	secret := os.Getenv("PAYPAL_SECRET")
	if len(secret) == 0 {
		panic("PAYPAL_SECRET environment variable is not set, but is needed to run tests!\n")
	}
	return secret
}

func GetTestBNCode() string {
	bn := os.Getenv("PAYPAL_BN_CODE")
	if len(bn) == 0 {
		panic("PAYPAL_BN_CODE environment variable is not set, but is needed to run tests!\n")
	}
	return bn
}

func GetTestPayerID() string {
	pid := os.Getenv("PAYPAL_PAYER_ID")
	if len(pid) == 0 {
		panic("PAYPAL_PAYER_ID environment variable is not set, but is needed to run tests!\n")
	}
	return pid
}

func GetPartnerReferralID(t *testing.T) string {
	partnerID := os.Getenv("PAYPAL_PARTNER_REFERRAL_ID")
	if partnerID == "" {
		t.Skip("PAYPAL_PARTNER_REFERRAL_ID environment variable is not set, but is needed for some tests.\n")
	}
	return partnerID
}

func GetTrackingID(t *testing.T) string {
	trackID := os.Getenv("PAYPAL_MERCHANT_TRACKING_ID")
	if trackID == "" {
		t.Skip("PAYPAL_MERCHANT_TRACKING_ID environment variable is not set, but is needed for some tests.\n")
	}
	return trackID
}

func GetMerchantID(t *testing.T) string {
	merchID := os.Getenv("PAYPAL_MERCHANT_ID")
	if merchID == "" {
		t.Skip("PAYPAL_MERCHANT_ID environment variable is not set, but is needed for some tests.\n")
	}
	return merchID
}

package market

import (
	"context"
	"testing"

	"github.com/greater-commons/paypal-marketplace/partner"
)

func TestCreatePartnerReferral(t *testing.T) {
	ctx := context.Background()
	c := NewClient(ctx, GetTestClientID(), GetTestSecret(), Sandbox)
	c.BNCode = GetTestBNCode()

	_, err := c.CreatePartnerReferral(ctx, &partner.CreatePartnerReferralParams{
		CustomerData: &partner.UserData{
			CustomerType: partner.CustomerTypeMerchant,
		},
		RequestedCapabilities: []partner.CustomerCapabilitiesData{
			{
				Capability: partner.CapabilityApiIntegration,
				ApiIntegrationPreference: &partner.IntegrationDetailsData{
					PartnerID: GetTestBNCode(),
					RestAPIIntegration: &partner.RestAPIIntegrationData{
						IntegrationMethod: partner.IntegrationMethodPaypal,
						IntegrationType:   partner.IntegrationTypeThirdParty,
					},
					RestThirdPartyDetails: &partner.RestThirdPartyDetailsData{
						PartnerClientID: GetTestClientID(),
						FeatureList: []partner.ReferralDataRestFeaturesData{
							partner.ReferralDataRestFeaturesPayment,
						},
					},
				},
			},
		},
		CollectedConsents: []partner.LegalConsentData{
			{
				Type:    partner.LegalConsentTypeShareDataConsent,
				Granted: true,
			},
		},
		Products: []partner.ReferralDataProductNameData{
			partner.ReferralDataExpressCheckout,
		},
	})
	if err != nil {
		t.Fatal("Error attempting to create a partner referral:", err)
	}
}

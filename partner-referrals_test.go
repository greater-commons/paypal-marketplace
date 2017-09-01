package market

import (
	"context"
	"testing"
	"time"

	"github.com/greater-commons/paypal-marketplace/partner"
)

func TestCreatePartnerReferral(t *testing.T) {
	ctx := context.Background()
	c := NewClient(ctx, GetTestClientID(), GetTestSecret(), Sandbox)
	c.BNCode = GetTestBNCode()

	r, err := c.CreatePartnerReferral(ctx, &partner.CreatePartnerReferralParams{
		CustomerData: &partner.UserData{
			CustomerType: partner.CustomerTypeMerchant,
			PersonDetails: &partner.PersonDetailsData{
				EmailAddress: "test1234@example.com",
				Name: &partner.NameOfAPartyData{
					GivenName: "Test",
					Surname:   "Test1",
				},
				DateOfBirth: &partner.DateData{
					EventType: partner.EventTypeBirth,
					EventDate: time.Date(1998, time.February, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			BusinessDetails: &partner.BusinessDetailsData{
				Category: partner.CategoryEducation{
					partner.EducationSubCategoryColleges,
				},
			},
		},
		RequestedCapabilities: []partner.CustomerCapabilitiesData{
			{
				Capability: partner.CapabilityApiIntegration,
				ApiIntegrationPreference: &partner.IntegrationDetailsData{
					PartnerID: GetTestPayerID(),
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
	t.Logf("%+v", r)
	if err != nil {
		t.Fatal("Error attempting to create a partner referral:", err)
	}
}

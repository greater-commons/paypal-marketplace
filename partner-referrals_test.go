package market

import (
	"context"
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/greater-commons/paypal-marketplace/partner"
)

func TestCreatePartnerReferral(t *testing.T) {
	ctx := context.Background()
	c := NewClient(ctx, GetTestClientID(), GetTestSecret(), Sandbox)
	c.BNCode = GetTestBNCode()

	num, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	testEmail := "greatercommons-" + strconv.FormatInt(num.Int64(), 10) + "@test.com"

	r, err := c.CreatePartnerReferral(ctx, &partner.CreatePartnerReferralParams{
		CustomerData: &partner.UserData{
			CustomerType: partner.CustomerTypeMerchant,
			PersonDetails: &partner.PersonDetailsData{
				EmailAddress: testEmail,
				Name: &partner.NameOfAPartyData{
					GivenName: "Test",
					Surname:   "Test",
				},
				HomeAddress: &partner.SimplePostalAddressData{
					Line1:       "123 Test Ave",
					City:        "Austin",
					State:       "TX",
					CountryCode: "US",
					PostalCode:  "78701",
				},
				DateOfBirth: &partner.DateData{
					EventType: partner.EventTypeBirth,
					EventDate: time.Date(1998, time.February, 2, 0, 0, 0, 0, time.UTC),
				},
				IdentityDocuments: []partner.IdentityDocumentData{
					{
						Type:              partner.IdentityTypeSocialSecurityNumber,
						Value:             "1234",
						PartialValue:      true,
						IssuerCountryCode: "US",
					},
				},
				PhoneContacts: []partner.OnboardingCommonUserPhoneData{
					{
						PhoneNumberDetails: &partner.PhoneDetailsData{
							CountryCode:    "1",
							NationalNumber: "5121234567",
						},
						PhoneType: partner.PhoneTypeHome,
					},
				},
			},
			BusinessDetails: &partner.BusinessDetailsData{
				BusinessType: partner.BusinessTypeIndividual,
				Names: []partner.BusinessNameData{
					{
						Type: partner.BusinessNameTypeLegal,
						Name: "Test Test's Store",
					},
				},
				Category: partner.CategoryEducation{
					partner.EducationSubCategoryVocational,
				},
				BusinessAddress: &partner.SimplePostalAddressData{
					Line1:       "123 Test Ave",
					City:        "Austin",
					State:       "TX",
					CountryCode: "US",
					PostalCode:  "78701",
				},
				PhoneContacts: []partner.OnboardingCommonUserPhoneData{
					{
						PhoneNumberDetails: &partner.PhoneDetailsData{
							CountryCode:    "1",
							NationalNumber: "5127654321",
						},
						PhoneType: partner.PhoneTypeFax,
					},
				},
				AverageMonthlyVolumeRange: &partner.CurrencyRangeData{
					MinimumAmount: &partner.CurrencyData{
						Currency: "USD",
						Value:    "0",
					},
					MaximumAmount: &partner.CurrencyData{
						Currency: "USD",
						Value:    "4999",
					},
				},
			},
			ReferralUserPayerID: &partner.AccountIdentifierData{
				Type:  partner.AccountIdentifierTypePayerID,
				Value: GetTestPayerID(),
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
							partner.ReferralDataRestFeaturesRefund,
							partner.ReferralDataRestFeaturesPartnerFee,
							partner.ReferralDataRestFeaturesDelayDisbursement,
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

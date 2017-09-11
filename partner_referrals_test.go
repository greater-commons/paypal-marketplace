package market

import (
	"context"
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/greater-commons/paypal-marketplace/merchant"
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

	r, err := c.CreatePartnerReferral(ctx, &merchant.CreatePartnerReferralParams{
		CustomerData: &merchant.UserData{
			CustomerType: merchant.CustomerTypeMerchant,
			PersonDetails: &merchant.PersonDetailsData{
				EmailAddress: testEmail,
				Name: &merchant.NameOfAPartyData{
					GivenName: "Test",
					Surname:   "Test",
				},
				HomeAddress: &merchant.SimplePostalAddressData{
					Line1:       "123 Test Ave",
					City:        "Austin",
					State:       "TX",
					CountryCode: "US",
					PostalCode:  "78701",
				},
				DateOfBirth: &merchant.DateData{
					EventType: merchant.EventTypeBirth,
					EventDate: time.Date(1998, time.February, 2, 0, 0, 0, 0, time.UTC),
				},
				IdentityDocuments: []merchant.IdentityDocumentData{
					{
						Type:              merchant.IdentityTypeSocialSecurityNumber,
						Value:             "1234",
						PartialValue:      true,
						IssuerCountryCode: "US",
					},
				},
				PhoneContacts: []merchant.OnboardingCommonUserPhoneData{
					{
						PhoneNumberDetails: &merchant.PhoneDetailsData{
							CountryCode:    "1",
							NationalNumber: "5121234567",
						},
						PhoneType: merchant.PhoneTypeHome,
					},
				},
			},
			BusinessDetails: &merchant.BusinessDetailsData{
				BusinessType: merchant.BusinessTypeIndividual,
				Names: []merchant.BusinessNameData{
					{
						Type: merchant.BusinessNameTypeLegal,
						Name: "Test Test's Store",
					},
				},
				Category: merchant.CategoryEducation{
					merchant.EducationSubCategoryVocational,
				},
				BusinessAddress: &merchant.SimplePostalAddressData{
					Line1:       "123 Test Ave",
					City:        "Austin",
					State:       "TX",
					CountryCode: "US",
					PostalCode:  "78701",
				},
				PhoneContacts: []merchant.OnboardingCommonUserPhoneData{
					{
						PhoneNumberDetails: &merchant.PhoneDetailsData{
							CountryCode:    "1",
							NationalNumber: "5127654321",
						},
						PhoneType: merchant.PhoneTypeFax,
					},
				},
				AverageMonthlyVolumeRange: &merchant.CurrencyRangeData{
					MinimumAmount: &merchant.CurrencyData{
						Currency: "USD",
						Value:    "0",
					},
					MaximumAmount: &merchant.CurrencyData{
						Currency: "USD",
						Value:    "4999",
					},
				},
			},
			ReferralUserPayerID: &merchant.AccountIdentifierData{
				Type:  merchant.AccountIdentifierTypePayerID,
				Value: GetTestPayerID(),
			},
			PartnerSpecificIdentifiers: []merchant.PartnerSpecificIdentifierData{
				{
					Type:  merchant.PartnerSpecificIdentifierTypeTrackingID,
					Value: strconv.FormatInt(num.Int64(), 10),
				},
			},
		},
		RequestedCapabilities: []merchant.CustomerCapabilitiesData{
			{
				Capability: merchant.CapabilityApiIntegration,
				ApiIntegrationPreference: &merchant.IntegrationDetailsData{
					PartnerID: GetTestPayerID(),
					RestAPIIntegration: &merchant.RestAPIIntegrationData{
						IntegrationMethod: merchant.IntegrationMethodPaypal,
						IntegrationType:   merchant.IntegrationTypeThirdParty,
					},
					RestThirdPartyDetails: &merchant.RestThirdPartyDetailsData{
						PartnerClientID: GetTestClientID(),
						FeatureList: []merchant.ReferralDataRestFeaturesData{
							merchant.ReferralDataRestFeaturesPayment,
							merchant.ReferralDataRestFeaturesRefund,
							merchant.ReferralDataRestFeaturesPartnerFee,
							merchant.ReferralDataRestFeaturesDelayDisbursement,
						},
					},
				},
			},
		},
		CollectedConsents: []merchant.LegalConsentData{
			{
				Type:    merchant.LegalConsentTypeShareDataConsent,
				Granted: true,
			},
		},
		Products: []merchant.ReferralDataProductNameData{
			merchant.ReferralDataExpressCheckout,
		},
	})
	if err != nil {
		t.Fatal("Error attempting to create a partner referral:", err)
	}
	t.Logf("Response after creating: %+v\n", r)
}

func TestGetPartnerReferral(t *testing.T) {
	ctx := context.Background()

	c := NewClient(ctx, GetTestClientID(), GetTestSecret(), Sandbox)
	c.BNCode = GetTestBNCode()

	pr, err := c.GetPartnerReferral(ctx, GetPartnerReferralID(t))
	if err != nil {
		t.Fatal("Error attempting to get a partner referral:", err)
	}
	t.Logf("Partner Referral: %+v\n", pr)
}

package partner

import (
	"encoding/json"
	"time"
)

type CustomerTypeData string

const (
	CustomerTypeConsumer CustomerTypeData = "CONSUMER"
	CustomerTypeMerchant CustomerTypeData = "MERCHANT"
)

type NameOfAPartyData struct {
	Prefix            string `json:"prefix"`
	GivenName         string `json:"given_name"`
	Surname           string `json:"surname"`
	MiddleName        string `json:"middle_name"`
	Suffix            string `json:"suffix"`
	AlternameFullName string `json:"altername_full_name"`
}

type PhoneTypeData struct {
	CountryCode     string `json:"country_code,omitempty"`
	NationalNumber  string `json:"national_number,omitempty"`
	ExtensionNumber string `json:"extension_number,omitempty"`
}

type OnboardingCommonUserPhoneData struct {
	PhoneNumberDetails *PhoneTypeData `json:"phone_number_details,omitempty"`
	PhoneType          string         `json:"phone_type,omitempty"`
}

type SimplePostalAddressData struct {
	Line1       string `json:"line1,omitempty"`
	Line2       string `json:"line2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	PostalCode  string `json:"postal_cord,omitempty"`
}

type DateData struct {
	EventType string    `json:"event_type,omitempty"`
	EventDate time.Time `json:"event_date,omitempty"`
}

func (d *DateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EventType string `json:"event_type,omitempty"`
		EventDate string `json:"event_date,omitempty"`
	}{
		EventType: d.EventType,
		EventDate: d.EventDate.Format(time.RFC3339),
	})
}

func (d *DateData) UnmarshalJSON(b []byte) error {
	temp := struct {
		EventType string `json:"event_type,omitempty"`
		EventDate string `json:"event_date,omitempty"`
	}{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	d.EventType = temp.EventType
	date, err := time.Parse(time.RFC3339, temp.EventDate)
	if err != nil {
		return err
	}
	d.EventDate = date
	return nil
}

type IdentityDocumentData struct {
	Type              string `json:"type,omitempty"`
	Value             string `json:"value,omitempty"`
	PartialValue      *bool  `json:"partial_value,omitempty"`
	IssuerCountryCode string `json:"issuer_country_code,omitempty"`
}

type RelationData string

const (
	RelationMother RelationData = "MOTHER"
)

type MerchantRelationData struct {
	Name                     *NameOfAPartyData `json:"name,omitempty"`
	Relation                 RelationData      `json:"relation,omitempty"`
	CountryCodeOfNationality string            `json:"country_code_of_nationality,omitempty"`
}

type PersonDetailsData struct {
	EmailAddress              string                          `json:"email_address,omitempty"`
	Name                      *NameOfAPartyData               `json:"name,omitempty"`
	PhoneContacts             []OnboardingCommonUserPhoneData `json:"phone_contacts,omitempty"`
	HomeAddress               *SimplePostalAddressData        `json:"home_address,omitempty"`
	DateOfBirth               *DateData                       `json:"date_of_birth,omitempty"`
	NationalityCountryCode    string                          `json:"nationality_country_code,omitempty"`
	IdentityDocuments         []IdentityDocumentData          `json:"identity_documents,omitempty"`
	AccountOwnerRelationships []MerchantRelationData          `json:"account_owner_relationships,omitempty"`
}

type BusinessTypeData string

const (
	BusinessTypeIndividual                  BusinessTypeData = "INDIVIDUAL"
	BusinessTypeProprietorship              BusinessTypeData = "PROPRIETORSHIP"
	BusinessTypePartnership                 BusinessTypeData = "PARTNERSHIP"
	BusinessTypeCorporation                 BusinessTypeData = "CORPORATION"
	BusinessTypeNonprofit                   BusinessTypeData = "NONPROFIT"
	BusinessTypeGovernment                  BusinessTypeData = "GOVERNMENT"
	BusinessTypePublicCompany               BusinessTypeData = "PUBLIC_COMPANY"
	BusinessTypeRegisteredCooperative       BusinessTypeData = "REGISTERED_COOPERATIVE"
	BusinessTypeProprietoryCompany          BusinessTypeData = "PROPRIETORY_COMPANY"
	BusinessTypeAssociation                 BusinessTypeData = "ASSOCIATION"
	BusinessTypePrivateCorporation          BusinessTypeData = "PRIVATE_CORPORATION"
	BusinessTypeLimitedPartnership          BusinessTypeData = "LIMITED_PARTNERSHIP"
	BusinessTypeLimitedLiabilityProprietors BusinessTypeData = "LIMITED_LIABILITY_PROPRIETORS"
	BusinessTypeLimitedLiabilityPartnership BusinessTypeData = "LIMITED_LIABILITY_PARTNERSHIP"
	BusinessTypePublicCorporation           BusinessTypeData = "PUBLIC_CORPORATION"
	BusinessTypeOtherCorporateBody          BusinessTypeData = "OTHER_CORPORATE_BODY"
)

type BusinessNameTypeData string

const (
	BusinessNameTypeLegal            BusinessNameTypeData = "LEGAL"
	BusinessNameTypeDoingBusinessAs  BusinessNameTypeData = "DOING_BUSINESS_AS"
	BusinessNameTypeStockTradingName BusinessNameTypeData = "STOCK_TRADING_NAME"
)

type BusinessNameData struct {
	Type BusinessNameTypeData `json:"type,omitempty"`
	Name string               `json:"name,omitempty"`
}

type CurrencyData struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:value,omitempty"`
}

type CurrencyRangeData struct {
	MinimumAmount *CurrencyData `json:"minimum_amount,omitempty"`
	MaximumAmount *CurrencyData `json:"maximum_amount,omitempty"`
}

type EmailRoleData string

const (
	EmailRoleCustomerService EmailRoleData = "CUSTOMER_SERVICE"
)

type EmailData struct {
	EmailAddress string        `json:"email_address,omitempty"`
	Role         EmailRoleData `json:"role,omitempty"`
}

type BusinessDetailsData struct {
	PhoneContacts             []OnboardingCommonUserPhoneData `json:"phone_contacts,omitempty"`
	BusinessAddress           *SimplePostalAddressData        `json:"business_address,omitempty"`
	BusinessType              BusinessTypeData                `json:"business_type,omitempty"`
	Category                  string                          `json:"category,omitempty"`
	SubCategory               string                          `json:"sub_category,omitempty"`
	Names                     []BusinessNameData              `json:"names,omitempty"`
	BusinessDescription       string                          `json:"business_description,omitempty"`
	EventDates                []DateData                      `json:"event_dates,omitempty"`
	WebsiteURLS               []string                        `json:"website_urls"`
	AnnualSalesVolumeRange    *CurrencyRangeData              `json:"annual_sales_volume_range,omitempty"`
	AverageMonthlyVolumeRange *CurrencyRangeData              `json:"average_monthly_volume_range,omitempty"`
	IdentityDocuments         []IdentityDocumentData          `json:"identity_documents,omitempty"`
	EmailContacts             []EmailData                     `json:"email_contacts,omitempty"`
}

type BankAccountTypeData string

const (
	BankAccountTypeChecking BankAccountTypeData = "CHECKING"
	BankAccountTypeSavings  BankAccountTypeData = "SAVINGS"
)

type BankDetailsData struct {
	NickName       string                   `json:"nick_name,omitempty"`
	AccountNumber  string                   `json:"account_number,omitempty"`
	AccountType    BankAccountTypeData      `json:"account_type,omitempty"`
	CurrencyCode   string                   `json:"currency_code,omitempty"`
	Identifiers    []string                 `json:"identifiers,omitempty"`
	BranchLocation *SimplePostalAddressData `json:"branch_location,omitempty"`
	MandateAgreed  *bool                    `json:"mandate_agreed,omitempty"`
}

type FinancialInstrumentDataType struct {
	BankDetails *BankDetailsData `json:"bank_details,omitempty"`
}

type AccountIdentifierTypeData string

const (
	AccountIdentifierTypePayerID AccountIdentifierTypeData = "PAYER_ID"
)

type AccountIdentifierData struct {
	Type  AccountIdentifierTypeData `json:"type,omitempty"`
	Value string                    `json:"value,omitempty"`
}

type PartnerSpecificIdentifierTypeData string

const (
	PartnerSpecificIdentifierTypeTrackingID       PartnerSpecificIdentifierTypeData = "TRACKING_ID"
	PartnerSpecificIdentifierTypeAccountLinkingID PartnerSpecificIdentifierTypeData = "ACCOUNT_LINKING_ID"
)

type PartnerSpecificIdentifierData struct {
	Type  PartnerSpecificIdentifierTypeData `json:"type,omitempty"`
	Value string                            `json:"value,omitempty"`
}

type UserData struct {
	CustomerType               CustomerTypeData                `json:"customer_type,omitempty"`
	PersonDetails              *PersonDetailsData              `json:"person_details,omitempty"`
	BusinessDetails            *BusinessDetailsData            `json:"business_details,omitempty"`
	FinancialInstrumentData    *FinancialInstrumentDataType    `json:"financial_instrument_data,omitempty"`
	PreferredLanguageCode      string                          `json:"preferred_language_code,omitempty"`
	PrimaryCurrencyCode        string                          `json:"primary_currency_code,omitempty"`
	ReferralUserPayerID        *AccountIdentifierData          `json:"referral_user_payer_id,omitempty"`
	PartnerSpecificIdentifiers []PartnerSpecificIdentifierData `json:"partner_specific_identifiers,omitempty"`
}

type CapabilityData string

const (
	CapabilityApiIntegration             CapabilityData = "API_INTEGRATION"
	CapabilityBankAddition               CapabilityData = "BANK_ADDITION"
	CapabilityBillingAgreement           CapabilityData = "BILLING_AGREEMENT"
	CapabilityContextualMarketingConsent CapabilityData = "CONTEXTUAL_MARKETING_CONSENT"
)

type ClassicIntegrationTypeData string

const (
	ClassicIntegrationThirdParty              ClassicIntegrationTypeData = "THIRD_PARTY"
	ClassicIntegrationFirstPartyIntegrated    ClassicIntegrationTypeData = "FIRST_PARTY_INTEGRATED"
	ClassicIntegrationFirstPartyNonIntegrated ClassicIntegrationTypeData = "FIRST_PARTY_NON_INTEGRATED"
)

type IntegrationMethodData string

const (
	IntegrationMethodBrainTree IntegrationMethodData = "BRAINTREE"
	IntegrationMethodPaypal    IntegrationMethodData = "PAYPAL"
)

type IntegrationTypeData string

const (
	IntegrationTypeThirdParty IntegrationTypeData = "THIRD_PARTY"
)

type RestAPIIntegrationData struct {
	IntegrationMethod *IntegrationMethodData `json:"integration_method,omitempty"`
	IntegrationType   *IntegrationTypeData   `json:"integration_type,omitempty"`
}

type ReferralDataClassicPermissionData string

const (
	ExpressCheckout                ReferralDataClassicPermissionData = "EXPRESS_CHECKOUT"
	Refund                         ReferralDataClassicPermissionData = "REFUND"
	DirectPayment                  ReferralDataClassicPermissionData = "DIRECT_PAYMENT"
	AuthCapture                    ReferralDataClassicPermissionData = "AUTH_CAPTURE"
	ButtonManager                  ReferralDataClassicPermissionData = "BUTTON_MANAGER"
	AccountBalance                 ReferralDataClassicPermissionData = "ACCOUNT_BALANCE"
	TransactionDetails             ReferralDataClassicPermissionData = "TRANSACTION_DETAILS"
	TransactionSearch              ReferralDataClassicPermissionData = "TRANSACTION_SEARCH"
	ReferenceTransaction           ReferralDataClassicPermissionData = "REFERENCE_TRANSACTION"
	RecurringPayments              ReferralDataClassicPermissionData = "RECURRING_PAYMENTS"
	ManagePendingTransactionStatus ReferralDataClassicPermissionData = "MANAGE_PENDING_TRANSACTION_STATUS"
	NonReferencedCredit            ReferralDataClassicPermissionData = "NON_REFERENCED_CREDIT"
	EncryptedWebsitePayments       ReferralDataClassicPermissionData = "ENCRYPTED_WEBSITE_PAYMENTS"
	MobileCheckout                 ReferralDataClassicPermissionData = "MOBILE_CHECKOUT"
	AirTravel                      ReferralDataClassicPermissionData = "AIR_TRAVEL"
	Invoicing                      ReferralDataClassicPermissionData = "INVOICING"
	AccessBasicPersonalData        ReferralDataClassicPermissionData = "ACCESS_BASIC_PERSONAL_DATA"
)

type SupportedClassicPermissionsData struct {
	ReferralDataClassicPermissions *ReferralDataClassicPermissionData `json:"referral_data-classic_permission_enum,omitempty"`
}

type ClassicThirdPartyDetailsData struct {
	PermissionList []SupportedClassicPermissionsData `json:"permission_list,omitempty"`
}

type ClassicFirstPartyDetailsData string

const (
	ClassicFirstPartySignature   ClassicFirstPartyDetailsData = "SIGNATURE"
	ClassicFirstPartyCertificate ClassicFirstPartyDetailsData = "CERTIFICATE"
)

type ReferralDataRestFeaturesData string

const (
	ReferralDataRestFeaturesPayment       ReferralDataRestFeaturesData = "PAYMENT"
	ReferralDataRestFeaturesRefund        ReferralDataRestFeaturesData = "REFUND"
	ReferralDataRestFeaturesFuturePayment ReferralDataRestFeaturesData = "FUTURE_PAYMENT"
	ReferralDataRestFeaturesDirectPayment ReferralDataRestFeaturesData = "DIRECT_PAYMENT"
	ReferralDataRestFeaturesPartnerFee    ReferralDataRestFeaturesData = "PARTNER_FEE"
)

type RestThirdPartyDetailsData struct {
	PartnerClientID string                         `json:"partner_client_id,omitempty"`
	FeatureList     []ReferralDataRestFeaturesData `json:"feature_list,omitempty"`
}

type IntegrationDetailsData struct {
	PartnerID                 string                        `json:"partner_id"`
	ClassicAPIIntegrationType *ClassicIntegrationTypeData   `json:"classic_api_integration_type,omitempty"`
	RestAPIIntegration        *RestAPIIntegrationData       `json:"rest_api_integration,omitempty"`
	ClassicThirdPartyDetails  *ClassicThirdPartyDetailsData `json:"classic_third_party_details,omitempty"`
	ClassicFirstPartyDetails  *ClassicFirstPartyDetailsData `json:"classic_first_party_details,omitempty"`
	RestThirdPartyDetails     *RestThirdPartyDetailsData    `json:"rest_third_party_details,omitempty"`
}

type BillingExperiencePreferenceData struct {
	ExperienceID      string `json:"experience_id,omitempty"`
	BillingContextSet *bool  `json:"billing_context_set,omitempty"`
}

type BillingAgreementData struct {
	Description                 string                           `json:"description,omitempty"`
	BillingExperiencePreference *BillingExperiencePreferenceData `json:"billing_experience_preference,omitempty"`
	MerchantCustomData          string                           `json:"merchant_custom_data,omitempty"`
	ApprovalURL                 string                           `json:"approval_url,omitempty"`
	ECToken                     string                           `json:"ec_token,omitempty"`
}

type CustomerCapabilitiesData struct {
	Capability               *CapabilityData         `json:"capability,omitempty"`
	ApiIntegrationPreference *IntegrationDetailsData `json:"api_integration_preference,omitempty"`
	BillingAgreement         *BillingAgreementData   `json:"billing_agreement,omitempty"`
}

type WebExperiencePreferenceData struct {
	PartnerLogoURL          string `json:"partner_logo_url,omitempty"`
	ReturnURL               string `json:"return_url,omitempty"`
	ReturnURLDescription    string `json:"return_url_description,omitempty"`
	ActionRenewalURL        string `json:"action_renewal_url,omitempty"`
	ShowAddCreditCard       *bool  `json:"show_add_credit_card,omitempty"`
	ShowMobileConfirm       *bool  `json:"show_mobile_confirm,omitempty"`
	UseMiniBrowser          *bool  `json:"use_mini_browser,omitempty"`
	UseHuaEmailConfirmation *bool  `json:"use_hua_email_confirmation,omitempty"`
}

type LegalConsentTypeData string

const (
	LegalConsentTypeShareDataConsent LegalConsentTypeData = "SHARE_DATA_CONSENT"
)

type LegalConsentData struct {
	Type    LegalConsentTypeData `json:"type"`
	Granted *bool                `json:"granted,omitempty"`
}

type ReferralDataProductNameData string

const (
	ReferralDataExpressCheckout ReferralDataProductNameData = "EXPRESS_CHECKOUT"
)

type ProductsToOnboardData struct {
	ReferralDataProductName ReferralDataProductNameData `json:"referral_data-product_name,omitempty"`
}

type CreatePartnerReferralParams struct {
	CustomerData            *UserData                    `json:"customer_data,omitempty"`
	RequestedCapabilities   []CustomerCapabilitiesData   `json:"requested_capabilities,omitempty"`
	WebExperiencePreference *WebExperiencePreferenceData `json:"web_experience_preference,omitempty"`
	CollectedConsents       []LegalConsentData           `json:"collected_consents,omitempty"`
	Products                []ProductsToOnboardData      `json:"products,omitempty"`
}

package partner

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
	IntegrationMethod IntegrationMethodData `json:"integration_method"`
	IntegrationType   IntegrationTypeData   `json:"integration_type"`
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
	ReferralDataClassicPermissions ReferralDataClassicPermissionData `json:"referral_data-classic_permission_enum"`
}

type ClassicThirdPartyDetailsData struct {
	PermissionList []SupportedClassicPermissionsData `json:"permission_list"`
}

type ClassicFirstPartyDetailsData string

const (
	ClassicFirstPartySignature   ClassicFirstPartyDetailsData = "SIGNATURE"
	ClassicFirstPartyCertificate ClassicFirstPartyDetailsData = "CERTIFICATE"
)

type ReferralDataRestFeaturesData string

type RestThirdPartyDetails struct {
	PartnerClientID string                         `json:"partner_client_id"`
	FeatureList     []ReferralDataRestFeaturesData `json:"feature_list"`
}

type IntegrationDetailsData struct {
	PartnerID                 string                       `json:"partner_id"`
	ClassicAPIIntegrationType ClassicIntegrationTypeData   `json:"classic_api_integration_type"`
	RestAPIIntegration        RestAPIIntegrationData       `json:"rest_api_integration"`
	ClassicThirdPartyDetails  ClassicThirdPartyDetailsData `json:"classic_third_party_details"`
	ClassicFirstPartyDetails  ClassicFirstParyDetailsData  `json:"classic_first_party_details"`
	RestThirdPartyDetails     RestThirdPartyDetailsData    `json:"rest_third_party_details"`
}

type CustomerCapabilities struct {
	Capability               CapabilityData         `json:"capability"`
	ApiIntegrationPreference IntegrationDetailsData `json:"api_integration_preference"`
	BillingAgreement         BillingAgreementData   `json:"billing_agreement"`
}

type CreatePartnerReferralParams struct {
	RequestedCapabilities []CustomerCapabilities `json:"requested_capabilities"`
}

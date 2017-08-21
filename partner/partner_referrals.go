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
	RequestedCapabilities   []CustomerCapabilitiesData   `json:"requested_capabilities,omitempty"`
	WebExperiencePreference *WebExperiencePreferenceData `json:"web_experience_preference,omitempty"`
	CollectedConsents       []LegalConsentData           `json:"collected_consents,omitempty"`
	Products                []ProductsToOnboardData      `json:"products,omitempty"`
}

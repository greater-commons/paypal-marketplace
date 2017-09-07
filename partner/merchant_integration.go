package partner

import "time"

type ProductTypeData string

const (
	ProductExpressCheckout           ProductTypeData = "EXPRESS_CHECKOUT"
	ProductWebsitePaymentsStandard   ProductTypeData = "WEBSITE_PAYMENTS_STANDARD"
	ProductMassPayment               ProductTypeData = "MASS_PAYMENT"
	ProductEmailPayments             ProductTypeData = "EMAIL_PAYMENTS"
	ProductEbayCheckout              ProductTypeData = "EBAY_CHECKOUT"
	ProductPayflowLink               ProductTypeData = "PAYFLOW_LINK"
	ProductPayflowPro                ProductTypeData = "PAYFLOW_PRO"
	ProductWebsitePaymentsPro3       ProductTypeData = "WEBSITE_PAYMENTS_PRO_3_0"
	ProductWebsitePaymentsPro2       ProductTypeData = "WEBSITE_PAYMENTS_PRO_2_0"
	ProductVirtualTerminal           ProductTypeData = "VIRTUAL_TERMINAL"
	ProductHostedSoleSolution        ProductTypeData = "HOSTED_SOLE_SOLUTION"
	ProductBillMeLater               ProductTypeData = "BILL_ME_LATER"
	ProductMobileExpressCheckout     ProductTypeData = "MOBILE_EXPRESS_CHECKOUT"
	ProductPaypalHere                ProductTypeData = "PAYPAL_HERE"
	ProductMobileInStore             ProductTypeData = "MOBILE_IN_STORE"
	ProductPaypalStandard            ProductTypeData = "PAYPAL_STANDARD"
	ProductMobilePaypalStandard      ProductTypeData = "MOBILE_PAYPAL_STANDARD"
	ProductMobilePaymentAcceptance   ProductTypeData = "MOBILE_PAYMENT_ACCEPTANCE"
	ProductPaypalAdvanced            ProductTypeData = "PAYPAL_ADVANCED"
	ProductPaypalPro                 ProductTypeData = "PAYPAL_PRO"
	ProductEnhancedRecurringPayments ProductTypeData = "ENHANCED_RECURRING_PAYMENTS"
)

type VettingStatusData string

const (
	VettingStatusApproved VettingStatusData = "APPROVED"
	VettingStatusPending  VettingStatusData = "PENDING"
	VettingStatusDeclined VettingStatusData = "DECLINED"
)

type ProductData struct {
	Name          ProductTypeData
	VettingStatus VettingStatusData
	Active        bool
}

type SignatureData struct {
	ApiUserName string
	ApiPassword string
	Signature   string
}

type CertificateData struct {
	ApiUserName  string
	ApiPassword  string
	Fingerprint  string
	DownloadLink string
}

type CredentialData struct {
	Signature   SignatureData
	Certificate CertificateData
}

type OAuthIntegrationTypeData string

const (
	OAuthIntegrationTypeFirstPartyIntegrated    OAuthIntegrationTypeData = "FIRST_PARTY_INTEGRATED"
	OAuthIntegrationTypeFirstPartyNonIntegrated OAuthIntegrationTypeData = "FIRST_PARTY_NON_INTEGRATED"
	OAuthIntegrationTypeThirdParty              OAuthIntegrationTypeData = "THIRD_PARTY"
	OAuthIntegrationTypeOAuthThirdParty         OAuthIntegrationTypeData = "OAUTH_THIRD_PARTY"
)

type IntegrationStatusData string

const (
	IntegrationStatusA IntegrationStatusData = "A"
	IntegrationStatusI IntegrationStatusData = "I"
)

type OAuthThirdPartyData struct {
	PartnerClientID  string
	MerchantClientID string
	Scopes           []string
	AccessToken      string
	RefreshToken     string
}

type OAuthIntegrationData struct {
	IntegrationType            OAuthIntegrationTypeData
	IntegrationMethod          IntegrationMethodData
	Status                     IntegrationStatusData
	OAuthThirdPartyIntegration OAuthThirdPartyData
}

type LimitationData struct {
	Name         string
	Restrictions []string
}

type GetAccountTrackingResponse struct {
	TrackingID            string
	MerchantID            string
	Products              []ProductData
	PaymentsReceivable    bool
	PrimaryEmailConfirmed bool
	PrimaryEmail          string
	DateCreated           time.Time
	GrantedPermissions    []string
	ApiCredentials        *CredentialData
	OAuthIntegrations     []OAuthIntegrationData
	Limitations           []LimitationData
}

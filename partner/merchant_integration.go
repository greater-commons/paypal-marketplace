package partner

import (
	"encoding/json"
	"time"
)

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
	Name          ProductTypeData   `json:"name"`
	VettingStatus VettingStatusData `json:"vetting_status"`
	Active        bool              `json:"active"`
}

type SignatureData struct {
	ApiUserName string `json:"api_user_name"`
	ApiPassword string `json:"api_password"`
	Signature   string `json:"signature"`
}

type CertificateData struct {
	ApiUserName  string `json:"api_user_name"`
	ApiPassword  string `json:"api_password"`
	Fingerprint  string `json:"fingerprint"`
	DownloadLink string `json:"download_link"`
}

type CredentialData struct {
	Signature   SignatureData   `json:"signature"`
	Certificate CertificateData `json:"certificate"`
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
	PartnerClientID  string   `json:"partner_client_id"`
	MerchantClientID string   `json:"merchant_client_id"`
	Scopes           []string `json:"scopes"`
	AccessToken      string   `json:"access_token"`
	RefreshToken     string   `json:"refresh_token"`
}

type OAuthIntegrationData struct {
	IntegrationType            OAuthIntegrationTypeData `json:"integration_type"`
	IntegrationMethod          IntegrationMethodData    `json:"integration_method"`
	Status                     IntegrationStatusData    `json:"status"`
	OAuthThirdPartyIntegration []OAuthThirdPartyData    `json:"oauth_third_party"`
}

type LimitationData struct {
	Name         string
	Restrictions []string
}

type MerchantDetailsData struct {
	TrackingID            string                 `json:"tracking_id"`
	MerchantID            string                 `json:"merchant_id"`
	Products              []ProductData          `json:"products"`
	PaymentsReceivable    bool                   `json:"payments_receivable"`
	PrimaryEmailConfirmed bool                   `json:"primary_email_confirmed"`
	PrimaryEmail          string                 `json:"primary_email"`
	DateCreated           time.Time              `json:"date_created"`
	GrantedPermissions    []string               `json:"granted_permissions"`
	ApiCredentials        *CredentialData        `json:"api_credentials"`
	OAuthIntegrations     []OAuthIntegrationData `json:"oauth_integrations"`
	Limitations           []LimitationData       `json:"limitations"`
}

func (m *MerchantDetailsData) UnmarshalJSON(b []byte) error {
	s := struct {
		TrackingID            string                 `json:"tracking_id"`
		MerchantID            string                 `json:"merchant_id"`
		Products              []ProductData          `json:"products"`
		PaymentsReceivable    bool                   `json:"payments_receivable"`
		PrimaryEmailConfirmed bool                   `json:"primary_email_confirmed"`
		PrimaryEmail          string                 `json:"primary_email"`
		DateCreated           string                 `json:"date_created"`
		GrantedPermissions    []string               `json:"granted_permissions"`
		ApiCredentials        *CredentialData        `json:"api_credentials"`
		OAuthIntegrations     []OAuthIntegrationData `json:"oauth_integrations"`
		Limitations           []LimitationData       `json:"limitations"`
	}{}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	m.TrackingID = s.TrackingID
	m.MerchantID = s.MerchantID
	m.Products = s.Products
	m.PaymentsReceivable = s.PaymentsReceivable
	m.PrimaryEmailConfirmed = s.PrimaryEmailConfirmed
	m.PrimaryEmail = s.PrimaryEmail
	if s.DateCreated != "" {
		m.DateCreated, err = time.Parse(s.DateCreated, time.RFC3339)
		if err != nil {
			return err
		}
	}
	m.GrantedPermissions = s.GrantedPermissions
	m.ApiCredentials = s.ApiCredentials
	m.OAuthIntegrations = s.OAuthIntegrations
	m.Limitations = s.Limitations

	return nil
}

package orders

import (
	"encoding/json"
	"time"
)

type OrderIntentData string

const (
	OrderIntentSale      OrderIntentData = "SALE"
	OrderIntentAuthorize OrderIntentData = "AUTHORIZE"
)

type DetailsData struct {
	Subtotal         string `json:"subtotal,omitempty"`
	Shipping         string `json:"shipping,omitempty"`
	Tax              string `json:"tax,omitempty"`
	HandlingFee      string `json:"handling_fee,omitempty"`
	ShippingDiscount string `json:"shipping_discount,omitempty"`
	Insurance        string `json:"insurance,omitempty"`
	GiftWrap         string `json:"gift_wrap,omitempty"`
}

type AmountData struct {
	Currency string      `json:"currency"`
	Total    string      `json:"total"`
	Details  DetailsData `json:"details"`
}

type DisplayPhoneData struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}

type PayeeDisplayMetadata struct {
	Email        string            `json:"email,omitempty"`
	DisplayPhone *DisplayPhoneData `json:"display_phone,omitempty"`
	BrandName    string            `json:"brand_name,omitempty"`
}

type PayeeData struct {
	Email                string                `json:"email,omitempty"`
	MerchantID           string                `json:"merchant_id,omitempty"`
	PayeeDisplayMetadata *PayeeDisplayMetadata `json:"payee_display_metadata"`
}

type ItemData struct {
	Sku         string `json:"sku,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Price       string `json:"price,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Tax         string `json:"tax,omitempty"`
	URL         string `json:"url,omitempty"`
}

type NormalizationStatusData string

const (
	NormalizationStatusUnknown                   NormalizationStatusData = "UNKNOWN"
	NormalizationStatusUnnormalizedUserPreferred NormalizationStatusData = "UNNORMALIZED_USER_PREFERRED"
	NormalizationStatusNormalized                NormalizationStatusData = "NORMALIZED"
	NormalizationStatusUnnormalized              NormalizationStatusData = "UNNORMALIZED"
)

type ShippingAddressData struct {
	Line1               string                  `json:"line1,omitempty"`
	Line2               string                  `json:"line2,omitempty"`
	City                string                  `json:"city,omitempty"`
	CountryCode         string                  `json:"country_code,omitempty"`
	PostalCode          string                  `json:"postal_code,omitempty"`
	State               string                  `json:"state,omitempty"`
	Phone               string                  `json:"phone,omitempty"`
	NormalizationStatus NormalizationStatusData `json:"normalization_status,omitempty"`
	Type                string                  `json:"type,omitempty"`
	RecipientName       string                  `json:"recipient_name,omitempty"`
}

type CurrencyData struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type PartnerFeeDetailsData struct {
	Receiver PayeeData    `json:"receiver"`
	Amount   CurrencyData `json:"amount"`
}

type NameValuePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MetadataData struct {
	PostbackData      []NameValuePair `json:"postback_data,omitempty"`
	SupplementaryData []NameValuePair `json:"supplementary_data,omitempty"`
}

type CaptureStatusData string

const (
	CaptureStatusPending           CaptureStatusData = "PENDING"
	CaptureStatusCompleted         CaptureStatusData = "COMPLETED"
	CaptureStatusRefunded          CaptureStatusData = "REFUNDED"
	CaptureStatusPartiallyRefunded CaptureStatusData = "PARTIALLY_REFUNDED"
)

type CaptureReasonCodeData string

const (
	CaptureReasonCodeChargeback                              CaptureReasonCodeData = "CHARGEBACK"
	CaptureReasonCodeGuarantee                               CaptureReasonCodeData = "GUARANTEE"
	CaptureReasonCodeBuyerComplaint                          CaptureReasonCodeData = "BUYER_COMPLAINT"
	CaptureReasonCodeRefund                                  CaptureReasonCodeData = "REFUND"
	CaptureReasonCodeUnconfirmedShippingAddress              CaptureReasonCodeData = "UNCONFIRMED_SHIPPING_ADDRESS"
	CaptureReasonCodeECheck                                  CaptureReasonCodeData = "ECHECK"
	CaptureReasonCodeInternationalWithdrawal                 CaptureReasonCodeData = "INTERNATIONAL_WITHDRAWAL"
	CaptureReasonCodeReceivingPreferenceMandatesManualAction CaptureReasonCodeData = "RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION"
	CaptureReasonCodePaymentReview                           CaptureReasonCodeData = "PAYMENT_REVIEW"
	CaptureReasonCodeRegulatoryReview                        CaptureReasonCodeData = "REGULATORY_REVIEW"
	CaptureReasonCodeUnilateral                              CaptureReasonCodeData = "UNILATERAL"
	CaptureReasonCodeVerificationRequired                    CaptureReasonCodeData = "VERIFICATION_REQUIRED"
	CaptureReasonCodeDelayedDisbursement                     CaptureReasonCodeData = "DELAYED_DISBURSEMENT"
)

type LinkData struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type CaptureData struct {
	ID             string                `json:"id,omitempty"`
	Amount         *AmountData           `json:"amount,omitempty"`
	Status         CaptureStatusData     `json:"status,omitempty"`
	ReasonCode     CaptureReasonCodeData `json:"reason_code,omitempty"`
	TransactionFee *CurrencyData         `json:"transaction_fee,omitempty"`
	Links          []LinkData            `json:"links,omitempty"`
}

type RefundStatusData string

const (
	RefundStatusPending   RefundStatusData = "PENDING"
	RefundStatusCompleted RefundStatusData = "COMPLETED"
	RefundStatusFailed    RefundStatusData = "FAILED"
)

type RefundData struct {
	ID        string           `json:"id,omitempty"`
	Amount    *AmountData      `json:"amount,omitempty"`
	CaptureID string           `json:"capture_id,omitempty"`
	SaleID    string           `json:"sale_id,omitempty"`
	Status    RefundStatusData `json:"status,omitempty"`
	Links     []LinkData       `json:"links,omitempty"`
}

type SaleStateData string

const (
	SaleStateCompleted         SaleStateData = "COMPLETED"
	SaleStatePartiallyRefunded SaleStateData = "PARTIALLY_REFUNDED"
	SaleStatePending           SaleStateData = "PENDING"
	SaleStateRefunded          SaleStateData = "REFUNDED"
	SaleStateDenied            SaleStateData = "DENIED"
)

type SaleData struct {
	ID             string        `json:"id,omitempty"`
	Amount         *AmountData   `json:"amount,omitempty"`
	TransactionFee *CurrencyData `json:"transaction_fee,omitempty"`
	State          SaleStateData `json:"state,omitempty"`
	CreateTime     time.Time     `json:"create_time,omitempty"`
	UpdateTime     time.Time     `json:"update_time,omitempty"`
	Links          []LinkData    `json:"links,omitempty"`
}

func (s *SaleData) MarshalJSON() ([]byte, error) {
	data := struct {
		ID             string        `json:"id,omitempty"`
		Amount         *AmountData   `json:"amount,omitempty"`
		TransactionFee *CurrencyData `json:"transaction_fee,omitempty"`
		State          SaleStateData `json:"state,omitempty"`
		CreateTime     string        `json:"create_time,omitempty"`
		UpdateTime     string        `json:"update_time,omitempty"`
		Links          []LinkData    `json:"links,omitempty"`
	}{
		ID:             s.ID,
		Amount:         s.Amount,
		TransactionFee: s.TransactionFee,
		State:          s.State,
		CreateTime:     s.CreateTime.Format(time.RFC3339Nano),
		UpdateTime:     s.UpdateTime.Format(time.RFC3339Nano),
		Links:          s.Links,
	}
	if s.CreateTime.IsZero() {
		data.CreateTime = ""
	}
	if s.UpdateTime.IsZero() {
		data.UpdateTime = ""
	}
	return json.Marshal(&data)
}

func (s *SaleData) UnmarshalJSON(b []byte) error {
	data := struct {
		ID             string        `json:"id,omitempty"`
		Amount         *AmountData   `json:"amount,omitempty"`
		TransactionFee *CurrencyData `json:"transaction_fee,omitempty"`
		State          SaleStateData `json:"state,omitempty"`
		CreateTime     string        `json:"create_time,omitempty"`
		UpdateTime     string        `json:"update_time,omitempty"`
		Links          []LinkData    `json:"links,omitempty"`
	}{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	s.ID = data.ID
	s.Amount = data.Amount
	s.TransactionFee = data.TransactionFee
	s.State = data.State
	if data.CreateTime != "" {
		s.CreateTime, err = time.Parse(time.RFC3339Nano, data.CreateTime)
		if err != nil {
			return err
		}
	} else {
		s.CreateTime = time.Time{}
	}
	if data.UpdateTime != "" {
		s.UpdateTime, err = time.Parse(time.RFC3339Nano, data.UpdateTime)
		if err != nil {
			return err
		}
	} else {
		s.UpdateTime = time.Time{}
	}
	s.Links = data.Links
	return nil
}

type PaymentSummaryData struct {
	Captures       []CaptureData `json:"captures,omitempty"`
	Refunds        []RefundData  `json:"refunds,omitempty"`
	Sales          []SaleData    `json:"sales,omitempty"`
	Authorizations []SaleData    `json:"authorizations,omitempty"`
}

type PurchaseStatusData string

const (
	PurchaseStatusNotProcessed PurchaseStatusData = "NOT_PROCESSED"
	PurchaseStatusPending      PurchaseStatusData = "PENDING"
	PurchaseStatusVoided       PurchaseStatusData = "VOIDED"
	PurchaseStatusAuthorized   PurchaseStatusData = "AUTHORIZED"
	PurchaseStatusCaptured     PurchaseStatusData = "CAPTURED"
)

type PurchaseUnitReasonCodeData string

const (
	PurchaseUnitReasonCodePayerShippingUnconfirmed PurchaseUnitReasonCodeData = "PAYER_SHIPPING_UNCONFIRMED"
	PurchaseUnitReasonCodeMultiCurrency            PurchaseUnitReasonCodeData = "MULTI_CURRENCY"
	PurchaseUnitReasonCodeRiskReview               PurchaseUnitReasonCodeData = "RISK_REVIEW"
	PurchaseUnitReasonCodeRegulatoryReview         PurchaseUnitReasonCodeData = "REGULATORY_REVIEW"
	PurchaseUnitReasonCodeVerificationRequired     PurchaseUnitReasonCodeData = "VERIFICATION_REQUIRED"
	PurchaseUnitReasonCodeOrder                    PurchaseUnitReasonCodeData = "ORDER"
	PurchaseUnitReasonCodeOther                    PurchaseUnitReasonCodeData = "OTHER"
	PurchaseUnitReasonCodeDeclinedByPolicy         PurchaseUnitReasonCodeData = "DECLINED_BY_POLICY"
)

type PurchaseUnitData struct {
	ReferenceID        string                     `json:"reference_id"`
	Amount             *AmountData                `json:"amount,omitempty"`
	Payee              *PayeeData                 `json:"payee,omitempty"`
	Description        string                     `json:"description,omitempty"`
	Custom             string                     `json:"custom,omitempty"`
	InvoiceNumber      string                     `json:"invoice_number,omitempty"`
	PaymentDescriptor  string                     `json:"payment_descriptor,omitempty"`
	Items              []ItemData                 `json:"items,omitempty"`
	NotifyURL          string                     `json:"notify_url,omitempty"`
	ShippingAddress    *ShippingAddressData       `json:"shipping_address,omitempty"`
	ShippingMethod     string                     `json:"shipping_method,omitempty"`
	PartnerFeeDetails  *PartnerFeeDetailsData     `json:"partner_fee_details,omitempty"`
	PaymentLinkedGroup int                        `json:"payment_linked_group,omitempty"`
	Metadata           *MetadataData              `json:"metadata,omitempty"`
	PaymentSummary     *PaymentSummaryData        `json:"payment_summary,omitempty"`
	Status             PurchaseStatusData         `json:"status,omitempty"`
	ReasonCode         PurchaseUnitReasonCodeData `json:"reason_code,omitempty"`
}

type ShippingPreferencesData string

const (
	ShippingPreferencesNoShipping         ShippingPreferencesData = "NO_SHIPPING"
	ShippingPreferencesGetFromFile        ShippingPreferencesData = "GET_FROM_FILE"
	ShippingPreferencesSetProvidedAddress ShippingPreferencesData = "SET_PROVIDED_ADDRESS"
)

type ApplicationContextData struct {
	BrandName           string                  `json:"brand_name,omitempty"`
	Locale              string                  `json:"locale,omitempty"`
	LandingPage         string                  `json:"landing_page,omitempty"`
	ShippingPreferences ShippingPreferencesData `json:"shipping_preferences,omitempty"`
	UserAction          string                  `json:"user_action,omitempty"`
	PostbackData        []NameValuePair         `json:"postback_data,omitempty"`
	SupplementaryData   []NameValuePair         `json:"supplementary_data,omitempty"`
}

type PhoneTypeData string

const (
	PhoneTypeHome   PhoneTypeData = "HOME"
	PhoneTypeWork   PhoneTypeData = "WORK"
	PhoneTypeMobile PhoneTypeData = "MOBILE"
	PhoneTypeOther  PhoneTypeData = "OTHER"
)

type TaxIDTypeData string

const (
	TaxIDTypeBRCPF  TaxIDTypeData = "BR_CPF"
	TaxIDTypeBRCNPJ TaxIDTypeData = "BR_CNPJ"
)

type AddressData struct {
	Line1               string                  `json:"line1,omitempty"`
	Line2               string                  `json:"line2,omitempty"`
	City                string                  `json:"city,omitempty"`
	CountryCode         string                  `json:"country_code,omitempty"`
	PostalCode          string                  `json:"postal_code,omitempty"`
	State               string                  `json:"state,omitempty"`
	Phone               string                  `json:"phone,omitempty"`
	NormalizationStatus NormalizationStatusData `json:"normalization_status,omitempty"`
	Type                string                  `json:"type,omitempty"`
}

type PayerInfoData struct {
	Email          string        `json:"email,omitempty"`
	Salutation     string        `json:"salutation,omitempty"`
	FirstName      string        `json:"first_name,omitempty"`
	MiddleName     string        `json:"middle_name,omitempty"`
	LastName       string        `json:"last_name,omitempty"`
	Suffix         string        `json:"suffix,omitempty"`
	PayerID        string        `json:"payer_id,omitempty"`
	Phone          string        `json:"phone,omitempty"`
	PhoneType      PhoneTypeData `json:"phone_type,omitempty"`
	BirthDate      time.Time     `json:"birth_date,omitempty"`
	TaxID          string        `json:"tax_id,omitempty"`
	TaxIDType      TaxIDTypeData `json:"tax_id_type,omitempty"`
	CountryCode    string        `json:"country_code,omitempty"`
	BillingAddress AddressData   `json:"billing_address,omitempty"`
}

func (p *PayerInfoData) MarshalJSON() ([]byte, error) {
	data := struct {
		Email          string        `json:"email,omitempty"`
		Salutation     string        `json:"salutation,omitempty"`
		FirstName      string        `json:"first_name,omitempty"`
		MiddleName     string        `json:"middle_name,omitempty"`
		LastName       string        `json:"last_name,omitempty"`
		Suffix         string        `json:"suffix,omitempty"`
		PayerID        string        `json:"payer_id,omitempty"`
		Phone          string        `json:"phone,omitempty"`
		PhoneType      PhoneTypeData `json:"phone_type,omitempty"`
		BirthDate      string        `json:"birth_date,omitempty"`
		TaxID          string        `json:"tax_id,omitempty"`
		TaxIDType      TaxIDTypeData `json:"tax_id_type,omitempty"`
		CountryCode    string        `json:"country_code,omitempty"`
		BillingAddress AddressData   `json:"billing_address,omitempty"`
	}{
		Email:          p.Email,
		Salutation:     p.Salutation,
		FirstName:      p.FirstName,
		MiddleName:     p.MiddleName,
		LastName:       p.LastName,
		Suffix:         p.Suffix,
		PayerID:        p.PayerID,
		Phone:          p.Phone,
		PhoneType:      p.PhoneType,
		BirthDate:      p.BirthDate.Format("2006-01-02"),
		TaxID:          p.TaxID,
		TaxIDType:      p.TaxIDType,
		CountryCode:    p.CountryCode,
		BillingAddress: p.BillingAddress,
	}
	if p.BirthDate.IsZero() {
		data.BirthDate = ""
	}
	return json.Marshal(&data)
}

func (p *PayerInfoData) UnmarshalJSON(b []byte) error {
	data := struct {
		Email          string        `json:"email,omitempty"`
		Salutation     string        `json:"salutation,omitempty"`
		FirstName      string        `json:"first_name,omitempty"`
		MiddleName     string        `json:"middle_name,omitempty"`
		LastName       string        `json:"last_name,omitempty"`
		Suffix         string        `json:"suffix,omitempty"`
		PayerID        string        `json:"payer_id,omitempty"`
		Phone          string        `json:"phone,omitempty"`
		PhoneType      PhoneTypeData `json:"phone_type,omitempty"`
		BirthDate      string        `json:"birth_date,omitempty"`
		TaxID          string        `json:"tax_id,omitempty"`
		TaxIDType      TaxIDTypeData `json:"tax_id_type,omitempty"`
		CountryCode    string        `json:"country_code,omitempty"`
		BillingAddress AddressData   `json:"billing_address,omitempty"`
	}{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	p.Email = data.Email
	p.Salutation = data.Salutation
	p.FirstName = data.FirstName
	p.MiddleName = data.MiddleName
	p.LastName = data.LastName
	p.Suffix = data.Suffix
	p.PayerID = data.PayerID
	p.Phone = data.Phone
	p.PhoneType = data.PhoneType
	if data.BirthDate != "" {
		p.BirthDate, err = time.Parse(time.RFC3339Nano, data.BirthDate)
		if err != nil {
			return nil
		}
	} else {
		p.BirthDate = time.Time{}
	}
	p.TaxID = data.TaxID
	p.TaxIDType = data.TaxIDType
	p.CountryCode = data.CountryCode
	p.BillingAddress = data.BillingAddress
	return nil
}

type RedirectURLsData struct {
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

type CreateOrderParams struct {
	Intent             OrderIntentData         `json:"intent,omitempty"`
	PurchaseUnits      []PurchaseUnitData      `json:"purchase_units"`
	ApplicationContext *ApplicationContextData `json:"application_context,omitempty"`
	PayerInfo          *PayerInfoData          `json:"payer_info,omitempty"`
	RedirectURLs       *RedirectURLsData       `json:"redirect_urls"`
}

type DisbursementModeData string

const (
	DisbursementModeInstant DisbursementModeData = "INSTANT"
	DisbursementModeDelayed DisbursementModeData = "DELAYED"
)

type PaymentDetailsData struct {
	PaymentID        string               `json:"payment_id"`
	DisbursementMode DisbursementModeData `json:"disbursement_mode"`
}

type OrderStatusData string

const (
	OrderStatusCreated   OrderStatusData = "CREATED"
	OrderStatusApproved  OrderStatusData = "APPROVED"
	OrderStatusCompleted OrderStatusData = "COMPLETED"
	OrderStatusFailed    OrderStatusData = "FAILED"
)

type CreateOrderResponse struct {
	ID                 string                  `json:"id"`
	Intent             OrderIntentData         `json:"intent"`
	PurchaseUnits      []PurchaseUnitData      `json:"purchase_units"`
	PaymentDetails     *PaymentDetailsData     `json:"payment_details"`
	ApplicationContext *ApplicationContextData `json:"application_context"`
	PayerInfo          *PayerInfoData          `json:"payer_info"`
	Metadata           *MetadataData           `json:"metadata"`
	Status             OrderStatusData         `json:"status"`
	RedirectURLs       *RedirectURLsData       `json:"redirect_urls"`
	CreateTime         time.Time               `json:"create_time"`
	UpdateTime         time.Time               `json:"update_time"`
	Links              []LinkData              `json:"links"`
}

func (c *CreateOrderResponse) UnmarshalJSON(b []byte) error {
	data := struct {
		ID                 string                  `json:"id"`
		Intent             OrderIntentData         `json:"intent"`
		PurchaseUnits      []PurchaseUnitData      `json:"purchase_units"`
		PaymentDetails     *PaymentDetailsData     `json:"payment_details"`
		ApplicationContext *ApplicationContextData `json:"application_context"`
		PayerInfo          *PayerInfoData          `json:"payer_info"`
		Metadata           *MetadataData           `json:"metadata"`
		Status             OrderStatusData         `json:"status"`
		RedirectURLs       *RedirectURLsData       `json:"redirect_urls"`
		CreateTime         string                  `json:"create_time"`
		UpdateTime         string                  `json:"update_time"`
		Links              []LinkData              `json:"links"`
	}{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	c.ID = data.ID
	c.Intent = data.Intent
	c.PurchaseUnits = data.PurchaseUnits
	c.PaymentDetails = data.PaymentDetails
	c.ApplicationContext = data.ApplicationContext
	c.PayerInfo = data.PayerInfo
	c.Metadata = data.Metadata
	c.Status = data.Status
	c.RedirectURLs = data.RedirectURLs
	if data.CreateTime != "" {
		c.CreateTime, err = time.Parse(time.RFC3339Nano, data.CreateTime)
		if err != nil {
			return err
		}
	} else {
		c.CreateTime = time.Time{}
	}
	if data.UpdateTime != "" {
		c.UpdateTime, err = time.Parse(time.RFC3339Nano, data.UpdateTime)
		if err != nil {
			return err
		}
	} else {
		c.UpdateTime = time.Time{}
	}
	c.Links = data.Links
	return nil
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

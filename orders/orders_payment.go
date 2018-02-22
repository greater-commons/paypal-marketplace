package orders

import "time"

type PayOrderResponse struct {
	OrderID        string              `json:"order_id"`
	Status         OrderStatusData     `json:"status"`
	Intent         OrderIntentData     `json:"intent"`
	PayerInfo      *PayerInfoData      `json:"payer_info"`
	PurchaseUnits  []PurchaseUnitData  `json:"purchase_units"`
	CreateTime     time.Time           `json:"create_time"`
	UpdateTime     time.Time           `json:"update_time"`
	ID             string              `json:"id,omitempty"`
	Links          []LinkData          `json:"links"`
	PaymentDetails *PaymentDetailsData `json:"payment_details"`
}

type ResponsePreferenceData string

const (
	ResponsePreferenceAsync ResponsePreferenceData = "respond-async"
	ResponsePreferenceSync  ResponsePreferenceData = "respond-sync"
)

type ProcessingStateData struct {
	Status string `json:"status"`
}

type DisbursementCurrencyData struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type FinalizeDisbursementResponse struct {
	ItemID              string                    `json:"item_id"`
	ProcessingState     *ProcessingStateData      `json:"processing_state"`
	ReferenceID         string                    `json:"reference_id"`
	ReferenceType       string                    `json:"reference_type"`
	PayoutTransactionID string                    `json:"payout_transaction_id"`
	ExternalReferenceID string                    `json:"external_reference_id"`
	PayoutAmount        *DisbursementCurrencyData `json:"payout_amount"`
	PayoutDestination   string                    `json:"payout_destination"`
	Links               []LinkData                `json:"links"`
}

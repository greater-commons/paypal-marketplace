package orders

import "time"

type PayOrderResponse struct {
	OrderID       string             `json:"order_id"`
	Status        OrderStatusData    `json:"status"`
	Intent        OrderIntentData    `json:"intent"`
	PayerInfo     *PayerInfoData     `json:"payer_info"`
	PurchaseUnits []PurchaseUnitData `json:"purchase_units"`
	CreateTime    time.Time          `json:"create_time"`
	UpdateTime    time.Time          `json:"update_time"`
	Links         []LinkData         `json:"links"`
}

type ResponsePreferenceData string

const (
	ResponsePreferenceAsync ResponsePreferenceData = "respond-async"
	ResponsePreferenceSync  ResponsePreferenceData = "respond-sync"
)

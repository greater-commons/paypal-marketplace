package orders

type RequestRefundParams struct {
	Amount        AmountData `json:"amount"`
	Custom        string     `json:"custom"`
	InvoiceNumber string     `json:"invoice_number"`
}

type RequestRefundResponse struct {
	ID                       string     `json:"id"`
	State                    string     `json:"state"`
	Amount                   AmountData `json:"amount"`
	RefundFromTransactionFee AmountData `json:"refund_from_transaction_fee"`
	TotalRefundedAmount      AmountData `json:"total_refunded_amount"`
	RefundFromReceivedAmount AmountData `json:"refund_from_received_amount"`
	CaptureID                string     `json:"capture_id"`
	InvoiceNumber            string     `json:"invoice_number"`
	Links                    LinkData   `json:"links"`
}

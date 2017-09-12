package market

import (
	"context"
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
	"testing"

	"github.com/greater-commons/paypal-marketplace/orders"
)

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()

	c := NewClient(ctx, GetTestClientID(), GetTestSecret(), Sandbox)
	c.BNCode = GetTestBNCode()

	num, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	trackingID := strconv.FormatInt(num.Int64(), 10)
	err = c.SaveTransactionContext(ctx, GetTestPayerID(), trackingID, nil)
	if err != nil {
		t.Fatal("Error attempting to get a transaction context")
	}
	t.Log("Tracking ID is:", trackingID)
	resp, err := c.CreateOrder(ctx, trackingID, &orders.CreateOrderParams{
		Intent: orders.OrderIntentSale,
		PurchaseUnits: []orders.PurchaseUnitData{
			{
				ReferenceID: "abc",
				Amount: &orders.AmountData{
					Currency: "USD",
					Details: orders.DetailsData{
						Subtotal: "20",
					},
					Total: "20",
				},
				Items: []orders.ItemData{
					{
						Name:     "Test Item",
						Quantity: 1,
						Price:    "20",
						Currency: "USD",
					},
				},
			},
		},
		RedirectURLs: &orders.RedirectURLsData{
			ReturnURL: "http://localhost:8080/return",
			CancelURL: "http://localhost:8080/cancel",
		},
	})
	if err != nil {
		t.Fatal("Error attempting to create an order:", err)
	}
	t.Logf("Order: %+v\n", resp)
}

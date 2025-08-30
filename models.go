package zeno

import (
	"bytes"
	"encoding/json"
	"log"
	"time"
)

// newData constructs and returns zenno.Data
func newData(paymentData PaymentData) (data *bytes.Buffer) {

	// construct and return data
	orderData := map[string]any{
		"order_id":    paymentData.OrderID,
		"buyer_email": paymentData.Email,
		"buyer_name":  paymentData.Name,
		"buyer_phone": paymentData.Phone,
		"amount":      paymentData.Amount,
		"metadata":    paymentData.MetaData,
		"webhook_url": apiConfigOptions.CallbackURL,
	}

	log.Printf("orderData: %v\n", orderData)

	// convert data to json
	jsonData, err := json.Marshal(orderData)
	if err != nil {
		zLog(err.Error())
		return data
	}

	return bytes.NewBuffer(jsonData)
}

// zenoRes holds data about status json data returned by the Zeno API
type zenoRes struct {
	Status     string `json:"status"`
	ResultCode string `json:"resultcode"`
	Message    string `json:"message"`
	OrderID    string `json:"order_id"`
}

// Holds API key for the ZenoAPI, timeout and the callback endpoint
type Options struct {
	APIKey      string
	CallbackURL string
	Timeout     time.Duration
}

// HookData holds data passed to the hook request
type HookData struct {
	OrderID       string `json:"order_id"`
	PaymentStatus string `json:"payment_status"`
	Reference     string `json:"reference"`
	MetaData      any    `json:"metadata"`
}

// PaymentData holds data passed to the payment request
type PaymentData struct {
	OrderID  string
	Amount   float64
	Name     string
	Phone    string
	Email    string
	MetaData any
}

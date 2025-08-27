package zeno

import (
	"bytes"
	"encoding/json"
	"time"
)

// newData constructs and returns zenno.Data
func newData(orderID string, amount float64, name, phone, email string) (data *bytes.Buffer) {

	// construct and return data
	orderData := map[string]any{
		"order_id":    orderID,
		"buyer_email": email,
		"buyer_name":  name,
		"buyer_phone": phone,
		"amount":      amount,
		"webhook_url": apiConfigOptions.CallbackURL,
	}

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

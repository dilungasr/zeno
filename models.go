package zeno

import (
	"bytes"
	"net/url"
)

// newData constructs and returns zenno.Data
func newData(amount string, name, phone, email string) *bytes.Buffer {

	// get API access data
	accountID := apiConfigData.AccountID
	APIKey := apiConfigData.APIKey
	secreteKey := apiConfigData.SecreteKey

	// construct and return data
	values := url.Values{}

	values.Set("create_order", "1")
	values.Set("buyer_name", name)
	values.Set("buyer_phone", phone)
	values.Set("buyer_email", email)
	values.Set("amount", amount)
	values.Set("account_id", accountID)
	values.Set("secrete_key", secreteKey)
	values.Set("api_key", APIKey)

	data := values.Encode()

	return bytes.NewBufferString(data)
}

// zenoRes holds data about status json data returned by the Zeno API
type zenoRes struct {
	Status        string `json:"status"`
	PaymentStatus string `json:"payment_status"`
	Message       string `json:"message"`
	OrderID       string `json:"order_id"`
}

// Holds API access data for the ZenoAPI
type apiConfig struct {
	AccountID  string
	SecreteKey string
	APIKey     string
}

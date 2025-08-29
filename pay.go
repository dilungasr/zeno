package zeno

import (
	"encoding/json"
	"errors"
	"net/http"
)

// API URL
const (
	zenoURL        string = "https://zenoapi.com/api/payments"
	mobileMoneyURL string = zenoURL + "/mobile_money_tanzania"
)

// Pay makes payment request to the Zeno API
func Pay(paymentData PaymentData, timeoutFn func(orderID string)) (err error) {

	// construct data in buffer of url.Values
	data := newData(paymentData)

	// prepare and make request
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, mobileMoneyURL, data)
	if err != nil {
		msg := zLog(err.Error())
		return errors.New(msg)
	}

	// add format and api key
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiConfigOptions.APIKey)

	res, err := client.Do(req)
	if err != nil {
		msg := zLog(err.Error())
		return errors.New(msg)

	}

	if res != nil {
		defer res.Body.Close()
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		msg := zLog("Unexpected status code %v", res.StatusCode)
		return errors.New(msg)
	}

	//  decode json data
	var zRes zenoRes

	err = json.NewDecoder(res.Body).Decode(&zRes)
	if err != nil {
		msg := zLog(err.Error())
		return errors.New(msg)
	}

	if zRes.Status != "success" {
		msg := zLog(zRes.Message)
		return errors.New(msg)
	}

	// monitor timeout
	if timeoutFn != nil {
		go timeoutStatus(paymentData.OrderID, timeoutFn)
	}

	zLog("Payment for order %v initiated...", zRes.OrderID)

	return nil
}

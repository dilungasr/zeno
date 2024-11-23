package zeno

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// API URL
var zenoURL string = "https://api.zeno.africa"

// Pay makes payment request to the Zeno API
func Pay(amount, name, phone, email string, callback func(orderID string, ok bool)) (orderID string, err error) {
	// construct data in buffer of url.Values
	data := newData(amount, name, phone, email)

	// prepare and make request
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, zenoURL, data)
	if err != nil {
		msg := zLog(err.Error())
		return orderID, fmt.Errorf(msg)
	}

	// add required headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		msg := zLog(err.Error())
		return orderID, fmt.Errorf(msg)

	}

	if res != nil {
		defer res.Body.Close()
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		msg := zLog("Unexpected status code %v", res.StatusCode)
		return orderID, fmt.Errorf(msg)
	}

	//  decode json data
	var zRes zenoRes

	err = json.NewDecoder(res.Body).Decode(&zRes)
	if err != nil {
		msg := zLog(err.Error())
		return orderID, fmt.Errorf(msg)
	}

	if zRes.Status != "success" {
		msg := zLog(zRes.Message)
		return orderID, fmt.Errorf(msg)
	}

	// start background polling
	go pollPaymentStatus(zRes.OrderID, callback)

	zLog("Payment for order %v initiated...", zRes.OrderID)

	return zRes.OrderID, nil
}

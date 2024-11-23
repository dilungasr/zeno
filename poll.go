package zeno

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// pollPaymentStatus periodically checks the status of the order payment every 5 seconds for 50 seconds max
func pollPaymentStatus(orderID string, callback func(orderID string, ok bool)) {
	ticker := time.NewTicker(5 * time.Second)
	timeout := time.After(50 * time.Second)

	// loop the channels
	for {
		select {
		case <-ticker.C:
			err := checkPaymentStatus(orderID)

			if err != nil {
				continue
			} else {
				// completed
				callback(orderID, true)
				ticker.Stop()
				return
			}

		case <-timeout:
			zLog("Payment for order %v failed", orderID)
			// update the database to FAILED
			callback(orderID, false)
			ticker.Stop()
			return
		}
	}
}

// checkPaymentStatus makes the request to the payment gateway API to check the status of the order
func checkPaymentStatus(orderID string) (err error) {
	statusURL := zenoURL + "/order-status"

	// get API access data
	APIKey := apiConfigData.APIKey
	secreteKey := apiConfigData.SecreteKey

	values := url.Values{}

	values.Set("check_status", "1")
	values.Set("order_id", orderID)
	values.Set("api_key", APIKey)
	values.Set("secrete_key", secreteKey)

	data := values.Encode()

	req, err := http.NewRequest(http.MethodPost, statusURL, bytes.NewBufferString(data))
	if err != nil {
		msg := zLog(err.Error())
		return fmt.Errorf(msg)
	}

	// add required headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		msg := zLog(err.Error())
		return fmt.Errorf(msg)
	}

	if res != nil {
		defer res.Body.Close()
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		msg := zLog("Unexpected status code %v", res.StatusCode)
		return fmt.Errorf(msg)
	}

	//  decode json data
	var statusData zenoRes

	err = json.NewDecoder(res.Body).Decode(&statusData)
	if err != nil {
		msg := zLog(err.Error())
		return fmt.Errorf(msg)
	}

	if statusData.PaymentStatus != "COMPLETED" {
		return fmt.Errorf("[zeno-payment] PAYMENT INCOMPLETE: %v", statusData.Message)
	}

	return err

}

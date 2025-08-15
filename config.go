package zeno

import "time"

var apiConfigData apiConfig

func ConfigAPI(accountID, APIKey, secreteKey string, timeout ...time.Duration) {
	var timeoutData time.Duration
	if len(timeout) > 0 {
		timeoutData = timeout[0]
	} else {
		timeoutData = time.Second * 50
	}

	// set API access data
	apiConfigData = apiConfig{
		AccountID:  accountID,
		APIKey:     APIKey,
		SecreteKey: secreteKey,
		Timeout:    timeoutData,
	}
}

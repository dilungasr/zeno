package zeno

var apiConfigData apiConfig

func ConfigAPI(accountID, APIKey, secreteKey string) {
	apiConfigData = apiConfig{
		AccountID:  accountID,
		APIKey:     APIKey,
		SecreteKey: secreteKey,
	}
}

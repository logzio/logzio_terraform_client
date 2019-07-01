package test_utils

import (
	"os"
)

const ENV_LOGZIO_API_TOKEN string = "LOGZIO_API_TOKEN"
const ENV_LOGZIO_ACCOUNT_ID string = "LOGZIO_ACCOUNT_ID"

func GetApiToken() string {
	api_token := os.Getenv(ENV_LOGZIO_API_TOKEN)
	return api_token
}

func GetAccountId() string {
	account_id := os.Getenv(ENV_LOGZIO_ACCOUNT_ID);
	return account_id
}
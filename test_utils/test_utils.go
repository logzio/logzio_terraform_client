package test_utils

import (
	"os"
	"strconv"
)

const ENV_LOGZIO_API_TOKEN string = "LOGZIO_API_TOKEN"
const ENV_LOGZIO_ACCOUNT_ID string = "LOGZIO_ACCOUNT_ID"

func GetApiToken() string {
	api_token := os.Getenv(ENV_LOGZIO_API_TOKEN)
	return api_token
}

func GetAccountId() int32 {
	account_id_string := os.Getenv(ENV_LOGZIO_ACCOUNT_ID);
	account_id, _ := strconv.ParseInt(account_id_string, 10, 32)
	return int32(account_id)
}


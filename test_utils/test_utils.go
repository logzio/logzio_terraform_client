package test_utils

import (
	"fmt"
	"os"
	"strconv"
)

const ENV_LOGZIO_API_TOKEN string = "LOGZIO_API_TOKEN"
const ENV_LOGZIO_ACCOUNT_ID string = "LOGZIO_ACCOUNT_ID"

func GetApiToken() (string, error) {
	api_token := os.Getenv(ENV_LOGZIO_API_TOKEN)
	if len(api_token) > 0 {
		return api_token, nil
	}
	return "", fmt.Errorf("%s env var not specified", ENV_LOGZIO_API_TOKEN)
}

func GetAccountId() (int32, error) {
	account_id_string := os.Getenv(ENV_LOGZIO_ACCOUNT_ID)
	if len(account_id_string) == 0 {
		return -1, fmt.Errorf("%s env var not specified", ENV_LOGZIO_ACCOUNT_ID)
	}
	account_id, _ := strconv.ParseInt(account_id_string, 10, 32)
	return int32(account_id), nil
}

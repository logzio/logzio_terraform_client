package test_utils

import (
	"fmt"
	"os"
	"strconv"
)

const ENV_LOGZIO_BASE_URL = "LOGZIO_BASE_URL"
const ENV_LOGZIO_API_TOKEN string = "LOGZIO_API_TOKEN"
const ENV_LOGZIO_ACCOUNT_ID string = "LOGZIO_ACCOUNT_ID"
const LOGZIO_BASE_URL string = "https://api.logz.io"

func GetApiToken() (string, error) {
	api_token := os.Getenv(ENV_LOGZIO_API_TOKEN)
	if len(api_token) > 0 {
		return api_token, nil
	}
	return "", fmt.Errorf("%s env var not specified", ENV_LOGZIO_API_TOKEN)
}

func GetAccountId() (int64, error) {
	account_id_string := os.Getenv(ENV_LOGZIO_ACCOUNT_ID)
	if len(account_id_string) == 0 {
		return -1, fmt.Errorf("%s env var not specified", ENV_LOGZIO_ACCOUNT_ID)
	}
	account_id, _ := strconv.ParseInt(account_id_string, 10, 32)
	return int64(account_id), nil
}

func GetLogzIoBaseUrl() string {
	if len(os.Getenv(ENV_LOGZIO_BASE_URL)) > 0 {
		return os.Getenv(ENV_LOGZIO_BASE_URL)
	}
	return LOGZIO_BASE_URL
}

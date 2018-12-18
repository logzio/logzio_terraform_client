package test_utils

import (
	"os"
)

const ENV_LOGZIO_API_TOKEN string = "LOGZIO_API_TOKEN"

func GetApiToken() string {
	api_token := os.Getenv(ENV_LOGZIO_API_TOKEN)
	return api_token
}

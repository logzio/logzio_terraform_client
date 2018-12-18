package client

import (
	"log"
	"os"
)

type Client struct {
	ApiToken string
	log      log.Logger
}

const ENV_LOGZIO_BASE_URL = "LOGZIO_BASE_URL"
const LOGZIO_BASE_URL string = "https://api.logz.io"

var logzioBaseUrl string = LOGZIO_BASE_URL

func GetLogzioBaseUrl() string {
	if len(os.Getenv(ENV_LOGZIO_BASE_URL)) > 0 {
		logzioBaseUrl = os.Getenv(ENV_LOGZIO_BASE_URL)
	}
	return logzioBaseUrl
}

func New(apiToken string) *Client {
	var c Client
	c.ApiToken = apiToken
	return &c
}

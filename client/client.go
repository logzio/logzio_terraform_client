package client

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

const ENV_LOGZIO_BASE_URL = "LOGZIO_BASE_URL"
const LOGZIO_BASE_URL string = "https://api.logz.io"

const (
	ERROR_CODE    = "errorCode"
	ERROR_MESSAGE = "errorMessage"
)

type Client struct {
	ApiToken  string
	BaseUrl   string
	log       log.Logger
}

var logzioBaseUrl string = LOGZIO_BASE_URL

// Entry point into the logz.io client
func New(apiToken string) *Client {
	var c Client
	c.ApiToken = apiToken
	return &c
}

func GetLogzioBaseUrl() string {
	if len(os.Getenv(ENV_LOGZIO_BASE_URL)) > 0 {
		logzioBaseUrl = os.Getenv(ENV_LOGZIO_BASE_URL)
	}
	return logzioBaseUrl
}

func GetHttpClient(req *http.Request) *http.Client {
	url, err := http.ProxyFromEnvironment(req)
	if url != nil && err == nil {
		tr := &http.Transport{
			Proxy:           http.ProxyURL(url),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		return &http.Client{Transport: tr}
	} else {
		return &http.Client{}
	}
}

func IsErrorResponse(response map[string]interface{}) (bool, string) {
	if _, ok := response[ERROR_CODE]; ok {
		return true, response[ERROR_CODE].(string)
	}
	if _, ok := response[ERROR_MESSAGE]; ok {
		return true, response[ERROR_MESSAGE].(string)
	}
	return false, ""
}

func (c *Client) SetBaseUrl(BaseUrl string) {
	c.BaseUrl = BaseUrl
}

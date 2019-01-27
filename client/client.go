package client

import (
	"log"
	"net/http"
	"os"
)

type Client struct {
	ApiToken string
	BaseUrl  string
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

func (c *Client) SetBaseUrl(BaseUrl string) {
	c.BaseUrl = BaseUrl
}

func GetHttpClient(req *http.Request) *http.Client {
	url, _ := http.ProxyFromEnvironment(req)
	tr := &http.Transport{
		Proxy: http.ProxyURL(url),
	}
	httpClient := &http.Client{Transport: tr}
	return httpClient
}

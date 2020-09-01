package log_shipping

import (
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	serviceEndpoint = "%s/v1/log-shipping"
)

const (
	fldId       string = "id"
	fldLogsType string = "logsType"
)

type LogShippingClient struct {
	*client.Client
}

// Creates a new entry point into the users functions, accepts the user's logz.io API token and account Id
func New(apiToken, baseUrl string) (*LogShippingClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &LogShippingClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

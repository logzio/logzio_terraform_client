package kibana_objects

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	kibanaObjectsExportServiceEndpoint = "%s/v1/kibana/export"
	kibanaObjectsImportServiceEndpoint = "%s/v1/kibana/import"
	loggerName                         = "logzio-client"
)

type KibanaObjectsClient struct {
	*client.Client
	logger hclog.Logger
}

// Creates a new entry point into the kibana objects functions, accepts the
// user's logz.io API token and API base URL
func New(apiToken string, baseUrl string) (*KibanaObjectsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	client := &KibanaObjectsClient{
		Client: client.New(apiToken, baseUrl),
		logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Debug,
			Name:       loggerName,
			JSONFormat: true,
		}),
	}
	return client, nil
}

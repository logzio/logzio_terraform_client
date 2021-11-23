package grafana_objects

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaObjectsDashboardsByUID        = "%s/v1/grafana/api/dashboards/uid/%s"
	grafanaObjectsDashboardsCreateUpdate = "%s/v1/grafana/api/dashboards/db"
	loggerName                           = "logzio-client"
)

type GrafanaObjectsClient struct {
	*client.Client
	logger hclog.Logger
}

// Creates a new entry point into the grafana objects functions, accepts the
// user's logz.io API token and API base URL
func New(apiToken string, baseUrl string) (*GrafanaObjectsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	client := &GrafanaObjectsClient{
		Client: client.New(apiToken, baseUrl),
		logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Debug,
			Name:       loggerName,
			JSONFormat: true,
		}),
	}
	return client, nil
}

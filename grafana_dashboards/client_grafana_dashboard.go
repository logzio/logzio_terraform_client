package grafana_dashboards

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaObjectServiceEndpoint = "%s/v1/grafana/api/dashboards"
	loggerName                   = "logzio-client"

	dashboardCreateUpdate = "CreateUpdateGrafanaDashboard"
	dashboardDelete       = "DeleteGrafanaDashboard"
	dashboardGet          = "GetGrafanaDashboard"

	dashboardResourceName = "grafana dashboard"

	GrafanaSuccessStatus = "success"
)

type GrafanaObjectsClient struct {
	*client.Client
	logger hclog.Logger
}

type CreateUpdatePayload struct {
	Dashboard map[string]interface{} `json:"dashboard"`
	FolderId  float64                `json:"folderId,omitempty"`
	FolderUid string                 `json:"folderUid"`
	Message   string                 `json:"message"`
	Overwrite bool                   `json:"overwrite"`
}

type GetResults struct {
	Dashboard map[string]interface{} `json:"dashboard"`
	Meta      map[string]interface{} `json:"meta"`
}

type DeleteResults struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Id      int    `json:"id"`
}

type CreateUpdateResults struct {
	Id      int    `json:"id"`
	Uid     string `json:"uid"`
	Status  string `json:"status"`
	Version int    `json:"version"`
	Url     string `json:"url"`
	Slug    string `json:"slug"`
}

// New Creates a new entry point into the grafana objects functions, accepts the
// user's logz.io API token and API base URL
func New(apiToken string, baseUrl string) (*GrafanaObjectsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	grafanaDashboardClient := &GrafanaObjectsClient{
		Client: client.New(apiToken, baseUrl),
		logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Debug,
			Name:       loggerName,
			JSONFormat: true,
		}),
	}
	return grafanaDashboardClient, nil
}

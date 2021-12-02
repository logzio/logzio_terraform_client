package grafana_objects

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaObjectServiceEndpoint = "%s/v1/grafana/api/dashboards"
	loggerName                   = "logzio-client"
)

type GrafanaObjectsClient struct {
	*client.Client
	logger hclog.Logger
}

type CreateUpdatePayload struct {
	Dashboard DashboardObject `json:"dashboard"`
	FolderId  int             `json:"folderId"`
	FolderUid int             `json:"folderUid"`
	Message   string          `json:"message"`
	Overwrite bool            `json:"overwrite"`
}

type GetResults struct {
	Dashboard DashboardObject `json:"dashboard"`
	Meta      DashboardMeta   `json:"meta"`
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

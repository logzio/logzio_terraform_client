package grafana_folders

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaFolderServiceEndpoint string = "%s/v1/grafana/api/folders"

	grafanaFolderResourceName = "grafana folder"

	operationCreateGrafanaFolder = "CreateGrafanaFolder"
	operationGetGrafanaFolder    = "GetGrafanaFolder"
	operationUpdateGrafanaFolder = "UpdateGrafanaFolder"
	operationDeleteGrafanaFolder = "DeleteGrafanaFolder"
	operationListGrafanaFolders  = "ListGrafanaFolders"
)

type GrafanaFolderClient struct {
	*client.Client
}

type GrafanaFolder struct {
	Uid     string `json:"uid"`
	Title   string `json:"title"`
	Id      int64  `json:"id"`
	Url     string `json:"url"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Version int64  `json:"version"`
}

type CreateUpdateFolder struct {
	Uid       string `json:"uid"`
	Title     string `json:"title"`
	Overwrite bool   `json:"overwrite"`
}

func New(apiToken string, baseUrl string) (*GrafanaFolderClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	grafanaFolderClient := &GrafanaFolderClient{
		Client: client.New(apiToken, baseUrl),
	}

	return grafanaFolderClient, nil
}

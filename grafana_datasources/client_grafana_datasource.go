package grafana_datasources

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaDatasourceServiceEndpoint = "%s/v1/grafana/api/datasources"

	grafanaDatasourceResourceName = "grafana datasource"

	GetAllForAccountGrafanaDatasourceMethod = "GetAllForAccountDatasource"
)

type GrafanaDatasourceClient struct {
	*client.Client
}

type GrafanaDataSource struct {
	Id              int64                  `json:"id,omitempty"`
	Uid             string                 `json:"uid,omitempty"`
	Name            string                 `json:"name"`
	Type            string                 `json:"type"`
	TypeLogoUrl     string                 `json:"typeLogoUrl,omitempty"` // Only returned by the API. Depends on the Type.
	Url             string                 `json:"url"`
	Access          string                 `json:"access"`
	ReadOnly        bool                   `json:"readOnly"` // Only returned by the API. Can be set through the `editable` attribute of provisioned data sources.
	Database        string                 `json:"database,omitempty"`
	User            string                 `json:"user,omitempty"`
	OrgId           int64                  `json:"orgId,omitempty"`
	IsDefault       bool                   `json:"isDefault"`
	BasicAuth       bool                   `json:"basicAuth"`
	BasicAuthUser   string                 `json:"basicAuthUser,omitempty"`
	WithCredentials bool                   `json:"withCredentials,omitempty"`
	JsonData        map[string]interface{} `json:"jsonData,omitempty"`
	SecureJsonData  map[string]interface{} `json:"secureJsonData,omitempty"`
	Version         int                    `json:"version,omitempty"`
}

func New(apiToken string, baseUrl string) (*GrafanaDatasourceClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	grafanaDatasourceClient := &GrafanaDatasourceClient{
		Client: client.New(apiToken, baseUrl),
	}

	return grafanaDatasourceClient, nil
}

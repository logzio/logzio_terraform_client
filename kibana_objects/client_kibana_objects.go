package kibana_objects

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

type ExportType string

// Enums for exportType
const (
	kibanaObjectsServiceEndpoint = "%s/v1/kibana"
	loggerName                   = "logzio-client"

	ExportTypeSearch        ExportType = "search"
	ExportTypeVisualization ExportType = "visualization"
	ExportTypeDashboard     ExportType = "dashboard"

	exportKibanaObjectOperation = "ExportKibanaObject"
	importKibanaObjectOperation = "ImportKibanaObject"

	kibanaObjectResourceName = "kibana object"
)

type KibanaObjectsClient struct {
	*client.Client
	logger hclog.Logger
}

type KibanaObjectExportRequest struct {
	Type ExportType `json:"type"` // Required
}

type KibanaObjectExportResponse struct {
	KibanaVersion string        `json:"kibanaVersion"`
	Hits          []interface{} `json:"hits"`
}

type KibanaObjectImportRequest struct {
	KibanaVersion string                   `json:"kibanaVersion"`
	Override      *bool                    `json:"override,omitempty"`
	Hits          []map[string]interface{} `json:"hits"`
}

type KibanaObjectImportResponse struct {
	Created []string `json:"created"`
	Updated []string `json:"updated,omitempty"`
	Ignored []string `json:"ignored,omitempty"`
	Failed  []string `json:"failed,omitempty"`
}

func New(apiToken, baseUrl string) (*KibanaObjectsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &KibanaObjectsClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

func (s ExportType) String() string {
	return string(s)
}

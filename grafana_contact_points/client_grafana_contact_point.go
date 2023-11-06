package grafana_contact_points

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaContactPointServiceEndpoint = "%s/v1/grafana/api/v1/provisioning/contact-points"

	operationCreateGrafanaContactPoint = "CreateGrafanaContactPoint"

	grafanaContactPointResourceName = "grafana contact point"
)

type GrafanaContactPointClient struct {
	*client.Client
}

type GrafanaContactPoint struct {
	Uid                   string                 `json:"uid"`
	Name                  string                 `json:"name"`
	Type                  string                 `json:"type"`
	Settings              map[string]interface{} `json:"settings"`
	DisableResolveMessage bool                   `json:"disableResolveMessage"`
	Provenance            string                 `json:"provenance"`
}

func New(apiToken string, baseUrl string) (*GrafanaContactPointClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	grafanaContactPointClient := &GrafanaContactPointClient{
		Client: client.New(apiToken, baseUrl),
	}

	return grafanaContactPointClient, nil
}
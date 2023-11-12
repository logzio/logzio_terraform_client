package grafana_contact_points

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaContactPointServiceEndpoint = "%s/v1/grafana/api/v1/provisioning/contact-points"

	operationCreateGrafanaContactPoint = "CreateGrafanaContactPoint"
	operationGetAllGrafanaContactPoint = "GetAllGrafanaContactPoint"
	operationUpdateGrafanaContactPoint = "UpdateGrafanaContactPoint"
	operationDeleteGrafanaContactPoint = "DeleteGrafanaContactPoint"

	grafanaContactPointResourceName = "grafana contact point"

	GrafanaContactPointTypeEmail          ContactPointType = "email"
	GrafanaContactPointTypeGoogleChat     ContactPointType = "googlechat"
	GrafanaContactPointTypeOpsgenie       ContactPointType = "opsgenie"
	GrafanaContactPointTypePagerduty      ContactPointType = "pagerduty"
	GrafanaContactPointTypeSlack          ContactPointType = "slack"
	GrafanaContactPointTypeMicrosoftTeams ContactPointType = "teams"
	GrafanaContactPointTypeVictorps       ContactPointType = "victorops"
	GrafanaContactPointTypeWebhook        ContactPointType = "webhook"
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

type ContactPointType string

func (cpt ContactPointType) String() string {
	return string(cpt)
}

func GetSupportedContactPointTypes() []ContactPointType {
	return []ContactPointType{
		GrafanaContactPointTypeEmail,
		GrafanaContactPointTypeGoogleChat,
		GrafanaContactPointTypeOpsgenie,
		GrafanaContactPointTypePagerduty,
		GrafanaContactPointTypeSlack,
		GrafanaContactPointTypeMicrosoftTeams,
		GrafanaContactPointTypeVictorps,
		GrafanaContactPointTypeWebhook,
	}
}

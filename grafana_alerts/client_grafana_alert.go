package grafana_alerts

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
	"time"
)

const (
	grafanaObjectServiceEndpoint = "%s/v1/grafana/api/v1/provisioning/alert-rules"

	grafanaAlertResourceName = "grafana alert"

	operationCreateGrafanaAlert = "CreateGrafanaAlert"
	operationGetGrafanaAlert    = "GetGrafanaAlert"
	operationUpdateGrafanaAlert = "UpdateGrafanaAlert"
	operationDeleteGrafanaAlert = "DeleteGrafanaAlert"
	operationListGrafanaAlerts  = "ListGrafanaAlerts"
)

type GrafanaAlertClient struct {
	*client.Client
}

type AlertRule struct {
	Annotations  map[string]string `json:"annotations,omitempty"`
	Condition    string            `json:"condition"`    // Required
	Data         []*AlertQuery     `json:"data"`         // Required
	ExecErrState string            `json:"execErrState"` // Required
	FolderUID    string            `json:"folderUID"`    // Required
	For          string            `json:"for"`          // Required
	Id           int64             `json:"id,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	NoDataState  string            `json:"noDataState"` // Required
	OrgID        int64             `json:"orgID"`       // Required
	Provenance   string            `json:"provenance,omitempty"`
	RuleGroup    string            `json:"ruleGroup"` // Required
	Title        string            `json:"title"`     // Required
	Uid          string            `json:"uid,omitempty"`
	Updated      time.Time         `json:"updated"`
}

type AlertQuery struct {
	DatasourceUid     string               `json:"datasourceUid"`
	Model             interface{}          `json:"model"`
	QueryType         string               `json:"queryType"`
	RefId             string               `json:"refId"`
	RelativeTimeRange RelativeTimeRangeObj `json:"relativeTimeRange"`
}

type RelativeTimeRangeObj struct {
	From time.Duration `json:"from"`
	To   time.Duration `json:"to"`
}

func New(apiToken string, baseUrl string) (*GrafanaAlertClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	grafanaAlertClient := &GrafanaAlertClient{
		Client: client.New(apiToken, baseUrl),
	}

	return grafanaAlertClient, nil
}

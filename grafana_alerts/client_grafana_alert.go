package grafana_alerts

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
	"strings"
	"time"
)

type ExecErrState string
type NoDataState string

const (
	grafanaAlertServiceEndpoint = "%s/v1/grafana/api/v1/provisioning/alert-rules"

	grafanaAlertResourceName = "grafana alert rule"

	operationCreateGrafanaAlert = "CreateGrafanaAlert"
	operationGetGrafanaAlert    = "GetGrafanaAlert"
	operationUpdateGrafanaAlert = "UpdateGrafanaAlert"
	operationDeleteGrafanaAlert = "DeleteGrafanaAlert"
	operationListGrafanaAlerts  = "ListGrafanaAlerts"

	ErrOK          ExecErrState = "OK"
	ErrError       ExecErrState = "Error"
	ErrAlerting    ExecErrState = "Alerting"
	NoDataOk       NoDataState  = "OK"
	NoData         NoDataState  = "NoData"
	NoDataAlerting NoDataState  = "Alerting"
)

type GrafanaAlertClient struct {
	*client.Client
}

type GrafanaAlertRule struct {
	Annotations  map[string]string    `json:"annotations,omitempty"`
	Condition    string               `json:"condition"`    // Required
	Data         []*GrafanaAlertQuery `json:"data"`         // Required
	ExecErrState ExecErrState         `json:"execErrState"`
	FolderUID    string               `json:"folderUID"`    // Required
	For          string               `json:"for"`          // Required
	Id           int64                `json:"id,omitempty"`
	Labels       map[string]string    `json:"labels,omitempty"`
	NoDataState  NoDataState          `json:"noDataState"` // Required
	OrgID        int64                `json:"orgID"`       // Required
	Provenance   string               `json:"provenance,omitempty"`
	RuleGroup    string               `json:"ruleGroup"` // Required
	Title        string               `json:"title"`     // Required
	Uid          string               `json:"uid,omitempty"`
	Updated      time.Time            `json:"updated"`
	IsPaused     bool                 `json:"isPaused"`
}

type GrafanaAlertQuery struct {
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

func validateGrafanaAlertRuleCreateUpdate(payload GrafanaAlertRule, isUpdate bool) error {
	if len(payload.Condition) == 0 {
		return fmt.Errorf("Field condition must be set!")
	}

	if payload.Data == nil || len(payload.Data) == 0 {
		return fmt.Errorf("Field data must be set!")
	}

	if len(payload.FolderUID) == 0 {
		return fmt.Errorf("Field folderUID must be set!")
	}

	if payload.For == "" {
		return fmt.Errorf("Field for must be set!")
	}

	if len(payload.NoDataState) == 0 {
		return fmt.Errorf("Field noDataState must be set!")
	}

	if len(payload.RuleGroup) == 0 {
		return fmt.Errorf("Field ruleGroup must be set!")
	}

	if len(payload.Title) == 0 {
		return fmt.Errorf("Field title must be set!")
	} else if strings.Contains(payload.Title, "\\") || strings.Contains(payload.Title, "/") {
		return fmt.Errorf("alert Title cannot contain '/' or '\\\\' charchters")
	}

	if isUpdate {
		if len(payload.Uid) == 0 {
			return fmt.Errorf("Field uid must be set when updating a Grafana alert rule!")
		}

	} else {
		if payload.OrgID == 0 {
			return fmt.Errorf("Field orgID must be set!")
		}
	}

	return nil
}

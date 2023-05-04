package grafana_alerts

import (
	"github.com/logzio/logzio_terraform_client/client"
	"time"
)

const (
	grafanaObjectServiceEndpoint = "%s/v1/grafana/api/v1/provisioning/alert-rules"
)

type GrafanaAlertsClient struct {
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
	Provenance   ProvenanceObj     `json:"provenance"`
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

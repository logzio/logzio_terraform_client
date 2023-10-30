package grafana_notification_policies

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	grafanaNotificationPolicyServiceEndpoint = "%s/v1/grafana/api/v1/provisioning/policies"

	grafanaNotificationPolicyResourceName   = "grafana notification policy"
	operationSetGrafanaNotificationPolicy   = "SetNotificationPolicyTree"
	operationGetGrafanaNotificationPolicy   = "GetNotificationPolicyTree"
	operationResetGrafanaNotificationPolicy = "ResetNotificationPolicyTree"

	MatchTypeEqual     MatchType = "="
	MatchTypeNotEqual  MatchType = "!="
	MatchTypeRegexp    MatchType = "=~"
	MatchTypeNotRegexp MatchType = "!~"
)

type GrafanaNotificationPolicyClient struct {
	*client.Client
}

type GrafanaNotificationPolicyTree struct {
	GroupBy        []string                    `json:"group_by,omitempty"`
	GroupInterval  string                      `json:"group_interval,omitempty"`
	GroupWait      string                      `json:"group_wait,omitempty"`
	Provenance     string                      `json:"provenance,omitempty"`
	Receiver       string                      `json:"receiver,omitempty"`
	RepeatInterval string                      `json:"repeat_interval,omitempty"`
	Routes         []GrafanaNotificationPolicy `json:"routes,omitempty"`
}

type GrafanaNotificationPolicy struct {
	Continue          bool                        `json:"continue"`
	GroupBy           []string                    `json:"group_by,omitempty"`
	GroupInterval     string                      `json:"group_interval,omitempty"`
	GroupWait         string                      `json:"group_wait,omitempty"`
	MuteTimeIntervals []string                    `json:"mute_time_intervals,omitempty"`
	ObjectMatchers    MatchersObj                 `json:"object_matchers,omitempty"`
	Receiver          string                      `json:"receiver,omitempty"`
	RepeatInterval    string                      `json:"repeat_interval,omitempty"`
	Routes            []GrafanaNotificationPolicy `json:"routes,omitempty"`
}

type MatchersObj []MatcherObj

type MatcherObj []string

type MatchType string

func New(apiToken string, baseUrl string) (*GrafanaNotificationPolicyClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	grafanaNotificationPolicyClient := &GrafanaNotificationPolicyClient{
		Client: client.New(apiToken, baseUrl),
	}

	return grafanaNotificationPolicyClient, nil
}

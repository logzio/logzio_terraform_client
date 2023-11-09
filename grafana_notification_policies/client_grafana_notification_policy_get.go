package grafana_notification_policies

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getGrafanaNotificationPolicyServiceUrl      = grafanaNotificationPolicyServiceEndpoint
	getGrafanaNotificationPolicyServiceMethod   = http.MethodGet
	getGrafanaNotificationPolicyServiceSuccess  = http.StatusOK
	getGrafanaNotificationPolicyServiceNotFound = http.StatusNotFound
)

func (c *GrafanaNotificationPolicyClient) GetGrafanaNotificationPolicyTree() (GrafanaNotificationPolicyTree, error) {
	var grafanaNotificationPolicy GrafanaNotificationPolicyTree
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getGrafanaNotificationPolicyServiceMethod,
		Url:          fmt.Sprintf(getGrafanaNotificationPolicyServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{getGrafanaNotificationPolicyServiceSuccess},
		NotFoundCode: getGrafanaNotificationPolicyServiceNotFound,
		ResourceId:   nil,
		ApiAction:    operationGetGrafanaNotificationPolicy,
		ResourceName: grafanaNotificationPolicyResourceName,
	})

	if err != nil {
		return grafanaNotificationPolicy, err
	}

	err = json.Unmarshal(res, &grafanaNotificationPolicy)
	if err != nil {
		return grafanaNotificationPolicy, err
	}

	return grafanaNotificationPolicy, nil
}

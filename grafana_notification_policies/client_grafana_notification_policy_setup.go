package grafana_notification_policies

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	setGrafanaNotificationPolicyServiceUrl     = grafanaNotificationPolicyServiceEndpoint
	setGrafanaNotificationPolicyServiceMethod  = http.MethodPut
	setGrafanaNotificationPolicyMethodAccepted = http.StatusAccepted
	setGrafanaNotificationPolicyStatusNotFound = http.StatusNotFound
)

func (c *GrafanaNotificationPolicyClient) SetGrafanaNotificationPolicyTree(payload GrafanaNotificationPolicy) error {
	setGrafanaNotificationPolicyJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   setGrafanaNotificationPolicyServiceMethod,
		Url:          fmt.Sprintf(setGrafanaNotificationPolicyServiceUrl, c.BaseUrl),
		Body:         setGrafanaNotificationPolicyJson,
		SuccessCodes: []int{setGrafanaNotificationPolicyMethodAccepted},
		NotFoundCode: setGrafanaNotificationPolicyStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationSetGrafanaNotificationPolicy,
		ResourceName: grafanaNotificationPolicyResourceName,
	})

	return err
}

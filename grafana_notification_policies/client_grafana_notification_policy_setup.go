package grafana_notification_policies

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	setupGrafanaNotificationPolicyServiceUrl     = grafanaNotificationPolicyServiceEndpoint
	setupGrafanaNotificationPolicyServiceMethod  = http.MethodPut
	setupGrafanaNotificationPolicyMethodAccepted = http.StatusAccepted
	setupGrafanaNotificationPolicyStatusNotFound = http.StatusNotFound
)

func (c *GrafanaNotificationPolicyClient) SetupGrafanaNotificationPolicyTree(payload GrafanaNotificationPolicyTree) error {
	setGrafanaNotificationPolicyJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   setupGrafanaNotificationPolicyServiceMethod,
		Url:          fmt.Sprintf(setupGrafanaNotificationPolicyServiceUrl, c.BaseUrl),
		Body:         setGrafanaNotificationPolicyJson,
		SuccessCodes: []int{setupGrafanaNotificationPolicyMethodAccepted},
		NotFoundCode: setupGrafanaNotificationPolicyStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationSetGrafanaNotificationPolicy,
		ResourceName: grafanaNotificationPolicyResourceName,
	})

	return err
}

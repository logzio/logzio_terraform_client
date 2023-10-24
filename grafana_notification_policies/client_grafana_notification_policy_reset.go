package grafana_notification_policies

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	resetGrafanaNotificationPolicyServiceUrl    = grafanaNotificationPolicyServiceEndpoint
	resetGrafanaNotificationPolicyServiceMethod = http.MethodDelete
	resetGrafanaNotificationPolicyServiceOk     = http.StatusOK
	resetGrafanaNotificationPolicyNotFound      = http.StatusNotFound
)

func (c *GrafanaNotificationPolicyClient) ResetGrafanaNotificationPolicyTree() error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   resetGrafanaNotificationPolicyServiceMethod,
		Url:          fmt.Sprintf(resetGrafanaNotificationPolicyServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{resetGrafanaNotificationPolicyServiceOk},
		NotFoundCode: resetGrafanaNotificationPolicyNotFound,
		ResourceId:   nil,
		ApiAction:    operationResetGrafanaNotificationPolicy,
		ResourceName: grafanaNotificationPolicyResourceName,
	})

	return err
}

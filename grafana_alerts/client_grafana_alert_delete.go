package grafana_alerts

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteGrafanaAlertServiceUrl     = grafanaAlertServiceEndpoint + "/%s"
	deleteGrafanaAlertServiceMethod  = http.MethodDelete
	deleteGrafanaAlertServiceSuccess = http.StatusNoContent
	deleteGrafanaAlertNotFound       = http.StatusNotFound
)

func (c *GrafanaAlertClient) DeleteGrafanaAlertRule(uid string) error {
	if uid == "" {
		return fmt.Errorf("uid is empty")
	}
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteGrafanaAlertServiceMethod,
		Url:          fmt.Sprintf(deleteGrafanaAlertServiceUrl, c.BaseUrl, uid),
		Body:         nil,
		SuccessCodes: []int{deleteGrafanaAlertServiceSuccess},
		NotFoundCode: deleteGrafanaAlertNotFound,
		ResourceId:   uid,
		ApiAction:    operationDeleteGrafanaAlert,
		ResourceName: grafanaAlertResourceName,
	})

	return err
}

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
	// NOTE: the grafana api returns 204 even when you try to delete with a uid that doesn't exist,
	// so the following line is just for compatibility with the CallLogzioApi object
	deleteGrafanaAlertNotFound = http.StatusNotFound
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

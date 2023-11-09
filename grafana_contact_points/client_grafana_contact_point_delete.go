package grafana_contact_points

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteGrafanaContactPointServiceUrl             = grafanaContactPointServiceEndpoint + "/%s"
	deleteGrafanaContactPointServiceMethod          = http.MethodDelete
	deleteGrafanaContactPointMethodSuccessNoContent = http.StatusNoContent
	deleteGrafanaContactPointMethodSuccessAccepted  = http.StatusAccepted
	// The status not found part is just for compatability with the CallLogzioApi function,
	// The actual API doesn't return an error if you try to delete a contact point that doesn't exist
	deleteGrafanaContactPointMethodNotFound = http.StatusNotFound
)

func (c *GrafanaContactPointClient) DeleteGrafanaContactPoint(uid string) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteGrafanaContactPointServiceMethod,
		Url:          fmt.Sprintf(deleteGrafanaContactPointServiceUrl, c.BaseUrl, uid),
		Body:         nil,
		SuccessCodes: []int{deleteGrafanaContactPointMethodSuccessNoContent, deleteGrafanaContactPointMethodSuccessAccepted},
		NotFoundCode: deleteGrafanaContactPointMethodNotFound,
		ResourceId:   uid,
		ApiAction:    operationDeleteGrafanaContactPoint,
		ResourceName: grafanaContactPointResourceName,
	})

	return err
}

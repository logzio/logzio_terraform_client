package grafana_contact_points

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateGrafanaContactPointServiceUrl     = grafanaContactPointServiceEndpoint + "/%s"
	updateGrafanaContactPointServiceMethod  = http.MethodPut
	updateGrafanaContactPointMethodAccepted = http.StatusAccepted
	updateGrafanaContactPointMethodNotFound = http.StatusNotFound
)

func (c *GrafanaContactPointClient) UpdateContactPoint(contactPoint GrafanaContactPoint) error {
	if len(contactPoint.Uid) == 0 {
		return fmt.Errorf("uid must be set")
	}

	updateGrafanaContactPointJson, err := json.Marshal(contactPoint)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateGrafanaContactPointServiceMethod,
		Url:          fmt.Sprintf(updateGrafanaContactPointServiceUrl, c.BaseUrl, contactPoint.Uid),
		Body:         updateGrafanaContactPointJson,
		SuccessCodes: []int{updateGrafanaContactPointMethodAccepted},
		NotFoundCode: updateGrafanaContactPointMethodNotFound,
		ResourceId:   contactPoint.Uid,
		ApiAction:    operationUpdateGrafanaContactPoint,
		ResourceName: grafanaContactPointResourceName,
	})

	return err
}

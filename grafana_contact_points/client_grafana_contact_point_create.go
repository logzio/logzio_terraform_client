package grafana_contact_points

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createGrafanaContactPointServiceUrl     = grafanaContactPointServiceEndpoint
	createGrafanaContactPointServiceMethod  = http.MethodPost
	createGrafanaContactPointMethodAccepted = http.StatusAccepted
	createGrafanaContactPointStatusNotFound = http.StatusNotFound
)

func (c *GrafanaContactPointClient) CreateGrafanaContactPoint(payload GrafanaContactPoint) error {
	grafanaContactPointJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createGrafanaContactPointServiceMethod,
		Url:          fmt.Sprintf(createGrafanaContactPointServiceUrl, c.BaseUrl),
		Body:         grafanaContactPointJson,
		SuccessCodes: []int{createGrafanaContactPointMethodAccepted},
		NotFoundCode: createGrafanaContactPointStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateGrafanaContactPoint,
		ResourceName: grafanaContactPointResourceName,
	})

	return err
}

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

func (c *GrafanaContactPointClient) CreateGrafanaContactPoint(payload GrafanaContactPoint) (GrafanaContactPoint, error) {
	err := validateContactPoint(payload)
	if err != nil {
		return GrafanaContactPoint{}, err
	}

	grafanaContactPointJson, err := json.Marshal(payload)
	if err != nil {
		return GrafanaContactPoint{}, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
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

	if err != nil {
		return GrafanaContactPoint{}, err
	}

	var grafanaContactPoint GrafanaContactPoint
	err = json.Unmarshal(res, &grafanaContactPoint)
	if err != nil {
		return GrafanaContactPoint{}, err
	}

	return grafanaContactPoint, nil
}

func validateContactPoint(payload GrafanaContactPoint) error {
	if len(payload.Name) == 0 {
		return fmt.Errorf("name must be set!")
	}

	if len(payload.Type) == 0 {
		return fmt.Errorf("type must be set!")
	}

	if payload.Settings == nil || len(payload.Settings) == 0 {
		return fmt.Errorf("settings must be set!")
	}

	return nil
}

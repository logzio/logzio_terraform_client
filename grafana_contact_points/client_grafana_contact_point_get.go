package grafana_contact_points

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getAllGrafanaContactPointServiceUrl      = grafanaContactPointServiceEndpoint
	getAllGrafanaContactPointServiceMethod   = http.MethodGet
	getAllGrafanaContactPointServiceSuccess  = http.StatusOK
	getAllGrafanaContactPointServiceNotFound = http.StatusNotFound
)

func (c *GrafanaContactPointClient) GetAllGrafanaContactPoints() ([]GrafanaContactPoint, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getAllGrafanaContactPointServiceMethod,
		Url:          fmt.Sprintf(getAllGrafanaContactPointServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{getAllGrafanaContactPointServiceSuccess},
		NotFoundCode: getAllGrafanaContactPointServiceNotFound,
		ResourceId:   nil,
		ApiAction:    operationGetAllGrafanaContactPoint,
		ResourceName: grafanaContactPointResourceName,
	})

	if err != nil {
		return nil, err
	}

	var grafanaContactPoints []GrafanaContactPoint
	err = json.Unmarshal(res, &grafanaContactPoints)
	if err != nil {
		return nil, err
	}

	return grafanaContactPoints, nil
}

// GetGrafanaContactPointByUid - The actual API doesn't have functionality of getting specific Contact Point by uid.
// This function wraps the GetAllGrafanaContactPoints function, and looks for a match of a given uid
func (c *GrafanaContactPointClient) GetGrafanaContactPointByUid(uid string) (GrafanaContactPoint, error) {
	if len(uid) == 0 {
		return GrafanaContactPoint{}, fmt.Errorf("uid must be set")
	}

	contactPoints, err := c.GetAllGrafanaContactPoints()
	if err != nil {
		return GrafanaContactPoint{}, err
	}

	for _, cp := range contactPoints {
		if cp.Uid == uid {
			return cp, nil
		}
	}

	return GrafanaContactPoint{}, fmt.Errorf("API call %s failed with missing %s %s, data: %s",
		operationGetByUidGrafanaContactPoint, grafanaContactPointResourceName, uid, "")
}

func (c *GrafanaContactPointClient) GetGrafanaContactPointsByName(name string) ([]GrafanaContactPoint, error) {
	var contactPoints []GrafanaContactPoint
	if len(name) == 0 {
		return nil, fmt.Errorf("uid must be set")
	}

	getContactPoints, err := c.GetAllGrafanaContactPoints()
	if err != nil {
		return nil, err
	}

	for _, cp := range getContactPoints {
		if cp.Name == name {
			contactPoints = append(contactPoints, cp)
		}
	}

	if len(contactPoints) == 0 {
		return nil, fmt.Errorf("could not find grafana contact points with name %s", name)
	}

	return contactPoints, nil
}

package drop_filters

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const deactivateDropFilterServiceUrl string = dropFiltersServiceEndpoint + "/%s/deactivate"
const deactivateDropFilterServiceMethod string = http.MethodPost
const deactivateDropFilterMethodSuccess int = http.StatusOK
const deactivateDropFilterMethodNotFound int = http.StatusNotFound

// DeactivateDropFilter deactivates drop filter by id, returns the deactivated drop filter or error if occurred
func (c *DropFiltersClient) DeactivateDropFilter(dropFilterId string) (*DropFilter, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deactivateDropFilterServiceMethod,
		Url:          fmt.Sprintf(deactivateDropFilterServiceUrl, c.BaseUrl, dropFilterId),
		Body:         nil,
		SuccessCodes: []int{deactivateDropFilterMethodSuccess},
		NotFoundCode: deactivateDropFilterMethodNotFound,
		ResourceId:   dropFilterId,
		ApiAction:    deactivateDropFilterOperation,
	})

	if err != nil {
		return nil, err
	}

	var dropFilter DropFilter
	err = json.Unmarshal(res, &dropFilter)
	if err != nil {
		return nil, err
	}

	return &dropFilter, nil
}

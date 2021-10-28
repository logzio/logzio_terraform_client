package drop_filters

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const activateDropFilterServiceUrl string = dropFiltersServiceEndpoint + "/%s/activate"
const activateDropFilterServiceMethod string = http.MethodPost
const activateDropFilterMethodSuccess int = http.StatusOK
const activateDropFilterMethodNotFound int = http.StatusNotFound

// ActivateDropFilter activates drop filter by id, returns the activated drop filter or error if occurred
func (c *DropFiltersClient) ActivateDropFilter(dropFilterId string) (*DropFilter, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   activateDropFilterServiceMethod,
		Url:          fmt.Sprintf(activateDropFilterServiceUrl, c.BaseUrl, dropFilterId),
		Body:         nil,
		SuccessCodes: []int{activateDropFilterMethodSuccess},
		NotFoundCode: activateDropFilterMethodNotFound,
		ResourceId:   dropFilterId,
		ApiAction:    activateDropFilterOperation,
		ResourceName: dropFilterResourceName,
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

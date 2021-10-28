package drop_filters

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	retrieveDropFilterServiceUrl     = dropFiltersServiceEndpoint + "/search"
	retrieveDropFilterServiceMethod  = http.MethodPost
	retrieveDropFilterMethodSuccess  = http.StatusOK
	retrieveDropFilterStatusNotFound = http.StatusNotFound
)

// RetrieveDropFilters returns all the drop filters in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *DropFiltersClient) RetrieveDropFilters() ([]DropFilter, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   retrieveDropFilterServiceMethod,
		Url:          fmt.Sprintf(retrieveDropFilterServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{retrieveDropFilterMethodSuccess},
		NotFoundCode: retrieveDropFilterStatusNotFound,
		ResourceId:   nil,
		ApiAction:    retrieveDropFiltersOperation,
		ResourceName: dropFilterResourceName,
	})

	if err != nil {
		return nil, err
	}

	var dropFilters []DropFilter
	err = json.Unmarshal(res, &dropFilters)

	if err != nil {
		return nil, err
	}

	return dropFilters, nil
}

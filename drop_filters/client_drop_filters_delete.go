package drop_filters

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteDropFilterServiceUrl     = dropFiltersServiceEndpoint + "/%s"
	deleteDropFilterServiceMethod  = http.MethodDelete
	deleteDropFilterMethodSuccess  = http.StatusOK
	deleteDropFilterMethodNotFound = http.StatusNotFound
)

// DeleteDropFilter deletes a drop filter, specified by its unique id, returns an error if a problem is encountered
func (c *DropFiltersClient) DeleteDropFilter(dropFilterId string) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteDropFilterServiceMethod,
		Url:          fmt.Sprintf(deleteDropFilterServiceUrl, c.BaseUrl, dropFilterId),
		Body:         nil,
		SuccessCodes: []int{deleteDropFilterMethodSuccess},
		NotFoundCode: deleteDropFilterMethodNotFound,
		ResourceId:   dropFilterId,
		ApiAction:    deleteDropFilterOperation,
	})

	return err
}

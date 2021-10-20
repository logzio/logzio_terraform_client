package endpoints

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteEndpointServiceUrl             = endpointServiceEndpoint + "/%d"
	deleteEndpointServiceMethod          = http.MethodDelete
	deleteEndpointMethodSuccess          = http.StatusOK
	deleteEndpointMethodSuccessNoContent = http.StatusNoContent
	deleteEndpointMethodNotFound         = http.StatusNotFound
)

// DeleteEndpoint deletes an endpoint, specified by its unique id, returns an error if a problem is encountered
func (c *EndpointsClient) DeleteEndpoint(endpointId int64) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteEndpointServiceMethod,
		Url:          fmt.Sprintf(deleteEndpointServiceUrl, c.BaseUrl, endpointId),
		Body:         nil,
		SuccessCodes: []int{deleteEndpointMethodSuccess, deleteEndpointMethodSuccessNoContent},
		NotFoundCode: deleteEndpointMethodNotFound,
		ResourceId:   endpointId,
		ApiAction:    deleteEndpointMethod,
	})

	return err
}

package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/jonboydell/logzio_client/test_utils"
)

func setupEndpointsTest() (*endpoints.EndpointsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err == nil {
		return endpoints.New(apiToken)
	}
	return nil, err
}

package endpoints

import "github.com/jonboydell/logzio_client/test_utils"

var endpoints *Endpoints
var apiToken string
var createdEndpoints []int64

func setupEndpointsTest() {
	apiToken = test_utils.GetApiToken()
	endpoints, _ = New(apiToken)
	createdEndpoints = []int64{}
}

func teardownEndpointsTest() {
	for x := 0; x < len(createdEndpoints); x++ {
		endpoints.DeleteEndpoint(createdEndpoints[x])
	}
}

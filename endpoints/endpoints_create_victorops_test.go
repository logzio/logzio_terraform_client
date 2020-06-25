package endpoints_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	endpointId        = int64(1234567)
	victorOpsEndpoint = "/v1/endpoints/victorops"
)

func TestEndpoints_CreateVictorOpsEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc(victorOpsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "title")
		assert.Contains(t, target, "description")
		assert.Contains(t, target, "routingKey")
		assert.Contains(t, target, "messageType")
		assert.Contains(t, target, "serviceApiKey")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:         "some_endpoint",
			Description:   "my description",
			RoutingKey:    "someRoutingKey",
			MessageType:   "someMessageType",
			ServiceApiKey: "someServiceApiKey",
			EndpointType:  endpoints.EndpointTypeVictorOps,
		})
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.Equal(t, endpointId, endpoint.Id)
	}
}

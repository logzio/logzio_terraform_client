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

func TestEndpoints_CreatePagerDutyEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)

	mux.HandleFunc("/v1/endpoints/pager-duty", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "title")
		assert.Contains(t, target, "description")
		assert.Contains(t, target, "serviceKey")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "some_endpoint",
			Description:  "my description",
			ServiceKey:   "SomeServiceKey",
			EndpointType: endpoints.EndpointTypePagerDuty,
		})
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.Equal(t, endpointId, endpoint.Id)
	}
}

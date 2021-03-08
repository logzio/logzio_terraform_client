package endpoints_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestEndpoints_CreateCustomEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)

	mux.HandleFunc("/v1/endpoints/custom", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.Contains(t, target, "title")
		assert.Contains(t, target, "description")
		assert.Contains(t, target, "url")
		assert.Contains(t, target, "headers")
		Headers := strings.Split(fmt.Sprint(target["headers"]), ",")
		assert.Equal(t, 2, len(Headers))
		assert.Equal(t, strings.Split(Headers[1], "=")[1], "two words")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "testCreateCustomEndpoint",
			Method:       "POST",
			Description:  "my description",
			Url:          "https://this.is.com/some/other/webhook",
			EndpointType: endpoints.EndpointTypeCustom,
			Headers:      map[string]string{"hello": "there", "header": "two words"},
			BodyTemplate: map[string]string{"hello": "there", "header": "two"},
		})
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.Equal(t, endpointId, endpoint.Id)
	}
}

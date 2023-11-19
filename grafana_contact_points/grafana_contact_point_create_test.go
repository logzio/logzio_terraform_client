package grafana_contact_points_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/grafana_contact_points"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestGrafanaContactPoint_CreateGrafanaContactPoint(t *testing.T) {
	underTest, teardown, err := setupGrafanaContactPointTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/contact-points", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_contact_points.GrafanaContactPoint
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Name)
			assert.NotEmpty(t, target.Settings)
			assert.NotEmpty(t, target.Type)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprint(w, fixture("create_grafana_contact_point.json"))
		})

		createContactPoint := getGrafanaContactPointObject()
		contactPoint, err := underTest.CreateGrafanaContactPoint(createContactPoint)
		if assert.NoError(t, err) && assert.NotEmpty(t, contactPoint) {
			assert.NotEmpty(t, contactPoint.Uid)
			assert.Equal(t, createContactPoint.Name, contactPoint.Name)
			assert.Equal(t, createContactPoint.Type, contactPoint.Type)
			assert.Equal(t, createContactPoint.DisableResolveMessage, contactPoint.DisableResolveMessage)
			assert.True(t, reflect.DeepEqual(createContactPoint.Settings, createContactPoint.Settings))
		}
	}
}

func TestGrafanaContactPoint_CreateGrafanaContactPointInternalServerError(t *testing.T) {
	underTest, teardown, err := setupGrafanaContactPointTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/contact-points", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_contact_points.GrafanaContactPoint
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Name)
			assert.NotEmpty(t, target.Settings)
			assert.NotEmpty(t, target.Type)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
		})

		createContactPoint := getGrafanaContactPointObject()
		contactPoint, err := underTest.CreateGrafanaContactPoint(createContactPoint)
		assert.Error(t, err)
		assert.Empty(t, contactPoint)
	}
}

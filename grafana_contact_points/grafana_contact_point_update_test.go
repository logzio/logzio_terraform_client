package grafana_contact_points_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/grafana_contact_points"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGrafanaContactPoint_UpdateGrafanaContactPoint(t *testing.T) {
	underTest, teardown, err := setupGrafanaContactPointTest()
	defer teardown()

	uid := "some-uid"

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/contact-points/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Contains(t, r.URL.String(), uid)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_contact_points.GrafanaContactPoint
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Name)
			assert.NotEmpty(t, target.Type)
			assert.NotEmpty(t, target.Settings)
			assert.Equal(t, uid, target.Uid)
			w.WriteHeader(http.StatusAccepted)
		})

		updateAlert := getGrafanaContactPointObject()
		updateAlert.Uid = uid
		err = underTest.UpdateContactPoint(updateAlert)
		assert.NoError(t, err)
	}

}

func TestGrafanaContactPoint_UpdateGrafanaContactPointInternalServerError(t *testing.T) {
	underTest, teardown, err := setupGrafanaContactPointTest()
	defer teardown()

	uid := "some-uid"

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/contact-points/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Contains(t, r.URL.String(), uid)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_contact_points.GrafanaContactPoint
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Name)
			assert.NotEmpty(t, target.Type)
			assert.NotEmpty(t, target.Settings)
			assert.Equal(t, uid, target.Uid)
			w.WriteHeader(http.StatusInternalServerError)
		})

		updateAlert := getGrafanaContactPointObject()
		updateAlert.Uid = uid
		err = underTest.UpdateContactPoint(updateAlert)
		assert.Error(t, err)
	}

}

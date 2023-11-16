package grafana_contact_points_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaContactPoint_DeleteGrafanaContactPoint(t *testing.T) {
	underTest, teardown, err := setupGrafanaContactPointTest()
	defer teardown()

	if assert.NoError(t, err) {
		grafanaAlertUid := "delete-me"

		mux.HandleFunc("/v1/grafana/api/v1/provisioning/contact-points/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), grafanaAlertUid)
			w.WriteHeader(http.StatusAccepted)
		})

		err = underTest.DeleteGrafanaContactPoint(grafanaAlertUid)
		assert.NoError(t, err)
	}
}

func TestGrafanaContactPoint_DeleteGrafanaContactPointInternalServerError(t *testing.T) {
	underTest, teardown, err := setupGrafanaContactPointTest()
	defer teardown()

	if assert.NoError(t, err) {
		grafanaAlertUid := "delete-me"

		mux.HandleFunc("/v1/grafana/api/v1/provisioning/contact-points/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), grafanaAlertUid)
			w.WriteHeader(http.StatusInternalServerError)
		})

		err = underTest.DeleteGrafanaContactPoint(grafanaAlertUid)
		assert.Error(t, err)
	}
}

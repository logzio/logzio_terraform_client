package grafana_contact_points_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaContactPoint_GetAllGrafanaContactPoints(t *testing.T) {
	underTest, teardown, err := setupGrafanaContactPointTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/contact-points", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("get_all_grafana_contact_points.json"))
		})

		contactPoints, err := underTest.GetAllGrafanaContactPoints()
		if assert.NoError(t, err) && assert.NotEmpty(t, contactPoints) {
			assert.Equal(t, 2, len(contactPoints))
		}
	}

}

func TestGrafanaContactPoint_GetAllGrafanaContactPointsInternalServerError(t *testing.T) {
	underTest, teardown, err := setupGrafanaContactPointTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/contact-points", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
		})

		contactPoints, err := underTest.GetAllGrafanaContactPoints()
		assert.Error(t, err)
		assert.Empty(t, contactPoints)
	}

}

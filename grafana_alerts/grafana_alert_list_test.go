package grafana_alerts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaAlert_ListGrafanaAlerts(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_grafana_alert_res.json"))
	})

	alerts, err := underTest.ListGrafanaAlertRules()
	assert.NoError(t, err)
	assert.NotNil(t, alerts)
	assert.Equal(t, 1, len(alerts))
}

func TestGrafanaAlert_ListGrafanaAlertsInternalServerError(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	alerts, err := underTest.ListGrafanaAlertRules()
	assert.Error(t, err)
	assert.Nil(t, alerts)
}

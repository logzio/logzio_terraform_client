package unified_alerts_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
)

func TestUnifiedAlerts_GetAlert(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertId := "get-alert-789"
	alertType := unified_alerts.UrlTypeLogs

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/v2/unified-alerts/%s/%s", alertType, alertId), func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("get_unified_alert.json"))
		})

		alert, err := underTest.GetUnifiedAlert(alertType, alertId)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		assert.Equal(t, alertId, alert.Id)
		assert.NotNil(t, alert.AlertConfiguration)
		assert.Equal(t, unified_alerts.TypeLogAlert, alert.AlertConfiguration.Type)
	}
}

func TestUnifiedAlerts_GetAlertNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertId := "non-existent-id"
	alertType := unified_alerts.UrlTypeLogs

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/v2/unified-alerts/%s/%s", alertType, alertId), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err := underTest.GetUnifiedAlert(alertType, alertId)
		assert.Error(t, err)
	}
}

func TestUnifiedAlerts_GetAlertAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertId := "test-alert-id"
	alertType := unified_alerts.UrlTypeLogs

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/v2/unified-alerts/%s/%s", alertType, alertId), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err := underTest.GetUnifiedAlert(alertType, alertId)
		assert.Error(t, err)
	}
}

func TestUnifiedAlerts_GetAlertEmptyId(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err := underTest.GetUnifiedAlert(unified_alerts.UrlTypeLogs, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "alertId must be set")
	}
}

func TestUnifiedAlerts_GetAlertInvalidType(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err := underTest.GetUnifiedAlert("invalid-type", "some-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "alertType must be one of")
	}
}

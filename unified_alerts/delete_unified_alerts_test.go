package unified_alerts_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
)

func TestUnifiedAlerts_DeleteAlert(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertId := "delete-alert-111"
	alertType := unified_alerts.UrlTypeMetrics

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/poc/unified-alerts/%s/%s", alertType, alertId), func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), alertId)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("delete_unified_alert.json"))
		})

		alert, err := underTest.DeleteUnifiedAlert(alertType, alertId)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		assert.Equal(t, alertId, alert.Id)
	}
}

func TestUnifiedAlerts_DeleteAlertNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertId := "non-existent-id"
	alertType := unified_alerts.UrlTypeLogs

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/poc/unified-alerts/%s/%s", alertType, alertId), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err := underTest.DeleteUnifiedAlert(alertType, alertId)
		assert.Error(t, err)
	}
}

func TestUnifiedAlerts_DeleteAlertAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertId := "test-alert-id"
	alertType := unified_alerts.UrlTypeLogs

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/poc/unified-alerts/%s/%s", alertType, alertId), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err := underTest.DeleteUnifiedAlert(alertType, alertId)
		assert.Error(t, err)
	}
}

func TestUnifiedAlerts_DeleteAlertEmptyId(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err := underTest.DeleteUnifiedAlert(unified_alerts.UrlTypeLogs, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "alertId must be set")
	}
}

func TestUnifiedAlerts_DeleteAlertInvalidType(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err := underTest.DeleteUnifiedAlert("invalid-type", "some-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "alertType must be one of")
	}
}

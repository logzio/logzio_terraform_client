package unified_alerts_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/unified_alerts"
	"github.com/stretchr/testify/assert"
)

func TestUnifiedAlerts_UpdateAlert(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertId := "update-alert-999"
	alertType := unified_alerts.UrlTypeLogs

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/poc/unified-alerts/%s/%s", alertType, alertId), func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Contains(t, r.URL.String(), alertId)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target unified_alerts.CreateUnifiedAlert
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("update_unified_alert.json"))
		})

		updateAlert := getCreateLogAlertType()
		updateAlert.Title = "Updated Alert"

		alert, err := underTest.UpdateUnifiedAlert(alertType, alertId, updateAlert)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		assert.Equal(t, alertId, alert.Id)
		assert.Equal(t, "Updated Alert", alert.Title)
	}
}

func TestUnifiedAlerts_UpdateAlertNotFound(t *testing.T) {
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

		updateAlert := getCreateLogAlertType()
		_, err := underTest.UpdateUnifiedAlert(alertType, alertId, updateAlert)
		assert.Error(t, err)
	}
}

func TestUnifiedAlerts_UpdateAlertAPIFail(t *testing.T) {
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

		updateAlert := getCreateLogAlertType()
		_, err := underTest.UpdateUnifiedAlert(alertType, alertId, updateAlert)
		assert.Error(t, err)
	}
}

func TestUnifiedAlerts_UpdateAlertEmptyId(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		updateAlert := getCreateLogAlertType()
		_, err := underTest.UpdateUnifiedAlert(unified_alerts.UrlTypeLogs, "", updateAlert)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "alertId must be set")
	}
}

func TestUnifiedAlerts_UpdateAlertInvalidType(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		updateAlert := getCreateLogAlertType()
		_, err := underTest.UpdateUnifiedAlert("invalid-type", "some-id", updateAlert)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "alertType must be one of")
	}
}

func TestUnifiedAlerts_UpdateAlertNoTitle(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		updateAlert := getCreateLogAlertType()
		updateAlert.Title = ""
		_, err := underTest.UpdateUnifiedAlert(unified_alerts.UrlTypeLogs, "some-id", updateAlert)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "title must be set")
	}
}

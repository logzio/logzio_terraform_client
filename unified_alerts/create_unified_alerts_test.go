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

func TestUnifiedAlerts_CreateLogAlert(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertType := unified_alerts.UrlTypeLogs

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/v2/unified-alerts/%s", alertType), func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target unified_alerts.CreateUnifiedAlert
			err = json.Unmarshal(jsonBytes, &target)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotNil(t, target.AlertConfiguration)
			assert.Equal(t, unified_alerts.TypeLogAlert, target.AlertConfiguration.Type)
			assert.NotEmpty(t, target.AlertConfiguration.SubComponents)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_unified_alert_log.json"))
		})

		testAlert := getCreateLogAlertType()

		alert, err := underTest.CreateUnifiedAlert(alertType, testAlert)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		assert.Equal(t, "test-log-alert-123", alert.Id)
		assert.NotNil(t, alert.AlertConfiguration)
		assert.Equal(t, unified_alerts.TypeLogAlert, alert.AlertConfiguration.Type)
		assert.True(t, alert.Enabled)
	}
}

func TestUnifiedAlerts_CreateMetricAlert(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertType := unified_alerts.UrlTypeMetrics

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/v2/unified-alerts/%s", alertType), func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target unified_alerts.CreateUnifiedAlert
			err = json.Unmarshal(jsonBytes, &target)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotNil(t, target.AlertConfiguration)
			assert.Equal(t, unified_alerts.TypeMetricAlert, target.AlertConfiguration.Type)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_unified_alert_metric.json"))
		})

		testAlert := getCreateMetricAlertType()

		alert, err := underTest.CreateUnifiedAlert(alertType, testAlert)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		assert.Equal(t, "test-metric-alert-456", alert.Id)
		assert.NotNil(t, alert.AlertConfiguration)
		assert.Equal(t, unified_alerts.TypeMetricAlert, alert.AlertConfiguration.Type)
	}
}

func TestUnifiedAlerts_CreateAlertAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	alertType := unified_alerts.UrlTypeLogs

	if assert.NoError(t, err) {
		mux.HandleFunc(fmt.Sprintf("/v2/unified-alerts/%s", alertType), func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("api_error.txt"))
		})
	}

	_, err = underTest.CreateUnifiedAlert(alertType, getCreateLogAlertType())
	assert.Error(t, err)
}

func TestUnifiedAlerts_CreateAlertNoTitle(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.Title = ""
	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title must be set")
}

func TestUnifiedAlerts_CreateAlertNoAlertConfiguration(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.AlertConfiguration = nil
	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "alertConfiguration must be set")
}

func TestUnifiedAlerts_CreateAlertNoType(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.AlertConfiguration.Type = ""
	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "alertConfiguration.type must be set")
}

func TestUnifiedAlerts_CreateAlertInvalidType(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.AlertConfiguration.Type = "INVALID_TYPE"
	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlertType)
	assert.Error(t, err)
}

func TestUnifiedAlerts_CreateAlertInvalidUrlType(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	_, err = underTest.CreateUnifiedAlert("invalid-url-type", createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "alertType must be one of")
}

func TestUnifiedAlerts_CreateLogAlertNoSubComponents(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.AlertConfiguration.SubComponents = nil
	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subComponents")
}

func TestUnifiedAlerts_CreateLogAlertEmptyQuery(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.AlertConfiguration.SubComponents[0].QueryDefinition.Query = ""
	_, err = underTest.CreateUnifiedAlert(unified_alerts.UrlTypeLogs, createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "query must be set")
}

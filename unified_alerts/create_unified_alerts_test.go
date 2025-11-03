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

	if assert.NoError(t, err) {
		mux.HandleFunc("/poc/unified-alerts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target unified_alerts.CreateUnifiedAlert
			err = json.Unmarshal(jsonBytes, &target)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.Equal(t, unified_alerts.TypeLogAlert, target.Type)
			assert.NotNil(t, target.LogAlert)
			assert.NotEmpty(t, target.LogAlert.SubComponents)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("create_unified_alert_log.json"))
		})

		testAlert := getCreateLogAlertType()

		alert, err := underTest.CreateUnifiedAlert(testAlert)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		assert.Equal(t, "test-log-alert-123", alert.Id)
		assert.Equal(t, unified_alerts.TypeLogAlert, alert.Type)
		assert.True(t, alert.Enabled)
	}
}

func TestUnifiedAlerts_CreateMetricAlert(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/poc/unified-alerts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			jsonBytes, _ := io.ReadAll(r.Body)
			var target unified_alerts.CreateUnifiedAlert
			err = json.Unmarshal(jsonBytes, &target)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.Equal(t, unified_alerts.TypeMetricAlert, target.Type)
			assert.NotNil(t, target.MetricAlert)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_unified_alert_metric.json"))
		})

		testAlert := getCreateMetricAlertType()

		alert, err := underTest.CreateUnifiedAlert(testAlert)
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		assert.Equal(t, "test-metric-alert-456", alert.Id)
		assert.Equal(t, unified_alerts.TypeMetricAlert, alert.Type)
	}
}

func TestUnifiedAlerts_CreateAlertAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/poc/unified-alerts", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("api_error.txt"))
		})
	}

	_, err = underTest.CreateUnifiedAlert(getCreateLogAlertType())
	assert.Error(t, err)
}

func TestUnifiedAlerts_CreateAlertNoTitle(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.Title = ""
	_, err = underTest.CreateUnifiedAlert(createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title must be set")
}

func TestUnifiedAlerts_CreateAlertNoType(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.Type = ""
	_, err = underTest.CreateUnifiedAlert(createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "type must be set")
}

func TestUnifiedAlerts_CreateAlertInvalidType(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.Type = "INVALID_TYPE"
	_, err = underTest.CreateUnifiedAlert(createAlertType)
	assert.Error(t, err)
}

func TestUnifiedAlerts_CreateLogAlertMissingLogAlert(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.LogAlert = nil
	_, err = underTest.CreateUnifiedAlert(createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "logAlert must be set")
}

func TestUnifiedAlerts_CreateMetricAlertMissingMetricAlert(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateMetricAlertType()
	createAlertType.MetricAlert = nil
	_, err = underTest.CreateUnifiedAlert(createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "metricAlert must be set")
}

func TestUnifiedAlerts_CreateLogAlertNoSubComponents(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.LogAlert.SubComponents = nil
	_, err = underTest.CreateUnifiedAlert(createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subComponents")
}

func TestUnifiedAlerts_CreateLogAlertEmptyQuery(t *testing.T) {
	underTest, err, teardown := setupUnifiedAlertsTest()
	defer teardown()

	createAlertType := getCreateLogAlertType()
	createAlertType.LogAlert.SubComponents[0].QueryDefinition.Query = ""
	_, err = underTest.CreateUnifiedAlert(createAlertType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "query must be set")
}

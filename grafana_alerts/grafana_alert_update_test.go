package grafana_alerts_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/grafana_alerts"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGrafanaAlert_UpdateGrafanaAlert(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	assert.NoError(t, err)
	defer teardown()

	alertUid := "some-uid"

	mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), alertUid)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target grafana_alerts.GrafanaAlertRule
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		assert.NotEmpty(t, target.Title)
		assert.Equal(t, alertUid, target.Uid)
	})

	updateAlert := getGrafanaAlertRuleObject()
	updateAlert.Uid = alertUid
	err = underTest.UpdateGrafanaAlertRule(updateAlert)
	assert.NoError(t, err)
}

func TestGrafanaAlert_UpdateGrafanaAlertInternalServerError(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	assert.NoError(t, err)
	defer teardown()

	alertUid := "client_test"

	mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), alertUid)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target grafana_alerts.GrafanaAlertRule
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		assert.NotEmpty(t, target.Title)
		assert.Equal(t, alertUid, target.Uid)
		w.WriteHeader(http.StatusInternalServerError)
	})

	updateAlert := getGrafanaAlertRuleObject()
	updateAlert.Uid = alertUid
	err = underTest.UpdateGrafanaAlertRule(updateAlert)
	assert.Error(t, err)
}

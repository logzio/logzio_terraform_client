package grafana_alerts_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/grafana_alerts"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGrafanaAlert_CreateGrafanaAlert(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_alerts.GrafanaAlertRule
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_grafana_alert_res.json"))
		})

		createGrafanaAlert := getGrafanaAlertRuleObject()
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.NoError(t, err)
		assert.NotNil(t, grafanaAlert)
		assert.Equal(t, int64(123456), grafanaAlert.Id)
		assert.Equal(t, "some_uid", grafanaAlert.Uid)
		assert.Equal(t, "folder_uid", grafanaAlert.FolderUID)
		assert.Equal(t, createGrafanaAlert.RuleGroup, grafanaAlert.RuleGroup)
	}
}

func TestGrafanaAlert_CreateGrafanaAlertInternalServerError(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_alerts.GrafanaAlertRule
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			w.WriteHeader(http.StatusInternalServerError)
		})

		createGrafanaAlert := getGrafanaAlertRuleObject()
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)
	}
}

func TestGrafanaAlert_CreateGrafanaAlertInvalidTitle(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_alerts.GrafanaAlertRule
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			w.WriteHeader(http.StatusInternalServerError)
		})

		// test '/' naming limitation
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.Title = "client/test/title"
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)

		// test '\' naming limitation
		createGrafanaAlert.Title = "client\\test\\title"
		grafanaAlert, err = underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)
	}
}

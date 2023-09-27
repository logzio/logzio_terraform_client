package grafana_alerts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaAlert_GetGrafanaAlert(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	assert.NoError(t, err)
	defer teardown()

	alertUid := "some-uid"

	mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), alertUid)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_grafana_alert_res.json"))
	})

	alert, err := underTest.GetGrafanaAlertRule(alertUid)
	assert.NoError(t, err)
	assert.NotNil(t, alert)
	assert.NoError(t, err)
	assert.NotNil(t, alert)
	assert.Equal(t, int64(123456), alert.Id)
	assert.Equal(t, "some_uid", alert.Uid)
	assert.Equal(t, "folder_uid", alert.FolderUID)
}

func TestGrafanaAlert_GetGrafanaAlertInternalError(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	assert.NoError(t, err)
	defer teardown()

	alertUid := "some-id"

	mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), alertUid)
		w.WriteHeader(http.StatusInternalServerError)
	})

	alert, err := underTest.GetGrafanaAlertRule(alertUid)
	assert.Error(t, err)
	assert.Nil(t, alert)
}

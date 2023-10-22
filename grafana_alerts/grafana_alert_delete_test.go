package grafana_alerts_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaAlert_DeleteGrafanaAlert(t *testing.T) {
	underTest, err, teardown := setupGrafanaAlertRuleTest()
	defer teardown()

	if assert.NoError(t, err) {
		grafanaAlertUid := "delete-me"

		mux.HandleFunc("/v1/grafana/api/v1/provisioning/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), grafanaAlertUid)
			w.WriteHeader(http.StatusNoContent)
		})

		err = underTest.DeleteGrafanaAlertRule(grafanaAlertUid)
		assert.NoError(t, err)
	}
}

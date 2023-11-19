package grafana_notification_policies_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaNotificationPolicy_ResetGrafanaNotificationPolicyTreeInternalError(t *testing.T) {
	underTest, err, teardown := setupGrafanaNotificationPolicyTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/policies", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
		})

		err := underTest.ResetGrafanaNotificationPolicyTree()
		assert.Error(t, err)
	}
}

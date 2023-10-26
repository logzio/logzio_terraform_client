package grafana_notification_policies_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/grafana_notification_policies"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGrafanaNotificationPolicy_SetupGrafanaNotificationPolicyTreeInternalError(t *testing.T) {
	underTest, err, teardown := setupGrafanaNotificationPolicyTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/policies", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_notification_policies.GrafanaNotificationPolicy
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Receiver)
			w.WriteHeader(http.StatusInternalServerError)
		})

		setupGrafanaNotificationPolicy := getGrafanaNotificationPolicyObject()
		err := underTest.SetupGrafanaNotificationPolicyTree(setupGrafanaNotificationPolicy)
		assert.Error(t, err)
	}
}

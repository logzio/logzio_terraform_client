package grafana_notification_policies_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaNotificationPolicy_GetGrafanaNotificationPolicyTreeInternalError(t *testing.T) {
	underTest, err, teardown := setupGrafanaNotificationPolicyTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/v1/provisioning/policies", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			fmt.Fprint(w, fixture("grafana_notification_policy.json"))
		})

		grafanaNotificationPolicy, err := underTest.GetGrafanaNotificationPolicyTree()
		assert.NoError(t, err)
		assert.NotNil(t, grafanaNotificationPolicy)
		assert.Equal(t, grafanaDefaultReceiver, grafanaNotificationPolicy.Receiver)
		assert.Equal(t, "10s", grafanaNotificationPolicy.GroupWait)
		assert.Equal(t, "5m", grafanaNotificationPolicy.GroupInterval)
		assert.Equal(t, "5h", grafanaNotificationPolicy.RepeatInterval)
		assert.Equal(t, 2, len(grafanaNotificationPolicy.GroupBy))
		assert.Equal(t, 3, len(grafanaNotificationPolicy.Routes))
		for _, route := range grafanaNotificationPolicy.Routes {
			assert.Equal(t, 3, len(route.ObjectMatchers[0]))
		}
	}
}

package grafana_notification_policies_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationGrafanaNotificationPolicy_GetGrafanaNotificationPolicyTree(t *testing.T) {
	underTest, err := setupGrafanaNotificationPolicyIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		grafanaNotificationPolicy, err := underTest.GetGrafanaNotificationPolicyTree()
		if assert.NoError(t, err) && assert.NotNil(t, grafanaNotificationPolicy) {
			assert.NotEmpty(t, grafanaNotificationPolicy.Receiver)
		}
	}
}

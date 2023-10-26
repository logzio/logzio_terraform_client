package grafana_notification_policies_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaNotificationPolicy_SetupGrafanaNotificationPolicyTree(t *testing.T) {
	underTest, err := setupGrafanaNotificationPolicyIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		setupGrafanaNotificationPolicy := getGrafanaNotificationPolicyObject()
		err = underTest.SetupGrafanaNotificationPolicyTree(setupGrafanaNotificationPolicy)
		if assert.NoError(t, err) {
			defer underTest.ResetGrafanaNotificationPolicyTree()
			time.Sleep(4 * time.Second)
		}
	}
}

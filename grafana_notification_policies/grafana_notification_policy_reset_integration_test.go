package grafana_notification_policies_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationGrafanaNotificationPolicy_ResetGrafanaNotificationPolicyTree(t *testing.T) {
	underTest, err := setupGrafanaNotificationPolicyIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		err = underTest.ResetGrafanaNotificationPolicyTree()
		assert.NoError(t, err)
	}
}

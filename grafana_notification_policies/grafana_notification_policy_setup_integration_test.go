package grafana_notification_policies_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_notification_policies"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"reflect"
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

func TestIntegrationGrafanaNotificationPolicy_SetupGrafanaNotificationPolicyTreeInvalidDefaultReceiver(t *testing.T) {
	underTest, err := setupGrafanaNotificationPolicyIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		setupGrafanaNotificationPolicy := getGrafanaNotificationPolicyObject()
		setupGrafanaNotificationPolicy.Receiver = "some-receiver"
		err = underTest.SetupGrafanaNotificationPolicyTree(setupGrafanaNotificationPolicy)
		assert.Error(t, err)
	}
}

func TestIntegrationGrafanaNotificationPolicy_SetupGrafanaNotificationPolicyTreeUpdate(t *testing.T) {
	underTest, err := setupGrafanaNotificationPolicyIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		setupGrafanaNotificationPolicy := getGrafanaNotificationPolicyObject()
		err = underTest.SetupGrafanaNotificationPolicyTree(setupGrafanaNotificationPolicy)
		if assert.NoError(t, err) {
			defer underTest.ResetGrafanaNotificationPolicyTree()
			time.Sleep(4 * time.Second)
			newRoute := grafana_notification_policies.GrafanaNotificationPolicy{
				Receiver:       grafanaDefaultReceiver,
				ObjectMatchers: grafana_notification_policies.MatchersObj{grafana_notification_policies.MatcherObj{"talktoyou", "=", "again"}},
				Continue:       true,
			}
			setupGrafanaNotificationPolicy.Routes = append(setupGrafanaNotificationPolicy.Routes, newRoute)
			err = underTest.SetupGrafanaNotificationPolicyTree(setupGrafanaNotificationPolicy)
			assert.NoError(t, err)
			time.Sleep(4 * time.Second)
			getNotificationPolicyTree, err := underTest.GetGrafanaNotificationPolicyTree()
			assert.NoError(t, err)
			assert.Equal(t, setupGrafanaNotificationPolicy.Receiver, getNotificationPolicyTree.Receiver)
			assert.Equal(t, setupGrafanaNotificationPolicy.GroupInterval, getNotificationPolicyTree.GroupInterval)
			assert.Equal(t, setupGrafanaNotificationPolicy.GroupWait, getNotificationPolicyTree.GroupWait)
			assert.Equal(t, setupGrafanaNotificationPolicy.RepeatInterval, getNotificationPolicyTree.RepeatInterval)
			assert.True(t, reflect.DeepEqual(setupGrafanaNotificationPolicy.Routes, getNotificationPolicyTree.Routes))
			assert.True(t, reflect.DeepEqual(setupGrafanaNotificationPolicy.GroupBy, getNotificationPolicyTree.GroupBy))
		}
	}
}

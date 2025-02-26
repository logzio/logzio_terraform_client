package grafana_alerts_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaAlert_CreateGrafanaAlert(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.Title = fmt.Sprintf("%s_%s", createGrafanaAlert.Title, "create")
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		if assert.NoError(t, err) && assert.NotNil(t, grafanaAlert) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteGrafanaAlertRule(grafanaAlert.Uid)
			assert.NotZero(t, grafanaAlert.Uid)
			assert.Equal(t, createGrafanaAlert.Title, grafanaAlert.Title)
		}
	}
}

func TestIntegrationGrafanaAlert_CreateGrafanaAlertNoTitle(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.Title = ""
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)
	}
}

func TestIntegrationGrafanaAlert_CreateGrafanaAlertNoData(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.Data = nil
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)
	}
}

func TestIntegrationGrafanaAlert_CreateGrafanaAlertNoFolderUid(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.FolderUID = ""
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)
	}
}

func TestIntegrationGrafanaAlert_CreateGrafanaAlertNoFor(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.For = ""
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)
	}
}

func TestIntegrationGrafanaAlert_CreateGrafanaAlertNoOrgId(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.OrgID = 0
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)
	}
}

func TestIntegrationGrafanaAlert_CreateGrafanaAlertNoOrgRuleGroup(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaAlert := getGrafanaAlertRuleObject()
		createGrafanaAlert.RuleGroup = ""
		grafanaAlert, err := underTest.CreateGrafanaAlertRule(createGrafanaAlert)
		assert.Error(t, err)
		assert.Nil(t, grafanaAlert)
	}
}

func TestIntegrationGrafanaAlert_CreateGrafanaAlertInvalidTitle(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
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

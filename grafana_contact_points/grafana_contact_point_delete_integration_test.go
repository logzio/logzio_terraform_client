package grafana_contact_points_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaContactPoint_DeleteGrafanaContactPoint(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createGrafanaContactPoint := getGrafanaContactPointObject()
		createGrafanaContactPoint.Name = fmt.Sprintf("%s_%s", createGrafanaContactPoint.Name, "delete")
		grafanaContactPoint, err := underTest.CreateGrafanaContactPoint(createGrafanaContactPoint)
		if assert.NoError(t, err) && assert.NotEmpty(t, grafanaContactPoint) && assert.NotEmpty(t, grafanaContactPoint.Uid) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteGrafanaContactPoint(grafanaContactPoint.Uid)
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationGrafanaAlert_DeleteGrafanaAlertEmptyUid(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		err = underTest.DeleteGrafanaContactPoint("")
		assert.Error(t, err)
	}
}

func TestIntegrationGrafanaAlert_DeleteGrafanaAlertUidNoExist(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		err = underTest.DeleteGrafanaContactPoint("some-uid")
		assert.Error(t, err)
	}
}

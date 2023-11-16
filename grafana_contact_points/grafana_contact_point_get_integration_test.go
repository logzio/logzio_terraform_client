package grafana_contact_points_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestIntegrationGrafanaContactPoint_GetAllGrafanaContactPoints(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		grafanaContactPoints, err := underTest.GetAllGrafanaContactPoints()
		assert.NoError(t, err)
		assert.NotNil(t, grafanaContactPoints)
	}
}

func TestIntegrationGrafanaContactPoint_GetGrafanaContactPoint(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaContactPoint := getGrafanaContactPointObject()
		createGrafanaContactPoint.Name = fmt.Sprintf("%s_%s", createGrafanaContactPoint.Name, "get")
		contactPoint, err := underTest.CreateGrafanaContactPoint(createGrafanaContactPoint)
		if assert.NoError(t, err) && assert.NotEmpty(t, contactPoint) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteGrafanaContactPoint(contactPoint.Uid)
			getContactPoint, err := underTest.GetGrafanaContactPointByUid(contactPoint.Uid)
			assert.NoError(t, err)
			assert.NotEmpty(t, getContactPoint)
			assert.True(t, reflect.DeepEqual(contactPoint, getContactPoint))
		}
	}
}

func TestIntegrationGrafanaContactPoint_GetGrafanaContactPointIdNotFound(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		getContactPoint, err := underTest.GetGrafanaContactPointByUid("some-uid")
		assert.Error(t, err)
		assert.Empty(t, getContactPoint)
	}
}

func TestIntegrationGrafanaContactPoint_GetGrafanaContactPointByName(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		contactPoints, err := underTest.GetGrafanaContactPointsByName("miri-slack2")
		assert.NoError(t, err)
		assert.NotEmpty(t, contactPoints)
	}
}

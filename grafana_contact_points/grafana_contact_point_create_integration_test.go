package grafana_contact_points_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaContactPoint_CreateGrafanaContactPoint(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaContactPoint := getGrafanaContactPointObject()
		createGrafanaContactPoint.Name = fmt.Sprintf("%s_%s", createGrafanaContactPoint.Name, "create")
		contactPoint, err := underTest.CreateGrafanaContactPoint(createGrafanaContactPoint)
		if assert.NoError(t, err) && assert.NotEmpty(t, contactPoint) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteGrafanaContactPoint(contactPoint.Uid)
			assert.NotEmpty(t, contactPoint.Uid)
			assert.Equal(t, createGrafanaContactPoint.Name, contactPoint.Name)
		}
	}
}

func TestIntegrationGrafanaContactPoint_CreateGrafanaContactPointNoName(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaContactPoint := getGrafanaContactPointObject()
		createGrafanaContactPoint.Name = ""
		contactPoint, err := underTest.CreateGrafanaContactPoint(createGrafanaContactPoint)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name must be set")
		assert.Empty(t, contactPoint)
	}
}

func TestIntegrationGrafanaContactPoint_CreateGrafanaContactPointNoType(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaContactPoint := getGrafanaContactPointObject()
		createGrafanaContactPoint.Type = ""
		contactPoint, err := underTest.CreateGrafanaContactPoint(createGrafanaContactPoint)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "type must be set")
		assert.Empty(t, contactPoint)
	}
}

func TestIntegrationGrafanaContactPoint_CreateGrafanaContactPointNoSettings(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaContactPoint := getGrafanaContactPointObject()
		createGrafanaContactPoint.Settings = nil
		contactPoint, err := underTest.CreateGrafanaContactPoint(createGrafanaContactPoint)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "settings must be set")
		assert.Empty(t, contactPoint)
	}
}

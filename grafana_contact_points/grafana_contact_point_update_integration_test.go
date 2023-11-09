package grafana_contact_points_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaContactPoint_UpdateGrafanaContactPoint(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()

	if assert.NoError(t, err) {
		createGrafanaContactPoint := getGrafanaContactPointObject()
		createGrafanaContactPoint.Name = fmt.Sprintf("%s_%s", createGrafanaContactPoint.Name, "update")
		grafanaContactPoint, err := underTest.CreateGrafanaContactPoint(createGrafanaContactPoint)
		if assert.NoError(t, err) && assert.NotEmpty(t, grafanaContactPoint) && assert.NotEmpty(t, grafanaContactPoint.Uid) {
			defer underTest.DeleteGrafanaContactPoint(grafanaContactPoint.Uid)
			time.Sleep(time.Second * 2)
			grafanaContactPoint.Name = "changed"
			err = underTest.UpdateContactPoint(grafanaContactPoint)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 4)
			getGrafanaContactPoint, err := underTest.GetGrafanaContactPointByUid(grafanaContactPoint.Uid)
			assert.NoError(t, err)
			assert.Equal(t, grafanaContactPoint.Name, getGrafanaContactPoint.Name)
		}
	}
}

func TestIntegrationGrafanaContactPoint_UpdateGrafanaContactPointUidNotFound(t *testing.T) {
	underTest, err := setupGrafanaContactPointIntegrationTest()

	if assert.NoError(t, err) {
		grafanaContactPoint := getGrafanaContactPointObject()
		grafanaContactPoint.Uid = "some-uid"
		err = underTest.UpdateContactPoint(grafanaContactPoint)
		assert.Error(t, err)
	}
}

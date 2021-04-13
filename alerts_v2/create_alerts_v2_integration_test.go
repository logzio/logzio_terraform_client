package alerts_v2_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAlertsV2_CreateAlert(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		alert, err := underTest.CreateAlert(createAlert)

		time.Sleep(3 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			t.Log(fmt.Sprintf("One alert id: %d", alert.AlertId))
			defer underTest.DeleteAlert(alert.AlertId)
		}
	}
}

func TestIntegrationAlertsV2_CreateAlertWithFilter(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		match := map[string]string{"type": "mytype"}
		must := map[string]interface{}{"match": match}
		createAlert.SubComponents[0].QueryDefinition.Filters.Bool.Must = append(createAlert.SubComponents[0].QueryDefinition.Filters.Bool.Must, must)
		alert, err := underTest.CreateAlert(createAlert)

		time.Sleep(3 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			defer underTest.DeleteAlert(alert.AlertId)
		}
	}
}

func TestIntegrationAlertsV2_CreateAlertInvalidFilter(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.SubComponents[0].QueryDefinition.Filters.Bool.Must = append(createAlert.SubComponents[0].QueryDefinition.Filters.Bool.Must, nil)

		alert, err := underTest.CreateAlert(createAlert)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}
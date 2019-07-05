package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	createAlert := createValidAlert()

	var alert *alerts.AlertType

	alert, err = underTest.CreateAlert(createAlert)
	assert.NoError(t, err)
	assert.NotNil(t, alert)

	alert, err = underTest.GetAlert(alert.AlertId)
	assert.NoError(t, err)
	assert.NotNil(t, alert)

	err = underTest.DeleteAlert(alert.AlertId)
	assert.NoError(t, err)

	_, err = underTest.GetAlert(12345)
	assert.Error(t, err)
}

package alerts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAlert(t *testing.T) {
	underTest, err := setupAlertsTest()

	createAlert := createValidAlert()
	alert, err := underTest.CreateAlert(createAlert)
	assert.NoError(t, err)
	if assert.NoError(t, err) && assert.NotNil(t, alert) {
		v, err := underTest.GetAlert(alert.AlertId)
		assert.NoError(t, err)
		assert.NotNil(t, v)
	}

	underTest.DeleteAlert(alert.AlertId)

	_, err = underTest.GetAlert(12345)
	assert.Error(t, err)
}

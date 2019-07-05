package alerts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateAlert(t *testing.T) {
	underTest, err := setupAlertsTest()
	if assert.NoError(t, err) {
		alert, err := underTest.CreateAlert(createValidAlert())
		assert.NoError(t, err)
		assert.NotNil(t, alert)
		underTest.DeleteAlert(alert.AlertId)
	}
}

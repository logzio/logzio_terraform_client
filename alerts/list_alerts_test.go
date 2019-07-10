package alerts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAlerts(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		_, err = underTest.ListAlerts()
		assert.NoError(t, err)
	}
}

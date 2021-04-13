package alerts_v2_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationAlertsV2_ListAlerts(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		_, err = underTest.ListAlerts()
		assert.NoError(t, err)
	}
}
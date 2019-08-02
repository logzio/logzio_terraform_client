// +build integration

package alerts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationAlerts_ListAlerts(t *testing.T) {
	underTest, err := setupAlertsIntegrationTest()

	if assert.NoError(t, err) {
		_, err = underTest.ListAlerts()
		assert.NoError(t, err)
	}
}

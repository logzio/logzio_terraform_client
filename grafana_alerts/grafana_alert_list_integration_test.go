package grafana_alerts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationGrafanaAlert_ListGrafanaAlerts(t *testing.T) {
	underTest, err := setupGrafanaAlertIntegrationTest()

	if assert.NoError(t, err) {
		alerts, err := underTest.ListGrafanaAlertRules()
		assert.NoError(t, err)
		assert.NotNil(t, alerts)
	}
}

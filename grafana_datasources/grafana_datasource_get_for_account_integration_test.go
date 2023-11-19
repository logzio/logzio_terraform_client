package grafana_datasources_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIntegrationGrafanaDatasource_GetForAccount(t *testing.T) {
	underTest, err := setupGrafanaDatasourceIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		if assert.NoError(t, err) {
			grafanaDatasource, err := underTest.GetForAccount(os.Getenv(envMetricsAccountName))
			assert.NoError(t, err)
			assert.NotNil(t, grafanaDatasource)
		}
	}
}

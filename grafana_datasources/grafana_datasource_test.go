package grafana_datasources_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_datasources"
	"github.com/logzio/logzio_terraform_client/test_utils"
)

const (
	envMetricsAccountName = "METRICS_ACCOUNT_NAME"
)

func setupGrafanaDatasourceIntegrationTest() (*grafana_datasources.GrafanaDatasourceClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := grafana_datasources.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

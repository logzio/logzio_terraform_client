package grafana_datasources

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getForAccountGrafanaDatasourceServiceUrl     = grafanaDatasourceServiceEndpoint + "/name/%s/summary"
	getForAccountGrafanaDatasourceServiceMethod  = http.MethodGet
	getForAccountGrafanaDatasourceMethodSuccess  = http.StatusOK
	getForAccountGrafanaDatasourceMethodNotFound = http.StatusNotFound
)

func (c *GrafanaDatasourceClient) GetForAccount(accountName string) (*GrafanaDataSource, error) {
	if len(accountName) == 0 {
		return nil, fmt.Errorf("account name must be set")
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getForAccountGrafanaDatasourceServiceMethod,
		Url:          fmt.Sprintf(getForAccountGrafanaDatasourceServiceUrl, c.BaseUrl, accountName),
		Body:         nil,
		SuccessCodes: []int{getForAccountGrafanaDatasourceMethodSuccess},
		NotFoundCode: getForAccountGrafanaDatasourceMethodNotFound,
		ResourceId:   nil,
		ApiAction:    GetAllForAccountGrafanaDatasourceMethod,
		ResourceName: grafanaDatasourceResourceName,
	})

	if err != nil {
		return nil, err
	}

	var grafanaDatasource GrafanaDataSource
	err = json.Unmarshal(res, &grafanaDatasource)
	if err != nil {
		return nil, err
	}

	return &grafanaDatasource, nil
}

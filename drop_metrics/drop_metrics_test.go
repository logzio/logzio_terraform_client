package drop_metrics_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
	"github.com/logzio/logzio_terraform_client/test_utils"
)

var (
	mux    *http.ServeMux
	client *drop_metrics.DropMetricsClient
	server *httptest.Server
)

func setupDropMetricsTest() (*drop_metrics.DropMetricsClient, *http.ServeMux, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = drop_metrics.New("test-token", server.URL)

	return client, mux, server.Close
}

func fixture(filename string) string {
	content, err := os.ReadFile(filepath.Join("testdata", "fixtures", filename))
	if err != nil {
		return fmt.Sprintf("fixture file not found: %s", filename)
	}
	return string(content)
}

func setupDropMetricsIntegrationTest() (*drop_metrics.DropMetricsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := drop_metrics.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func getCreateDropMetric() (drop_metrics.CreateUpdateDropMetric, error) {
	accountId, err := test_utils.GetMetricsAccountId()
	if err != nil {
		return drop_metrics.CreateUpdateDropMetric{}, err
	}
	active := true

	return drop_metrics.CreateUpdateDropMetric{
		AccountId: accountId,
		Active:    &active,
		Filter: drop_metrics.FilterObject{
			Operator: drop_metrics.OperatorAnd,
			Expression: []drop_metrics.FilterExpression{
				{
					Name:             "__name__",
					Value:            "test_metric",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
				{
					Name:             "service",
					Value:            "integration-test",
					ComparisonFilter: drop_metrics.ComparisonEq,
				},
			},
		},
	}, nil
}

func getBulkCreateDropMetrics() ([]drop_metrics.CreateUpdateDropMetric, error) {
	accountId, err := test_utils.GetMetricsAccountId()
	if err != nil {
		return nil, err
	}
	active := true

	return []drop_metrics.CreateUpdateDropMetric{
		{
			AccountId: accountId,
			Active:    &active,
			Filter: drop_metrics.FilterObject{
				Operator: drop_metrics.OperatorAnd,
				Expression: []drop_metrics.FilterExpression{
					{
						Name:             "__name__",
						Value:            "bulk_test_metric_1",
						ComparisonFilter: drop_metrics.ComparisonEq,
					},
				},
			},
		},
		{
			AccountId: accountId,
			Active:    &active,
			Filter: drop_metrics.FilterObject{
				Operator: drop_metrics.OperatorAnd,
				Expression: []drop_metrics.FilterExpression{
					{
						Name:             "__name__",
						Value:            "bulk_test_metric_2",
						ComparisonFilter: drop_metrics.ComparisonEq,
					},
				},
			},
		},
	}, nil
}

func getSearchDropMetricsRequest() (drop_metrics.SearchDropMetricsRequest, error) {
	accountId, err := test_utils.GetMetricsAccountId()
	if err != nil {
		return drop_metrics.SearchDropMetricsRequest{}, err
	}
	active := true

	return drop_metrics.SearchDropMetricsRequest{
		Filter: &drop_metrics.SearchFilter{
			AccountIds: []int64{accountId},
			Active:     &active,
		},
		Pagination: &drop_metrics.Pagination{
			PageNumber: 0,
			PageSize:   10,
		},
	}, nil
}

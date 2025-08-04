package drop_metrics_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_SearchDropMetrics(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression[0].Value = "test-metric-search"

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			searchReq := getSearchDropMetricsRequest()
			results, err := underTest.SearchDropMetrics(searchReq)

			if assert.NoError(t, err) && assert.NotNil(t, results) {
				assert.GreaterOrEqual(t, len(results), 1)

				// Find our created metric in the results
				found := false
				for _, result := range results {
					if result.Id == created.Id {
						found = true
						assert.Equal(t, created.AccountId, result.AccountId)
						assert.Equal(t, created.Enabled, result.Enabled)
						break
					}
				}
				assert.True(t, found, "Created drop metric should be found in search results")
			}
		}
	}
}

func TestIntegrationDropMetrics_SearchDropMetricsWithMetricName(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		testMetricName := "unique-test-metric-search-name"
		createReq.Filter.Expression[0].Value = testMetricName

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			searchReq := getSearchDropMetricsRequest()
			searchReq.Filter.MetricNames = []string{testMetricName}
			results, err := underTest.SearchDropMetrics(searchReq)

			if assert.NoError(t, err) && assert.NotNil(t, results) {
				assert.GreaterOrEqual(t, len(results), 1)

				// All results should have our metric name
				for _, result := range results {
					found := false
					for _, expr := range result.Filter.Expression {
						if expr.Name == "__name__" && expr.Value == testMetricName {
							found = true
							break
						}
					}
					assert.True(t, found, "Search result should contain the searched metric name")
				}
			}
		}
	}
}

func TestIntegrationDropMetrics_SearchDropMetricsWithPagination(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		searchReq := getSearchDropMetricsRequest()
		searchReq.Pagination.PageSize = 5
		searchReq.Pagination.PageNumber = 0

		results, err := underTest.SearchDropMetrics(searchReq)

		if assert.NoError(t, err) && assert.NotNil(t, results) {
			assert.LessOrEqual(t, len(results), 5, "Results should respect page size limit")
		}
	}
}

func TestIntegrationDropMetrics_SearchDropMetricsEmptyFilter(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		searchReq := drop_metrics.SearchDropMetricsRequest{
			Pagination: &drop_metrics.Pagination{
				PageNumber: 0,
				PageSize:   10,
			},
		}

		results, err := underTest.SearchDropMetrics(searchReq)

		assert.NoError(t, err)
		assert.NotNil(t, results)
	}
}

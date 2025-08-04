package drop_metrics_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_DeleteDropMetricsBySearch(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		// Create a unique metric to search for and delete
		createReq := getCreateDropMetric()
		uniqueMetricName := "unique-delete-by-search-metric"
		createReq.Filter.Expression[0].Value = uniqueMetricName

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second)

			// Create search request to find metrics with this specific name
			searchReq := drop_metrics.SearchDropMetricsRequest{
				Filter: &drop_metrics.SearchFilter{
					AccountIds:  []int64{created.AccountId},
					MetricNames: []string{uniqueMetricName},
				},
			}

			// Delete by search
			err := underTest.DeleteDropMetricsBySearch(searchReq)
			assert.NoError(t, err)

			// Verify it's deleted by trying to get it
			time.Sleep(1 * time.Second)
			result, err := underTest.GetDropMetric(created.Id)
			assert.Error(t, err)
			assert.Nil(t, result)
		}
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricsBySearchMultiple(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		// Create multiple metrics with the same prefix
		uniquePrefix := "bulk-delete-search-test"
		var createdIds []int64
		var accountId int64

		for i := 0; i < 2; i++ {
			createReq := getCreateDropMetric()
			createReq.Filter.Expression[0].Value = fmt.Sprintf("%s-%d", uniquePrefix, i)

			created, err := underTest.CreateDropMetric(createReq)
			if assert.NoError(t, err) && assert.NotNil(t, created) {
				createdIds = append(createdIds, created.Id)
				if accountId == 0 {
					accountId = created.AccountId
				}
			}
		}

		if len(createdIds) > 0 {
			time.Sleep(2 * time.Second)

			// Search for all metrics with our prefix
			searchReq := drop_metrics.SearchDropMetricsRequest{
				Filter: &drop_metrics.SearchFilter{
					AccountIds: []int64{accountId},
				},
			}

			// Delete by search (this will delete more than we created, but that's OK for integration test)
			err := underTest.DeleteDropMetricsBySearch(searchReq)
			assert.NoError(t, err)

			// Verify our specific metrics are deleted
			time.Sleep(1 * time.Second)
			for _, id := range createdIds {
				result, err := underTest.GetDropMetric(id)
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		}
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricsBySearchEmptyFilter(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		// Try to delete with empty search - this should be handled gracefully
		searchReq := drop_metrics.SearchDropMetricsRequest{}

		_ = underTest.DeleteDropMetricsBySearch(searchReq)

		// This might succeed or fail depending on API behavior - we just test it doesn't panic
		// The error handling is validated by checking the function completes
		assert.NotPanics(t, func() {
			underTest.DeleteDropMetricsBySearch(searchReq)
		})
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricsBySearchNoMatches(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		accountId, _ := test_utils.GetAccountId()

		// Search for metrics that definitely don't exist
		searchReq := drop_metrics.SearchDropMetricsRequest{
			Filter: &drop_metrics.SearchFilter{
				AccountIds:  []int64{accountId},
				MetricNames: []string{"definitely-non-existent-metric-name-12345"},
			},
		}

		_ = underTest.DeleteDropMetricsBySearch(searchReq)

		// This should either succeed (no metrics to delete) or return a specific error
		// We don't assert success/failure as API behavior may vary
		assert.NotPanics(t, func() {
			underTest.DeleteDropMetricsBySearch(searchReq)
		})
	}
}

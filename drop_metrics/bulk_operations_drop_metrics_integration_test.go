package drop_metrics_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_BulkCreateDropMetrics(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		bulkReq := getBulkCreateDropMetrics()
		// Make the metric names unique for this test
		bulkReq[0].Filter.Expression[0].Value = "bulk-test-metric-1-create"
		bulkReq[1].Filter.Expression[0].Value = "bulk-test-metric-2-create"

		results, err := underTest.BulkCreateDropMetrics(bulkReq)

		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, results) {
			// Clean up all created metrics
			for _, result := range results {
				defer underTest.DeleteDropMetric(result.Id)
			}

			assert.Len(t, results, 2)

			// Verify both metrics were created correctly
			for i, result := range results {
				assert.NotZero(t, result.Id)
				assert.Equal(t, bulkReq[i].AccountId, result.AccountId)
				// API defaults new metrics to active: true regardless of input
				assert.True(t, result.Active, "API should default new metrics to active: true")
				assert.Equal(t, bulkReq[i].Filter.Operator, result.Filter.Operator)
			}
		}
	}
}

func TestIntegrationDropMetrics_BulkDeleteDropMetrics(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		// First create some metrics to delete
		bulkCreateReq := getBulkCreateDropMetrics()
		bulkCreateReq[0].Filter.Expression[0].Value = "bulk-test-metric-1-delete"
		bulkCreateReq[1].Filter.Expression[0].Value = "bulk-test-metric-2-delete"

		created, err := underTest.BulkCreateDropMetrics(bulkCreateReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) && assert.Len(t, created, 2) {
			time.Sleep(2 * time.Second)

			// Extract IDs for bulk deletion
			idsToDelete := make([]int64, len(created))
			for i, metric := range created {
				idsToDelete[i] = metric.Id
			}

			// Bulk delete
			err := underTest.BulkDeleteDropMetrics(idsToDelete)
			assert.NoError(t, err)

			// Verify they're deleted by trying to get them
			time.Sleep(1 * time.Second)
			for _, id := range idsToDelete {
				result, err := underTest.GetDropMetric(id)
				assert.Error(t, err)
				assert.Nil(t, result)
			}
		}
	}
}

func TestIntegrationDropMetrics_BulkCreateDropMetricsInvalidRequest(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		bulkReq := getBulkCreateDropMetrics()
		// Make one request invalid
		bulkReq[0].AccountId = 0

		results, err := underTest.BulkCreateDropMetrics(bulkReq)

		assert.Error(t, err)
		assert.Nil(t, results)
	}
}

func TestIntegrationDropMetrics_BulkDeleteDropMetricsInvalidIds(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		idsToDelete := []int64{0, -1}

		err := underTest.BulkDeleteDropMetrics(idsToDelete)

		assert.Error(t, err)
	}
}

func TestIntegrationDropMetrics_BulkDeleteDropMetricsNotFound(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		idsToDelete := []int64{999999999, 999999998}

		err := underTest.BulkDeleteDropMetrics(idsToDelete)

		assert.Error(t, err)
	}
}

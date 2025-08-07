package drop_metrics_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_DeleteDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)

		createReq.Filter.Expression[0].Value = "test-metric-delete"

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second)

			err := underTest.DeleteDropMetric(created.Id)
			assert.NoError(t, err)

			// Verify it's deleted by trying to get it
			time.Sleep(1 * time.Second)
			result, err := underTest.GetDropMetric(created.Id)
			assert.Error(t, err)
			assert.Nil(t, result)
		}
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricMultiple(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		// Create multiple metrics and delete them one by one
		var createdIds []int64

		for i := 0; i < 3; i++ {
			createReq, err := getCreateDropMetric()
			assert.NoError(t, err)
			createReq.Filter.Expression[0].Value = fmt.Sprintf("test-metric-delete-multi-%d", i)

			created, err := underTest.CreateDropMetric(createReq)
			if assert.NoError(t, err) && assert.NotNil(t, created) {
				createdIds = append(createdIds, created.Id)
			}
		}

		time.Sleep(2 * time.Second)

		// Delete all created metrics
		for _, id := range createdIds {
			err := underTest.DeleteDropMetric(id)
			assert.NoError(t, err)

			// Verify deletion
			time.Sleep(1 * time.Second)
			result, err := underTest.GetDropMetric(id)
			assert.Error(t, err)
			assert.Nil(t, result)
		}
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricAfterDisable(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)
		createReq.Filter.Expression[0].Value = "test-metric-delete-after-disable"

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second)

			// First disable the metric
			_ = underTest.DisableDropMetric(created.Id)
			time.Sleep(1 * time.Second)

			// Then delete it
			err := underTest.DeleteDropMetric(created.Id)
			assert.NoError(t, err)

			// Verify it's deleted
			time.Sleep(1 * time.Second)
			result, err := underTest.GetDropMetric(created.Id)
			assert.Error(t, err)
			assert.Nil(t, result)
		}
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricNotFound(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.DeleteDropMetric(999999999)

		assert.Error(t, err)
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricInvalidId(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.DeleteDropMetric(0)

		assert.Error(t, err)
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricNegativeId(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.DeleteDropMetric(-1)

		assert.Error(t, err)
	}
}

func TestIntegrationDropMetrics_DeleteDropMetricTwice(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)
		createReq.Filter.Expression[0].Value = "test-metric-delete-twice"

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second)

			// First deletion should succeed
			err := underTest.DeleteDropMetric(created.Id)
			assert.NoError(t, err)

			time.Sleep(1 * time.Second)

			// Second deletion should fail with not found
			err = underTest.DeleteDropMetric(created.Id)
			assert.Error(t, err)
		}
	}
}

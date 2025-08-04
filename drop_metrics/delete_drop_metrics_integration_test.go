package drop_metrics_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_DeleteDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
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

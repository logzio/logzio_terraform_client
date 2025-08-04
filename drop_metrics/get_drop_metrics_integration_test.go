package drop_metrics_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_GetDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression[0].Value = "test-metric-get"

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			result, err := underTest.GetDropMetric(created.Id)
			if assert.NoError(t, err) && assert.NotNil(t, result) {
				assert.Equal(t, created.Id, result.Id)
				assert.Equal(t, created.AccountId, result.AccountId)
				// API uses "active" field - verify consistency
				assert.Equal(t, created.Active, result.Active, "Active field should match between create and get")
				assert.Equal(t, created.Filter.Operator, result.Filter.Operator)
				assert.Len(t, result.Filter.Expression, len(created.Filter.Expression))
			}
		}
	}
}

func TestIntegrationDropMetrics_GetDropMetricNotFound(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		result, err := underTest.GetDropMetric(999999999)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestIntegrationDropMetrics_GetDropMetricInvalidId(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		result, err := underTest.GetDropMetric(0)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

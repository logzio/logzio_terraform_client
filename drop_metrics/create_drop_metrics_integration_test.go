package drop_metrics_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_CreateDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression[0].Value = "test-metric-create"

		result, err := underTest.CreateDropMetric(createReq)

		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, result) {
			defer underTest.DeleteDropMetric(result.Id)
			assert.NotZero(t, result.Id)
			assert.Equal(t, createReq.AccountId, result.AccountId)
			assert.True(t, result.Enabled)
			assert.Equal(t, drop_metrics.OperatorAnd, result.Filter.Operator)
			assert.Len(t, result.Filter.Expression, 2)
		}
	}
}

func TestIntegrationDropMetrics_CreateDropMetricInvalidAccountId(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.AccountId = 0

		result, err := underTest.CreateDropMetric(createReq)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestIntegrationDropMetrics_CreateDropMetricNoFilter(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression = []drop_metrics.FilterExpression{}

		result, err := underTest.CreateDropMetric(createReq)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestIntegrationDropMetrics_CreateDropMetricInvalidOperator(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Operator = "INVALID"

		result, err := underTest.CreateDropMetric(createReq)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestIntegrationDropMetrics_CreateDropMetricInvalidComparison(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression[0].ComparisonFilter = "INVALID"

		result, err := underTest.CreateDropMetric(createReq)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestIntegrationDropMetrics_CreateDropMetricEmptyName(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression[0].Name = ""

		result, err := underTest.CreateDropMetric(createReq)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestIntegrationDropMetrics_CreateDropMetricEmptyValue(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression[0].Value = ""

		result, err := underTest.CreateDropMetric(createReq)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

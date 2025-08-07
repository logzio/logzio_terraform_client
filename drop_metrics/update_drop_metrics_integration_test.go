package drop_metrics_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_UpdateDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)
		createReq.Filter.Expression[0].Value = "test-metric-update"

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			// Update the metric
			active := created.Active
			updateReq := drop_metrics.UpdateDropMetric{
				AccountId: created.AccountId,
				Active:    &active,
				Filter: drop_metrics.FilterObject{
					Operator: drop_metrics.OperatorAnd,
					Expression: []drop_metrics.FilterExpression{
						{
							Name:             "__name__",
							Value:            "updated-test-metric",
							ComparisonFilter: drop_metrics.ComparisonEq,
						},
						{
							Name:             "service",
							Value:            "updated-integration-test",
							ComparisonFilter: drop_metrics.ComparisonEq,
						},
					},
				},
			}

			result, err := underTest.UpdateDropMetric(created.Id, updateReq)
			if assert.NoError(t, err) && assert.NotNil(t, result) {
				assert.Equal(t, created.Id, result.Id)
				assert.Equal(t, created.AccountId, result.AccountId)
				assert.Equal(t, updateReq.Filter.Operator, result.Filter.Operator)
				assert.Len(t, result.Filter.Expression, 2)

				// Verify the updated values
				found := false
				for _, expr := range result.Filter.Expression {
					if expr.Name == "__name__" && expr.Value == "updated-test-metric" {
						found = true
						break
					}
				}
				assert.True(t, found, "Updated metric name should be found")
			}
		}
	}
}

func TestIntegrationDropMetrics_UpdateDropMetricNotFound(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)
		updateReq := drop_metrics.UpdateDropMetric{
			AccountId: createReq.AccountId,
			Active:    createReq.Active,
			Filter:    createReq.Filter,
		}

		result, err := underTest.UpdateDropMetric(999999999, updateReq)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestIntegrationDropMetrics_UpdateDropMetricInvalidId(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)
		updateReq := drop_metrics.UpdateDropMetric{
			AccountId: createReq.AccountId,
			Active:    createReq.Active,
			Filter:    createReq.Filter,
		}

		result, err := underTest.UpdateDropMetric(0, updateReq)

		assert.Error(t, err)
		assert.Nil(t, result)
	}
}

func TestIntegrationDropMetrics_UpdateDropMetricInvalidAccountId(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)
		createReq.Filter.Expression[0].Value = "test-metric-update-invalid-account"

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			active := true
			updateReq := drop_metrics.UpdateDropMetric{
				AccountId: 0,
				Active:    &active,
				Filter:    createReq.Filter,
			}

			result, err := underTest.UpdateDropMetric(created.Id, updateReq)

			assert.Error(t, err)
			assert.Nil(t, result)
		}
	}
}

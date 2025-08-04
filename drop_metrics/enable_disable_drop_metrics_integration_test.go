package drop_metrics_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_EnableDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression[0].Value = "test-metric-enable"
		// Create as disabled
		disabled := false
		createReq.Enabled = &disabled

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			// Verify it's disabled initially
			assert.False(t, created.Enabled)

			// Enable it
			err := underTest.EnableDropMetric(created.Id)
			if assert.NoError(t, err) {
				// Verify by getting it again
				time.Sleep(1 * time.Second)
				retrieved, err := underTest.GetDropMetric(created.Id)
				if assert.NoError(t, err) && assert.NotNil(t, retrieved) {
					assert.True(t, retrieved.Enabled)
				}
			}
		}
	}
}

func TestIntegrationDropMetrics_DisableDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq := getCreateDropMetric()
		createReq.Filter.Expression[0].Value = "test-metric-disable"
		// Create as enabled
		enabled := true
		createReq.Enabled = &enabled

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			// Verify it's enabled initially
			assert.True(t, created.Enabled)

			// Disable it
			err := underTest.DisableDropMetric(created.Id)
			if assert.NoError(t, err) {
				// Verify by getting it again
				time.Sleep(1 * time.Second)
				retrieved, err := underTest.GetDropMetric(created.Id)
				if assert.NoError(t, err) && assert.NotNil(t, retrieved) {
					assert.False(t, retrieved.Enabled)
				}
			}
		}
	}
}

func TestIntegrationDropMetrics_EnableDropMetricNotFound(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.EnableDropMetric(999999999)

		assert.Error(t, err)
	}
}

func TestIntegrationDropMetrics_DisableDropMetricNotFound(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.DisableDropMetric(999999999)

		assert.Error(t, err)
	}
}

func TestIntegrationDropMetrics_EnableDropMetricInvalidId(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.EnableDropMetric(0)

		assert.Error(t, err)
	}
}

func TestIntegrationDropMetrics_DisableDropMetricInvalidId(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.DisableDropMetric(0)

		assert.Error(t, err)
	}
}

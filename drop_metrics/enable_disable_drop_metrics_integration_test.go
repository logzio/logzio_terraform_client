package drop_metrics_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropMetrics_EnableDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)
		createReq.Filter.Expression[0].Value = "test-metric-enable"
		// Create as disabled
		disabled := false
		createReq.Active = &disabled

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			t.Logf("Created metric - Active: %v, IsActive(): %v",
				created.Active, created.IsActive())

			// Note: API defaults to active: true regardless of input, so disable first if needed
			if created.Active {
				_ = underTest.DisableDropMetric(created.Id)
				time.Sleep(2 * time.Second)
			}

			// Enable it
			err := underTest.EnableDropMetric(created.Id)
			if assert.NoError(t, err) {
				// Verify by getting it again - wait longer for eventual consistency
				time.Sleep(5 * time.Second)
				retrieved, err := underTest.GetDropMetric(created.Id)
				if assert.NoError(t, err) && assert.NotNil(t, retrieved) {
					t.Logf("After enable - Active: %v, IsActive(): %v",
						retrieved.Active, retrieved.IsActive())

					// API uses "active" field - this should be true after enable
					assert.True(t, retrieved.Active, "Active field should be true after enable")
					assert.True(t, retrieved.IsActive(), "IsActive() should return true after enable")
				}
			}
		}
	}
}

func TestIntegrationDropMetrics_DisableDropMetric(t *testing.T) {
	underTest, err := setupDropMetricsIntegrationTest()
	if assert.NoError(t, err) {
		createReq, err := getCreateDropMetric()
		assert.NoError(t, err)

		createReq.Filter.Expression[0].Value = "test-metric-disable"
		// Create as enabled
		enabled := true
		createReq.Active = &enabled

		created, err := underTest.CreateDropMetric(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteDropMetric(created.Id)
			time.Sleep(2 * time.Second)

			t.Logf("Created metric - Active: %v, IsActive(): %v",
				created.Active, created.IsActive())

			// Note: API defaults to active: true, so metric should already be enabled
			// If it's disabled for some reason, enable it first
			if !created.Active {
				_ = underTest.EnableDropMetric(created.Id)
				time.Sleep(2 * time.Second)
			}

			// Disable it
			err := underTest.DisableDropMetric(created.Id)
			if assert.NoError(t, err) {
				// Verify by getting it again - wait longer for eventual consistency
				time.Sleep(5 * time.Second)
				retrieved, err := underTest.GetDropMetric(created.Id)
				if assert.NoError(t, err) && assert.NotNil(t, retrieved) {
					t.Logf("After disable - Active: %v, IsActive(): %v",
						retrieved.Active, retrieved.IsActive())

					// API uses "active" field - this should be false after disable
					assert.False(t, retrieved.Active, "Active field should be false after disable")
					assert.False(t, retrieved.IsActive(), "IsActive() should return false after disable")
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

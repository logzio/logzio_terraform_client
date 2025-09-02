package metrics_rollup_rules_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/metrics_rollup_rules"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationMetricsRollupRules_UpdateRollupRule(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.RollupFunction = metrics_rollup_rules.AggMax

		created, err := underTest.CreateRollupRule(req)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteRollupRule(created.Id)
			time.Sleep(2 * time.Second)

			req.MetricType = metrics_rollup_rules.MetricTypeCounter
			req.RollupFunction = ""

			updated, err := underTest.UpdateRollupRule(created.Id, req)
			if assert.NoError(t, err) && assert.NotNil(t, updated) {
				assert.Equal(t, req.MetricType, updated.MetricType)
			}
		}
	}
}

func TestIntegrationMetricsRollupRules_UpdateRollupRule_InvalidId(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		updateReq := metrics_rollup_rules.CreateUpdateRollupRule{MetricName: "x"}
		updated, err := underTest.UpdateRollupRule("", updateReq)
		assert.Error(t, err)
		assert.Nil(t, updated)
	}
}

func TestIntegrationMetricsRollupRules_UpdateRollupRule_NotFound(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		updateReq := metrics_rollup_rules.CreateUpdateRollupRule{MetricName: "x"}
		updated, err := underTest.UpdateRollupRule("non-existing-id", updateReq)
		assert.Error(t, err)
		assert.Nil(t, updated)
	}
}

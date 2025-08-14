package metrics_rollup_rules_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/metrics_rollup_rules"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationMetricsRollupRules_BulkCreateRollupRules(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		reqs, err := getBulkCreateRollupRules()
		assert.NoError(t, err)

		created, err := underTest.BulkCreateRollupRules(reqs)
		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, created) && assert.GreaterOrEqual(t, len(created), 1) {
			// cleanup
			ids := make([]string, 0, len(created))
			for _, r := range created {
				ids = append(ids, r.Id)
			}
			_ = underTest.BulkDeleteRollupRules(ids)
		}
	}
}

func TestIntegrationMetricsRollupRules_BulkCreateRollupRules_EmptyArray(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		results, err := underTest.BulkCreateRollupRules([]metrics_rollup_rules.CreateUpdateRollupRule{})
		assert.Error(t, err)
		assert.Nil(t, results)
	}
}

func TestIntegrationMetricsRollupRules_BulkDeleteRollupRules_InvalidIds(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.BulkDeleteRollupRules([]string{})
		assert.Error(t, err)
	}
}

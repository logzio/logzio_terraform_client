package metrics_rollup_rules_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationMetricsRollupRules_DeleteRollupRule(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)
		created, err := underTest.CreateRollupRule(req)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second)
			err := underTest.DeleteRollupRule(created.Id)
			assert.NoError(t, err)
		}
	}
}

func TestIntegrationMetricsRollupRules_DeleteRollupRule_NotFound(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		err := underTest.DeleteRollupRule("non-existent-id")
		assert.Error(t, err)
	}
}

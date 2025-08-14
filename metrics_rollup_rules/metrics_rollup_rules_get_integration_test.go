package metrics_rollup_rules_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationMetricsRollupRules_GetRollupRule(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)
		created, err := underTest.CreateRollupRule(req)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteRollupRule(created.Id)
			time.Sleep(2 * time.Second)

			got, err := underTest.GetRollupRule(created.Id)
			if assert.NoError(t, err) && assert.NotNil(t, got) {
				assert.Equal(t, created.Id, got.Id)
				assert.Equal(t, created.AccountId, got.AccountId)
			}
		}
	}
}

func TestIntegrationMetricsRollupRules_GetRollupRule_NotFound(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		got, err := underTest.GetRollupRule("non-existent-id")
		assert.Error(t, err)
		assert.Nil(t, got)
	}
}

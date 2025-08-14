package metrics_rollup_rules_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationMetricsRollupRules_CreateRollupRule(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		created, err := underTest.CreateRollupRule(req)
		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteRollupRule(created.Id)
			assert.NotEmpty(t, created.Id)
			assert.Equal(t, req.AccountId, created.AccountId)
			assert.Equal(t, req.MetricType, created.MetricType)
			assert.Equal(t, req.RollupFunction, created.RollupFunction)
		}
	}
}

func TestIntegrationMetricsRollupRules_CreateRollupRuleMissingMetricType(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.MetricType = ""

		created, err := underTest.CreateRollupRule(req)

		assert.Error(t, err)
		assert.Nil(t, created)
	}
}

func TestIntegrationMetricsRollupRules_CreateRollupRuleMissingRollupFunction(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.RollupFunction = ""

		created, err := underTest.CreateRollupRule(req)

		assert.Error(t, err)
		assert.Nil(t, created)
	}
}

func TestIntegrationMetricsRollupRules_CreateRollupRuleMissingLabelsEliminationMethod(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.LabelsEliminationMethod = ""

		created, err := underTest.CreateRollupRule(req)

		assert.Error(t, err)
		assert.Nil(t, created)
	}
}

func TestIntegrationMetricsRollupRules_CreateRollupRuleMissingLabels(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.Labels = nil

		created, err := underTest.CreateRollupRule(req)

		assert.Error(t, err)
		assert.Nil(t, created)
	}
}

func TestIntegrationMetricsRollupRules_CreateRollupRuleMissingMetricNameAndFilter(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.MetricName = ""
		req.Filter = nil

		created, err := underTest.CreateRollupRule(req)

		assert.Error(t, err)
		assert.Nil(t, created)
	}
}

func TestIntegrationMetricsRollupRules_CreateRollupRuleInvalidMetricType(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.MetricType = "INVALID"

		created, err := underTest.CreateRollupRule(req)

		assert.Error(t, err)
		assert.Nil(t, created)
	}
}

func TestIntegrationMetricsRollupRules_CreateRollupRuleInvalidRollupFunction(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.RollupFunction = "INVALID"

		created, err := underTest.CreateRollupRule(req)

		assert.Error(t, err)
		assert.Nil(t, created)
	}
}

func TestIntegrationMetricsRollupRules_CreateRollupRuleInvalidLabelsEliminationMethod(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)

		req.LabelsEliminationMethod = "INVALID"

		created, err := underTest.CreateRollupRule(req)

		assert.Error(t, err)
		assert.Nil(t, created)
	}
}

package security_rules_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestIntegrationSecurityRules_CreateRule(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test create"
		rule, err := underTest.CreateSecurityRule(createRule)

		if assert.NoError(t, err) && assert.NotNil(t, rule) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSecurityRule(rule.Id)
			assert.NotZero(t, rule.Id)
			assert.Equal(t, createRule.Title, rule.Title)
			assert.Equal(t, createRule.Description, rule.Description)
			assert.Equal(t, createRule.Tags, rule.Tags)
			assert.Equal(t, createRule.Output.Recipients.Emails, rule.Output.Recipients.Emails)
			assert.Equal(t, createRule.SubComponents[0].QueryDefinition.Query, rule.SubComponents[0].QueryDefinition.Query)
			assert.Equal(t, createRule.Output.SuppressNotificationsMinutes, rule.Output.SuppressNotificationsMinutes)
			assert.Equal(t, createRule.Output.Type, rule.Output.Type)
			assert.Equal(t, createRule.SearchTimeFrameMinutes, rule.SearchTimeFrameMinutes)
			assert.Equal(t, len(createRule.SubComponents), len(rule.SubComponents))
			assert.Equal(t, createRule.SubComponents[0].QueryDefinition.Query, rule.SubComponents[0].QueryDefinition.Query)
			assert.True(t, reflect.DeepEqual(createRule.SubComponents[0].QueryDefinition.Filters.Bool, rule.SubComponents[0].QueryDefinition.Filters.Bool))
			assert.Equal(t, createRule.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts, rule.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts)
			assert.Equal(t, createRule.SubComponents[0].Trigger.Operator, rule.SubComponents[0].Trigger.Operator)
			assert.Equal(t, createRule.SubComponents[0].Trigger.SeverityThresholdTiers, rule.SubComponents[0].Trigger.SeverityThresholdTiers)
			assert.Equal(t, len(createRule.SubComponents[0].Output.Columns), len(rule.SubComponents[0].Output.Columns))
			assert.Equal(t, createRule.SubComponents[0].Output.Columns[0].Sort, rule.SubComponents[0].Output.Columns[0].Sort)
			assert.Equal(t, createRule.SubComponents[0].Output.Columns[0].FieldName, rule.SubComponents[0].Output.Columns[0].FieldName)
			assert.Equal(t, *createRule.Enabled, rule.Enabled)
		}
	}
}

func TestIntegrationSecurityRules_CreateRuleNoTitle(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = ""
		rule, err := underTest.CreateSecurityRule(createRule)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Equal(t, "title must be set", err.Error())
	}
}

func TestIntegrationSecurityRules_CreateRuleNoSubComponent(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test no sub component"
		createRule.SubComponents = nil
		rule, err := underTest.CreateSecurityRule(createRule)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Equal(t, "sub components must be set", err.Error())
	}
}

func TestIntegrationSecurityRules_CreateRuleInvalidOutput(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test invalid output type"
		createRule.Output.Type = "invalid output type"
		rule, err := underTest.CreateSecurityRule(createRule)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "invalid output type.")
	}
}

func TestIntegrationSecurityRules_CreateRuleInvalidAggregationType(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test invalid aggregation type"
		createRule.SubComponents[0].QueryDefinition.Aggregation.AggregationType = "invalid agg type"
		rule, err := underTest.CreateSecurityRule(createRule)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "invalid aggregation type.")
	}
}

func TestIntegrationSecurityRules_CreateRuleInvalidTriggerOperator(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test invalid trigger operator"
		createRule.SubComponents[0].Trigger.Operator = "invalid trigger operator"
		rule, err := underTest.CreateSecurityRule(createRule)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "invalid operator.")
	}
}

func TestIntegrationSecurityRules_CreateRuleInvalidThresholdTier(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test invalid threshold tier"
		createRule.SubComponents[0].Trigger.SeverityThresholdTiers = map[string]float32{"SOMETHING": 15}
		rule, err := underTest.CreateSecurityRule(createRule)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "invalid severity threshold tiers.")
	}
}

func TestIntegrationSecurityRules_CreateRuleInvalidOutputSort(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test invalid output sort"
		createRule.SubComponents[0].Output.Columns[0].Sort = "invalid sort"
		rule, err := underTest.CreateSecurityRule(createRule)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "invalid sort.")
	}
}

func TestIntegrationSecurityRules_CreateRuleInvalidCorrelationOperator(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test invalid correlation operator"
		createRule.Correlations.CorrelationOperators = []string{"invalid"}
		rule, err := underTest.CreateSecurityRule(createRule)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "invalid correlation operator.")
	}
}

package security_rules_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSecurityRules_UpdateRule(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test update before"

		rule, err := underTest.CreateSecurityRule(createRule)
		if assert.NoError(t, err) && assert.NotNil(t, rule) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSecurityRule(rule.Id)
			createRule.Title = "test update AFTER"
			updated, err := underTest.UpdateSecurityRule(rule.Id, createRule)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, createRule.Title, updated.Title)
		}
	}
}

func TestIntegrationSecurityRules_UpdateRuleIdNotExist(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		updateRule := getCreateUpdateRule()
		updateRule.Title = "test update id not exist"
		updated, err := underTest.UpdateSecurityRule(int32(1234), updateRule)
		assert.Error(t, err)
		assert.Nil(t, updated)
		assert.Contains(t, err.Error(), "failed with missing security rule")
	}
}

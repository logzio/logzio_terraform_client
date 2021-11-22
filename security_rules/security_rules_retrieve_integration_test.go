package security_rules_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSecurityRules_RetrieveRule(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test retrieve rule"

		rule, err := underTest.CreateSecurityRule(createRule)
		if assert.NoError(t, err) && assert.NotNil(t, rule) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSecurityRule(rule.Id)
			getRule, err := underTest.RetrieveSecurityRule(rule.Id)
			assert.NoError(t, err)
			assert.NotNil(t, getRule)
			assert.Equal(t, rule.Id, getRule.Id)

		}
	}
}

func TestIntegrationSecurityRules_RetrieveRuleIdNotExist(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		getRule, err := underTest.RetrieveSecurityRule(int32(12345))
		assert.Error(t, err)
		assert.Nil(t, getRule)
		assert.Contains(t, err.Error(), "failed with missing security rule")
	}
}

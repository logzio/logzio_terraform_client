package security_rules_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSecurityRules_DisableRule(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test disable"
		rule, err := underTest.CreateSecurityRule(createRule)

		if assert.NoError(t, err) && assert.NotNil(t, rule) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSecurityRule(rule.Id)
			err = underTest.DisableSecurityRule(rule.Id)
			assert.NoError(t, err)
			// Double check if we disabled the rule
			time.Sleep(2 * time.Second) // give time for the rule to be disabled
			ruleGet, err := underTest.RetrieveSecurityRule(rule.Id)
			assert.NoError(t, err)
			assert.NotNil(t, ruleGet)
			assert.False(t, ruleGet.Enabled)
		}
	}
}

// TODO - negative test for non-existing id

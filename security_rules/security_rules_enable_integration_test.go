package security_rules_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSecurityRules_EnableRule(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test enable"
		*createRule.Enabled = false
		rule, err := underTest.CreateSecurityRule(createRule)

		if assert.NoError(t, err) && assert.NotNil(t, rule) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSecurityRule(rule.Id)
			err = underTest.EnableSecurityRule(rule.Id)
			assert.NoError(t, err)
			// Double check if we disabled the rule
			time.Sleep(2 * time.Second) // give time for the rule to be disabled
			ruleGet, err := underTest.RetrieveSecurityRule(rule.Id)
			assert.NoError(t, err)
			assert.NotNil(t, ruleGet)
			assert.True(t, ruleGet.Enabled)
		}
	}
}

func TestIntegrationSecurityRules_EnableRuleIdNotExists(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		err = underTest.EnableSecurityRule(1234)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing security rule")
	}
}

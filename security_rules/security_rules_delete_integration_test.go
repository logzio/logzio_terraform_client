package security_rules_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSecurityRules_DeleteRule(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test_delete"
		rule, err := underTest.CreateSecurityRule(createRule)
		if assert.NoError(t, err) && assert.NotNil(t, rule) {
			time.Sleep(4 * time.Second)
			defer func() {
				rule, err = underTest.DeleteSecurityRule(rule.Id)
				assert.NoError(t, err)
				assert.NotNil(t, rule)
			}()
		}
	}
}

func TestIntegrationSecurityRules_DeleteRuleIdNotExist(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()
	if assert.NoError(t, err) {
		rule, err := underTest.DeleteSecurityRule(int32(1234))
		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "failed with missing security rule")
	}
}

package security_rules_test

import (
	"github.com/logzio/logzio_terraform_client/security_rules"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationSecurityRules_RetrieveFilteredRules(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		createRule := getCreateUpdateRule()
		createRule.Title = "test retrieve filtered"
		rule, err := underTest.CreateSecurityRule(createRule)

		if assert.NoError(t, err) && assert.NotNil(t, rule) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteSecurityRule(rule.Id)
			retrieve := getRetrieveFilteredRequest(createRule.Title)
			rules, err := underTest.RetrieveFilteredSecurityRules(retrieve)
			assert.NoError(t, err)
			assert.NotNil(t, rules)
			assert.NotZero(t, rules.Total)
			assert.NotZero(t, len(rules.Results))
		}
	}
}

func TestIntegrationSecurityRules_RetrieveFilteredRulesNotFound(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		retrieve := getRetrieveFilteredRequest("foo bar")
		rules, err := underTest.RetrieveFilteredSecurityRules(retrieve)
		assert.NoError(t, err)
		assert.NotNil(t, rules)
		assert.Zero(t, rules.Total)
		assert.Zero(t, len(rules.Results))
	}
}

func TestIntegrationSecurityRules_RetrieveFilteredRulesPagination(t *testing.T) {
	underTest, err := setupSecurityRulesIntegrationTest()

	if assert.NoError(t, err) {
		retrieve := security_rules.RetrieveFiltered{
			Filter: security_rules.SecurityRuleFilter{
				Tags: []string{"access"},
			},
			Pagination: security_rules.SecurityRulePagination{
				PageNumber: 1,
			},
		}

		rules, err := underTest.RetrieveFilteredSecurityRules(retrieve)
		if assert.NoError(t, err) && assert.NotNil(t, rules) && assert.NotZero(t, rules.Total) {
			total := rules.Total
			retrievedRulesNum := int32(len(rules.Results))
			for retrievedRulesNum < total {
				retrieve.Pagination.PageNumber += 1
				rules, err = underTest.RetrieveFilteredSecurityRules(retrieve)
				assert.NoError(t, err)
				assert.NotNil(t, rules)
				assert.NotZero(t, len(rules.Results))
				retrievedRulesNum += int32(len(rules.Results))
			}

			assert.Equal(t, total, retrievedRulesNum)
		}
	}
}

func getRetrieveFilteredRequest(search string) security_rules.RetrieveFiltered {
	filter := security_rules.SecurityRuleFilter{
		Search: search,
	}

	return security_rules.RetrieveFiltered{
		Filter: filter,
	}
}

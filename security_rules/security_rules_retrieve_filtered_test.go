package security_rules_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/security_rules"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSecurityRules_RetrieveFilteredRules(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(securityRulesApiBasePath+"/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("retrieve_filtered_rules.json"))
		})

		tagToSearch := "access"
		retrieve := security_rules.RetrieveFiltered{
			Filter: security_rules.SecurityRuleFilter{
				Tags: []string{tagToSearch},
			},
			Pagination: security_rules.SecurityRulePagination{
				PageNumber: 1,
				PageSize:   25,
			},
		}

		rules, err := underTest.RetrieveFilteredSecurityRules(retrieve)

		assert.NoError(t, err)
		assert.NotNil(t, rules)
		assert.NotZero(t, rules.Total)
		assert.NotZero(t, len(rules.Results))
		for _, result := range rules.Results {
			assert.Contains(t, result.Tags, tagToSearch)
		}
	}
}

func TestSecurityRules_RetrieveFilteredRulesNoResults(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(securityRulesApiBasePath+"/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("retrieve_filtered_rules_no_results.json"))
		})

		tagToSearch := "access"
		retrieve := security_rules.RetrieveFiltered{
			Filter: security_rules.SecurityRuleFilter{
				Tags: []string{tagToSearch},
			},
			Pagination: security_rules.SecurityRulePagination{
				PageNumber: 1,
				PageSize:   25,
			},
		}

		rules, err := underTest.RetrieveFilteredSecurityRules(retrieve)

		assert.NoError(t, err)
		assert.NotNil(t, rules)
		assert.Zero(t, rules.Total)
		assert.Zero(t, len(rules.Results))
	}
}

func TestSecurityRules_RetrieveFilteredRulesApiFail(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(securityRulesApiBasePath+"/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
		})

		tagToSearch := "access"
		retrieve := security_rules.RetrieveFiltered{
			Filter: security_rules.SecurityRuleFilter{
				Tags: []string{tagToSearch},
			},
			Pagination: security_rules.SecurityRulePagination{
				PageNumber: 1,
				PageSize:   25,
			},
		}

		rules, err := underTest.RetrieveFilteredSecurityRules(retrieve)

		assert.Error(t, err)
		assert.Nil(t, rules)
	}
}

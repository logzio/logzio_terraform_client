package security_rules_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestSecurityRules_RetrieveRule(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("retrieve_rule.json"))
		})

		rule, err := underTest.RetrieveSecurityRule(ruleId)

		assert.NoError(t, err)
		assert.NotNil(t, rule)
		assert.Equal(t, ruleId, rule.Id)
	}
}

func TestSecurityRules_RetrieveRuleIdNotFound(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusNotFound)
		})

		rule, err := underTest.RetrieveSecurityRule(ruleId)

		assert.Error(t, err)
		assert.Nil(t, rule)
		assert.Contains(t, err.Error(), "failed with missing security rule")
	}
}

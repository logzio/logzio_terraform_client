package security_rules_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestSecurityRules_DeleteRule(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("delete_rule.json"))
		})

		deleteRule, err := underTest.DeleteSecurityRule(ruleId)
		assert.NoError(t, err)
		assert.NotNil(t, deleteRule)
		assert.Equal(t, ruleId, deleteRule.Id)
	}
}

func TestSecurityRules_DeleteRuleIdNotFound(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusNotFound)
		})

		deleteRule, err := underTest.DeleteSecurityRule(ruleId)
		assert.Error(t, err)
		assert.Nil(t, deleteRule)
		assert.Contains(t, err.Error(), "failed with missing security rule")
	}
}

func TestSecurityRules_DeleteRuleApiFail(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusNotFound)
		})

		deleteRule, err := underTest.DeleteSecurityRule(ruleId)
		assert.Error(t, err)
		assert.Nil(t, deleteRule)
	}
}

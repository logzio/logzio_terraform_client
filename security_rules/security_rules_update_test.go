package security_rules_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/security_rules"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestSecurityRules_UpdateRule(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target security_rules.CreateUpdateSecurityRule
			err = json.Unmarshal(jsonBytes, &target)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotEmpty(t, target.SubComponents)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("update_rule.json"))
		})

		update := getCreateUpdateRule()
		update.Title = "after update"
		updatedRule, err := underTest.UpdateSecurityRule(ruleId, update)
		assert.NoError(t, err)
		assert.NotNil(t, updatedRule)
		assert.Equal(t, update.Title, updatedRule.Title)
	}
}

func TestSecurityRules_UpdateRuleIdNotFound(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			w.WriteHeader(http.StatusNotFound)
		})

		update := getCreateUpdateRule()
		update.Title = "after update"
		updatedRule, err := underTest.UpdateSecurityRule(ruleId, update)
		assert.Error(t, err)
		assert.Nil(t, updatedRule)
		assert.Contains(t, err.Error(), "failed with missing security rule")
	}
}

func TestSecurityRules_UpdateRuleApiFail(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			w.WriteHeader(http.StatusInternalServerError)
		})

		update := getCreateUpdateRule()
		update.Title = "after update"
		updatedRule, err := underTest.UpdateSecurityRule(ruleId, update)
		assert.Error(t, err)
		assert.Nil(t, updatedRule)
	}
}

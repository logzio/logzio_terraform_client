package security_rules_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/security_rules"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestSecurityRules_CreateSecurityRule(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(securityRulesApiBasePath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target security_rules.CreateUpdateSecurityRule
			err = json.Unmarshal(jsonBytes, &target)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotEmpty(t, target.SubComponents[0])
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_rule.json"))
		})
	}

	createRule := getCreateUpdateRule()
	rule, err := underTest.CreateSecurityRule(createRule)
	assert.NoError(t, err)
	assert.NotNil(t, rule)
	assert.Equal(t, createRule.Title, rule.Title)
	assert.Equal(t, createRule.Description, rule.Description)
	assert.Equal(t, createRule.Tags, rule.Tags)
	assert.Equal(t, createRule.Output.Recipients.Emails, rule.Output.Recipients.Emails)
	assert.Equal(t, createRule.SubComponents[0].QueryDefinition.Query, rule.SubComponents[0].QueryDefinition.Query)
	assert.Equal(t, createRule.Output.SuppressNotificationsMinutes, rule.Output.SuppressNotificationsMinutes)
	assert.Equal(t, createRule.Output.Type, rule.Output.Type)
	assert.Equal(t, createRule.SearchTimeFrameMinutes, rule.SearchTimeFrameMinutes)
	assert.Equal(t, len(createRule.SubComponents), len(rule.SubComponents))
	assert.Equal(t, createRule.SubComponents[0].QueryDefinition.Query, rule.SubComponents[0].QueryDefinition.Query)
	assert.True(t, reflect.DeepEqual(createRule.SubComponents[0].QueryDefinition.Filters.Bool, rule.SubComponents[0].QueryDefinition.Filters.Bool))
	assert.Equal(t, createRule.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts, rule.SubComponents[0].QueryDefinition.ShouldQueryOnAllAccounts)
	assert.Equal(t, createRule.SubComponents[0].Trigger.Operator, rule.SubComponents[0].Trigger.Operator)
	assert.Equal(t, createRule.SubComponents[0].Trigger.SeverityThresholdTiers, rule.SubComponents[0].Trigger.SeverityThresholdTiers)
	assert.Equal(t, len(createRule.SubComponents[0].Output.Columns), len(rule.SubComponents[0].Output.Columns))
	assert.Equal(t, createRule.SubComponents[0].Output.Columns[0].Sort, rule.SubComponents[0].Output.Columns[0].Sort)
	assert.Equal(t, createRule.SubComponents[0].Output.Columns[0].FieldName, rule.SubComponents[0].Output.Columns[0].FieldName)
	assert.Equal(t, *createRule.Enabled, rule.Enabled)
}

func TestSecurityRules_CreateSecurityRuleAPIFail(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(securityRulesApiBasePath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target security_rules.CreateUpdateSecurityRule
			err = json.Unmarshal(jsonBytes, &target)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotEmpty(t, target.SubComponents[0])
			w.WriteHeader(http.StatusInternalServerError)
		})

		createRule := getCreateUpdateRule()
		rule, err := underTest.CreateSecurityRule(createRule)
		assert.Error(t, err)
		assert.Nil(t, rule)
	}
}

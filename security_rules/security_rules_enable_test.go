package security_rules_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestSecurityRules_EnableRule(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusNoContent)
		})

		err = underTest.EnableSecurityRule(ruleId)
		assert.NoError(t, err)
	}
}

func TestSecurityRules_EnableRuleIdNotFound(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusNotFound)
		})

		err = underTest.EnableSecurityRule(ruleId)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing security rule")
	}
}

func TestSecurityRules_EnableRuleApiFail(t *testing.T) {
	underTest, teardown, err := setupSecurityRulesTest()
	defer teardown()

	if assert.NoError(t, err) {
		ruleId := int32(12345)
		mux.HandleFunc(securityRulesApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(ruleId), 10))
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
		})

		err = underTest.EnableSecurityRule(ruleId)
		assert.Error(t, err)
	}
}

package log_shipping_tokens_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/log_shipping_tokens"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestLogShippingTokens_RetrieveLogShippingTokens(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target log_shipping_tokens.RetrieveLogShippingTokensRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Filter.Enabled)
			assert.Contains(t, []string{strconv.FormatBool(true), strconv.FormatBool(false)}, target.Filter.Enabled)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("retrieve_log_shipping_tokens.json"))
			w.WriteHeader(http.StatusOK)
		})
	}

	retrieveRequest := getCreateRetrieveTokensRequest()
	tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
	assert.NoError(t, err)
	assert.NotNil(t, tokens)
	assert.NotZero(t, tokens.Total)
	assert.NotZero(t, len(tokens.Results))
}

func TestLogShippingTokens_RetrieveLogShippingTokensAPIFailed(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target log_shipping_tokens.RetrieveLogShippingTokensRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Filter.Enabled)
			assert.Contains(t, []string{strconv.FormatBool(true), strconv.FormatBool(false)}, target.Filter.Enabled)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("retrieve_log_shipping_tokens_failed.txt"))
			w.WriteHeader(http.StatusInternalServerError)
		})
	}

	retrieveRequest := getCreateRetrieveTokensRequest()
	tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
	assert.Error(t, err)
	assert.Nil(t, tokens)
}

func TestLogShippingTokens_RetrieveLogShippingTokensNoFilter(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	retrieveRequest := getCreateRetrieveTokensRequest()
	retrieveRequest.Filter.Enabled = ""
	tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
	assert.Error(t, err)
	assert.Nil(t, tokens)
}

func TestLogShippingTokens_RetrieveLogShippingTokensInvalidFilter(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	retrieveRequest := getCreateRetrieveTokensRequest()
	retrieveRequest.Filter.Enabled = "invalid_value"
	tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
	assert.Error(t, err)
	assert.Nil(t, tokens)
}

func TestLogShippingTokens_RetrieveLogShippingTokensInvalidSortField(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	retrieveRequest := getCreateRetrieveTokensRequest()
	retrieveRequest.Sort[0].Field = "invalid_value"
	tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
	assert.Error(t, err)
	assert.Nil(t, tokens)
}

func TestLogShippingTokens_RetrieveLogShippingTokensInvalidSortDescending(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	retrieveRequest := getCreateRetrieveTokensRequest()
	retrieveRequest.Sort[0].Descending = "invalid_value"
	tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
	assert.Error(t, err)
	assert.Nil(t, tokens)
}

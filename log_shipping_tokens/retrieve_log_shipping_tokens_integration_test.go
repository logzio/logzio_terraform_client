package log_shipping_tokens_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationLogShippingTokens_RetrieveTokens(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		retrieveRequest := getCreateRetrieveTokensRequest()
		tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
		assert.NoError(t, err)
		assert.NotNil(t, tokens)
		assert.NotZero(t, tokens.Total)
		assert.NotZero(t, len(tokens.Results))
	}
}

func TestIntegrationLogShippingTokens_RetrieveTokensNoFilter(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		retrieveRequest := getCreateRetrieveTokensRequest()
		retrieveRequest.Filter.Enabled = ""
		tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
		assert.Error(t, err)
		assert.Nil(t, tokens)
	}
}

func TestIntegrationLogShippingTokens_RetrieveTokensInvalidFilter(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		retrieveRequest := getCreateRetrieveTokensRequest()
		retrieveRequest.Filter.Enabled = "invalid_value"
		tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
		assert.Error(t, err)
		assert.Nil(t, tokens)
	}
}

func TestIntegrationLogShippingTokens_RetrieveTokensInvalidSortField(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		retrieveRequest := getCreateRetrieveTokensRequest()
		retrieveRequest.Sort[0].Field = "invalid"
		tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
		assert.Error(t, err)
		assert.Nil(t, tokens)
	}
}

func TestIntegrationLogShippingTokens_RetrieveTokensInvalidSortDescending(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		retrieveRequest := getCreateRetrieveTokensRequest()
		retrieveRequest.Sort[0].Descending = "invalid"
		tokens, err := underTest.RetrieveLogShippingTokens(retrieveRequest)
		assert.Error(t, err)
		assert.Nil(t, tokens)
	}
}
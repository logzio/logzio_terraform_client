package log_shipping_tokens_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationLogShippingTokens_GetToken(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		if assert.NoError(t, err) && assert.NotNil(t, token) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteLogShippingToken(token.Id)
			tokenFromGet, err := underTest.GetLogShippingToken(token.Id)
			assert.NoError(t, err)
			assert.NotNil(t, tokenFromGet)
			assert.Equal(t, token.Id, tokenFromGet.Id)
		}
	}
}

func TestIntegrationLogShippingTokens_GetTokenNotExist(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		tokenFromGet, err := underTest.GetLogShippingToken(int32(123456))
		assert.Error(t, err)
		assert.Nil(t, tokenFromGet)
	}
}

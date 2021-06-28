package log_shipping_tokens_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationLogShippingTokens_CreateToken(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, token) {
			assert.Equal(t, createToken.Name, token.Name)
			defer underTest.DeleteLogShippingToken(token.Id)
		}
	}
}

func TestIntegrationLogShippingTokens_CreateTokenInvalidName(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		createToken.Name = ""
		token, err := underTest.CreateLogShippingToken(createToken)

		assert.Error(t, err)
		assert.Nil(t, token)
	}
}

func TestIntegrationLogShippingTokens_CreateTokenAlreadyExist(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, token) {
			defer underTest.DeleteLogShippingToken(token.Id)
			assert.Equal(t, createToken.Name, token.Name)
			newToken, err := underTest.CreateLogShippingToken(createToken)
			assert.Error(t, err)
			assert.Nil(t, newToken)
		}
	}
}
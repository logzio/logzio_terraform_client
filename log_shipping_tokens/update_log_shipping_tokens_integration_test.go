package log_shipping_tokens_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestIntegrationLogShippingTokens_UpdateTokenName(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		if assert.NoError(t, err) && assert.NotNil(t, token) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteLogShippingToken(token.Id)
			assert.Equal(t, createToken.Name, token.Name)

			updateRequest := getUpdateLogShippingToken()
			updateRequest.Name = "change_name"
			updatedToken, err := underTest.UpdateLogShippingToken(token.Id, updateRequest)
			assert.NoError(t, err)
			assert.NotNil(t, updatedToken)
			assert.Equal(t, updateRequest.Name, updatedToken.Name)
			assert.Equal(t, token.Id, updatedToken.Id)
		}
	}
}

func TestIntegrationLogShippingTokens_UpdateTokenEnabled(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		if assert.NoError(t, err) && assert.NotNil(t, token) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteLogShippingToken(token.Id)
			assert.Equal(t, createToken.Name, token.Name)

			updateRequest := getUpdateLogShippingToken()
			updateRequest.Enabled = strconv.FormatBool(false)
			updatedToken, err := underTest.UpdateLogShippingToken(token.Id, updateRequest)
			assert.NoError(t, err)
			assert.NotNil(t, updatedToken)
			assert.Equal(t, updateRequest.Enabled, strconv.FormatBool(updatedToken.Enabled))
			assert.Equal(t, token.Id, updatedToken.Id)
		}
	}
}

func TestIntegrationLogShippingTokens_UpdateTokenMissingToken(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		updateRequest := getUpdateLogShippingToken()
		token, err := underTest.UpdateLogShippingToken(int32(123), updateRequest)
		assert.Error(t, err)
		assert.Nil(t, token)
	}
}

func TestIntegrationLogShippingTokens_UpdateTokenNoName(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		if assert.NoError(t, err) && assert.NotNil(t, token) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteLogShippingToken(token.Id)
			assert.Equal(t, createToken.Name, token.Name)

			updateRequest := getUpdateLogShippingToken()
			updateRequest.Name = ""
			updatedToken, err := underTest.UpdateLogShippingToken(token.Id, updateRequest)
			assert.Error(t, err)
			assert.Nil(t, updatedToken)
		}
	}
}

func TestIntegrationLogShippingTokens_UpdateTokenInvalidName(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		if assert.NoError(t, err) && assert.NotNil(t, token) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteLogShippingToken(token.Id)
			assert.Equal(t, createToken.Name, token.Name)

			updateRequest := getUpdateLogShippingToken()
			updateRequest.Name = ""
			updatedToken, err := underTest.UpdateLogShippingToken(token.Id, updateRequest)
			assert.Error(t, err)
			assert.Nil(t, updatedToken)
		}
	}
}

func TestIntegrationLogShippingTokens_UpdateTokenInvalidEnabled(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		if assert.NoError(t, err) && assert.NotNil(t, token) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteLogShippingToken(token.Id)
			assert.Equal(t, createToken.Name, token.Name)

			updateRequest := getUpdateLogShippingToken()
			updateRequest.Enabled = "invalid_val"
			updatedToken, err := underTest.UpdateLogShippingToken(token.Id, updateRequest)
			assert.Error(t, err)
			assert.Nil(t, updatedToken)
		}
	}
}

func TestIntegrationLogShippingTokens_UpdateTokenNoEnabled(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		if assert.NoError(t, err) && assert.NotNil(t, token) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteLogShippingToken(token.Id)
			assert.Equal(t, createToken.Name, token.Name)

			updateRequest := getUpdateLogShippingToken()
			updateRequest.Enabled = ""
			updatedToken, err := underTest.UpdateLogShippingToken(token.Id, updateRequest)
			assert.Error(t, err)
			assert.Nil(t, updatedToken)
		}
	}
}
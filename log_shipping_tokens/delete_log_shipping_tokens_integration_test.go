package log_shipping_tokens_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationLogShippingTokens_DeleteToken(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)

		if assert.NoError(t, err) && assert.NotNil(t, token) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteLogShippingToken(token.Id)
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationLogShippingTokens_DeleteTokenNotFound(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteLogShippingToken(int32(123456))
		assert.Error(t, err)
	}
}

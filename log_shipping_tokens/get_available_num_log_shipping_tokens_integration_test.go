package log_shipping_tokens_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationLogShippingTokens_GetLimitsToken(t *testing.T) {
	underTest, err := setupLogShippingTokensIntegrationTest()
	if assert.NoError(t, err) {
		limits, err := underTest.GetLogShippingLimitsToken()
		assert.NoError(t, err)
		assert.NotNil(t, limits)
	}
}

package log_shipping_tokens_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLogShippingTokens_GetLimitsLogShippingToken(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/limits", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("get_limits_log_shipping_token.json"))
		})
	}

	limits, err := underTest.GetLogShippingLimitsToken()
	assert.NoError(t, err)
	assert.NotNil(t, limits)
}

func TestLogShippingTokens_GetLimitsLogShippingTokenAPIFailed(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/limits", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("get_limits_log_shipping_token.json"))
		})
	}

	limits, err := underTest.GetLogShippingLimitsToken()
	assert.Error(t, err)
	assert.Nil(t, limits)
}

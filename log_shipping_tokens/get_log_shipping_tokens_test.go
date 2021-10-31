package log_shipping_tokens_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestLogShippingTokens_GetLogShippingToken(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	tokenId := int64(123456)

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(tokenId, 10))
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("get_log_shipping_token.json"))
		})
	}

	token, err := underTest.GetLogShippingToken(int32(tokenId))
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, int32(tokenId), token.Id)
}

func TestLogShippingTokens_GetLogShippingTokenNotExist(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	tokenId := int64(123456)

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(tokenId, 10))
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("get_log_shipping_token_not_exist.txt"))
		})
	}

	token, err := underTest.GetLogShippingToken(int32(tokenId))
	assert.Error(t, err)
	assert.Nil(t, token)
	assert.Contains(t, err.Error(), "failed with missing log shipping token")
}

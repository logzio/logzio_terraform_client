package log_shipping_tokens_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestLogShippingTokens_DeleteLogShippingToken(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	tokenId := int64(123456)

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(tokenId, 10))
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("delete_log_shipping_token.txt"))
			w.WriteHeader(http.StatusOK)
		})
	}

	err = underTest.DeleteLogShippingToken(int32(tokenId))
	assert.NoError(t, err)
}

func TestLogShippingTokens_DeleteLogShippingTokenMissingToken(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	tokenId := int64(123456)

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(tokenId, 10))
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("delete_log_shipping_token_not_exist.txt"))
			w.WriteHeader(http.StatusNotFound)
		})
	}

	err = underTest.DeleteLogShippingToken(int32(tokenId))
	assert.NoError(t, err)
}

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

func TestLogShippingTokens_UpdateLogShippingToken(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target log_shipping_tokens.UpdateLogShippingToken
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Name)
			assert.NotEmpty(t, target.Enabled)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("update_log_shipping_token.json"))
			w.WriteHeader(http.StatusOK)
		})
	}

	tokenId := int32(123456)
	updateRequest := getUpdateLogShippingToken()
	updatedToken, err := underTest.UpdateLogShippingToken(tokenId, updateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, updatedToken)
	assert.Equal(t, updateRequest.Name, updatedToken.Name)
	assert.Equal(t, updateRequest.Enabled, strconv.FormatBool(updatedToken.Enabled))
}

func TestLogShippingTokens_UpdateLogShippingTokenIdNotFound(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target log_shipping_tokens.UpdateLogShippingToken
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Enabled)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("update_log_shipping_token_failed.json"))
			w.WriteHeader(http.StatusNotFound)
		})
	}

	updateRequest := getUpdateLogShippingToken()
	updatedToken, err := underTest.UpdateLogShippingToken(int32(123), updateRequest)
	assert.Error(t, err)
	assert.Nil(t, updatedToken)
}

func TestLogShippingTokens_UpdateLogShippingTokenNoName(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		updateRequest := getUpdateLogShippingToken()
		updateRequest.Name = ""
		updatedToken, err := underTest.UpdateLogShippingToken(int32(123456), updateRequest)
		assert.Error(t, err)
		assert.Nil(t, updatedToken)
	}
}

func TestLogShippingTokens_UpdateLogShippingTokenNoEnabled(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		updateRequest := getUpdateLogShippingToken()
		updateRequest.Enabled = ""
		updatedToken, err := underTest.UpdateLogShippingToken(int32(123456), updateRequest)
		assert.Error(t, err)
		assert.Nil(t, updatedToken)
	}
}

func TestLogShippingTokens_UpdateLogShippingTokenInvalidEnabled(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		updateRequest := getUpdateLogShippingToken()
		updateRequest.Enabled = "invalid_enabled"
		updatedToken, err := underTest.UpdateLogShippingToken(int32(123456), updateRequest)
		assert.Error(t, err)
		assert.Nil(t, updatedToken)
	}
}
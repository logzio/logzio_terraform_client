package log_shipping_tokens_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/log_shipping_tokens"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLogShippingTokens_CreateLogShippingToken(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target log_shipping_tokens.CreateLogShippingToken
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Name)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_log_shipping_token.json"))
			w.WriteHeader(http.StatusOK)
		})

		createToken := getCreateLogShippingToken()
		token, err := underTest.CreateLogShippingToken(createToken)
		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.Equal(t, createToken.Name, token.Name)
	}
}

func TestLogShippingTokens_CreateLogShippingTokenAPIFail(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/tokens", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_log_shipping_token_failed.txt"))
		})
	}

	_, err = underTest.CreateLogShippingToken(getCreateLogShippingToken())
	assert.Error(t, err)
}

func TestLogShippingTokens_CreateLogShippingTokenNoName(t *testing.T) {
	underTest, err, teardown := setupLogShippingTokenTest()
	defer teardown()

	createToken := getCreateLogShippingToken()
	createToken.Name = ""
	token, err := underTest.CreateLogShippingToken(createToken)
	assert.Error(t, err)
	assert.Nil(t, token)
}
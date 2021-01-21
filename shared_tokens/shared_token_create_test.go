package shared_tokens_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/shared_tokens"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSubAccount_CreateValidSubAccount(t *testing.T) {
	underTest, err, teardown := setupSharedTokensTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/shared-tokens/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "name")
		assert.Contains(t, target, "token")
		assert.Contains(t, target, "filters")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_sharedtoken.json"))
	})

	filterIds := make([]int32, 2)
	filterIds[0] = 123
	filterIds[1] = 456

	s := shared_tokens.SharedToken{
		Name:                 "test token",
		Token: "6c36edf51-cf93883aa35-5bc6ce6-7bcfe60d87",
		Id: 1234,
		FilterIds: filterIds,
	}

	sharedToken, err := underTest.CreateSharedToken(s)
	assert.NoError(t, err)
	assert.NotNil(t, sharedToken)
	assert.NotZero(t, sharedToken.Id)
}

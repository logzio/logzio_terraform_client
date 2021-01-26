package shared_tokens_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/shared_tokens"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestSharedToken_UpdateValidSharedToken(t *testing.T) {
	underTest, err, teardown := setupSharedTokensTest()
	assert.NoError(t, err)
	defer teardown()

	sharedTokenId := int64(123456)

	mux.HandleFunc("/v1/shared-tokens/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(sharedTokenId, 10))

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "name")
		assert.Contains(t, target, "filters")

		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_shared_token.json"))
	})

	s := shared_tokens.SharedToken{
		Name:      "new test",
		FilterIds: []int32{2345, 6789},
	}

	err = underTest.UpdateSharedToken(sharedTokenId, s)
	assert.NoError(t, err)
}

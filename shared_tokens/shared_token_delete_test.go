package shared_tokens_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestSubAccount_DeleteValidSharedToken(t *testing.T) {
	underTest, err, teardown := setupSharedTokensTest()
	assert.NoError(t, err)
	defer teardown()

	sharedTokenId := int64(123456)

	mux.HandleFunc("/v1/shared-tokens/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(sharedTokenId), 10))
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("delete_shared_token.json"))
	})

	err = underTest.DeleteSharedToken(sharedTokenId)
	assert.NoError(t, err)
}

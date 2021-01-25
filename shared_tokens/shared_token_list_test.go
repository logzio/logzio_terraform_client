package shared_tokens_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSharedToken_ListSharedtokens(t *testing.T) {
	underTest, err, teardown := setupSharedTokensTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/shared-tokens", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_shared_tokens.json"))
	})

	sharedTokens, err := underTest.ListSharedTokens()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(sharedTokens))
}
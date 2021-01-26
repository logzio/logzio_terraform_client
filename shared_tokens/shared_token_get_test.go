package shared_tokens_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestSubAccount_GetValidSharedToken(t *testing.T) {
	underTest, err, teardown := setupSharedTokensTest()
	assert.NoError(t, err)
	defer teardown()

	sharedTokenId := int32(123456)

	mux.HandleFunc("/v1/shared-tokens/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(sharedTokenId), 10))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_shared_token.json"))
	})

	sharedToken, err := underTest.GetSharedToken(sharedTokenId)
	assert.NoError(t, err)
	assert.NotNil(t, sharedToken)

	assert.Equal(t, sharedTokenId, sharedToken.Id)
	assert.Equal(t, "test token", sharedToken.Name)
	assert.Equal(t, "6c36edf51-cf93883aa35-5bc6ce6-7bcfe60d87", sharedToken.Token)
	filters := sharedToken.FilterIds
	assert.Equal(t, 2, len(filters))
	assert.Equal(t, int32(123), filters[0])
}

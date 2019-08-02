package sub_accounts_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestSubAccount_DeleteValidSubAccount(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(1234567)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(subAccountId), 10))
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("delete_subaccount.json"))
	})

	err = underTest.DeleteSubAccount(subAccountId)
	assert.NoError(t, err)
}

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
		assert.Contains(t, r.URL.String(), strconv.FormatInt(subAccountId, 10))
		w.WriteHeader(http.StatusNoContent)
	})

	err = underTest.DeleteSubAccount(subAccountId)
	assert.NoError(t, err)
}

func TestSubAccount_DeleteValidSubAccountNotFound(t *testing.T) {
	underTest, err, teardown := setupSubAccountsTest()
	assert.NoError(t, err)
	defer teardown()

	subAccountId := int64(12345)

	mux.HandleFunc("/v1/account-management/time-based-accounts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(subAccountId, 10))
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("delete_subaccount_failed.txt"))
	})

	err = underTest.DeleteSubAccount(subAccountId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed with missing sub account")
}

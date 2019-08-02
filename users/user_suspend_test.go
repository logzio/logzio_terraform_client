package users_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestUsers_SuspendUser(t *testing.T) {
	accountId := int64(123456)

	underTest, err, teardown := setupUsersTest()
	defer teardown()

	mux.HandleFunc("/v1/user-management/suspend/", func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.String(), strconv.FormatInt(accountId, 10))
		assert.Equal(t, http.MethodPost, r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	success, err := underTest.SuspendUser(accountId)
	assert.True(t, success)
	assert.NoError(t, err)
}

func TestUsers_UnSuspendUser(t *testing.T) {
	accountId := int64(123456)

	underTest, err, teardown := setupUsersTest()
	defer teardown()

	mux.HandleFunc("/v1/user-management/unsuspend/", func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.String(), strconv.FormatInt(accountId, 10))
		assert.Equal(t, http.MethodPost, r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	success, err := underTest.UnSuspendUser(accountId)
	assert.True(t, success)
	assert.NoError(t, err)
}

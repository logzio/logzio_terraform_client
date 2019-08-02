package users_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestUsers_GetUser(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	assert.NoError(t, err)
	defer teardown()

	accountId := int64(1234567)

	mux.HandleFunc("/v1/user-management/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(accountId, 10))

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_user.json"))
	})

	user, err := underTest.GetUser(accountId)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUsers_GetUserNotExist(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	assert.NoError(t, err)
	defer teardown()

	accountId := int64(1234567)

	mux.HandleFunc("/v1/user-management/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(accountId, 10))

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_user_not_exist.json"))
	})

	user, err := underTest.GetUser(accountId)
	assert.Error(t, err)
	assert.Nil(t, user)
}

package users_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUsers_ListUsers(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		mux.HandleFunc(usersApiBasePath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("list_users.json"))
		})

		users, err := underTest.ListUsers()
		assert.Equal(t, 3, len(users))
		assert.NoError(t, err)
	}
}

func TestUsers_ListUsersApiFail(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		mux.HandleFunc(usersApiBasePath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
		})

		users, err := underTest.ListUsers()
		assert.Error(t, err)
		assert.Nil(t, users)
	}
}

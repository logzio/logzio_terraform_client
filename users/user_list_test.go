package users_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUsers_ListUsers(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	defer teardown()

	mux.HandleFunc("/v1/user-management", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, fixture("list_users.json"))
		w.WriteHeader(http.StatusOK)
	})

	if assert.NoError(t, err) {
		users, err := underTest.ListUsers()
		assert.Equal(t, 3, len(users))
		assert.NoError(t, err)
	}
}

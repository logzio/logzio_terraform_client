package users_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUsers_DeleteUser(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	defer teardown()

	mux.HandleFunc("/v1/user-management/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodDelete {
			fmt.Fprint(w, fixture("delete_valid_user.json"))
		}
		w.WriteHeader(http.StatusOK)
	})

	err = underTest.DeleteUser(int64(123456))
	assert.NoError(t, err)
}

func TestUsers_DeleteNotExistantUser(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	defer teardown()

	mux.HandleFunc("/v1/user-management/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodDelete {
			fmt.Fprint(w, fixture("delete_not_exist.json"))
		}
		w.WriteHeader(http.StatusOK)
	})

	err = underTest.DeleteUser(int64(123456))
	assert.Error(t, err)
}

package users_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestUsers_DeleteUser(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		id := int32(123456)
		mux.HandleFunc(usersApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.WriteHeader(http.StatusNoContent)
		})

		err = underTest.DeleteUser(id)
		assert.NoError(t, err)
	}
}

func TestUsers_DeleteUserIdNotFound(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		id := int32(123456)
		mux.HandleFunc(usersApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, fixture("delete_not_exist.json"))

		})

		err = underTest.DeleteUser(id)
		assert.Error(t, err)
	}
}

func TestUsers_DeleteUserApiFail(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		id := int32(123456)
		mux.HandleFunc(usersApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.WriteHeader(http.StatusInternalServerError)
		})

		err = underTest.DeleteUser(id)
		assert.Error(t, err)
	}
}

package users_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/users"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestUsers_GetUser(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		id := int32(123456)

		mux.HandleFunc(usersApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("get_user.json"))
		})

		user, err := underTest.GetUser(id)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, id, user.Id)
		assert.Equal(t, "test_create_user@test.co", user.UserName)
		assert.Equal(t, "test create user", user.FullName)
		assert.Equal(t, int32(1234567), user.AccountId)
		assert.Equal(t, users.UserRoleReadOnly, user.Role)
		assert.True(t, user.Active)
	}
}

func TestUsers_GetUserIdNotFound(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		id := int32(123456)

		mux.HandleFunc(usersApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.WriteHeader(http.StatusNotFound)
		})

		user, err := underTest.GetUser(id)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "failed with missing user")
	}
}

func TestUsers_GetUserApiFail(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		id := int32(123456)

		mux.HandleFunc(usersApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.WriteHeader(http.StatusInternalServerError)
		})

		user, err := underTest.GetUser(id)
		assert.Error(t, err)
		assert.Nil(t, user)
	}
}

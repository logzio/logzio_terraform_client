package users_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/users"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUsers_CreateUser(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		if assert.NoError(t, err) {
			mux.HandleFunc(usersApiBasePath, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				jsonBytes, _ := ioutil.ReadAll(r.Body)
				var target users.CreateUpdateUser
				err = json.Unmarshal(jsonBytes, &target)
				assert.NoError(t, err)
				assert.NotNil(t, target)
				assert.NotEmpty(t, target.UserName)
				assert.NotEmpty(t, target.FullName)
				assert.NotZero(t, target.AccountId)
				assert.NotEmpty(t, target.Role)
				assert.Contains(t, []string{users.UserRoleReadOnly, users.UserRoleRegular, users.UserRoleAccountAdmin}, target.Role)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, fixture("create_user.json"))
			})

			resp, err := underTest.CreateUser(createUser)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, int32(123456), resp.Id)
		}
	}
}

func TestUsers_CreateUserApiFail(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		if assert.NoError(t, err) {
			mux.HandleFunc(usersApiBasePath, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				jsonBytes, _ := ioutil.ReadAll(r.Body)
				var target users.CreateUpdateUser
				err = json.Unmarshal(jsonBytes, &target)
				assert.NoError(t, err)
				assert.NotNil(t, target)
				assert.NotEmpty(t, target.UserName)
				assert.NotEmpty(t, target.FullName)
				assert.NotZero(t, target.AccountId)
				assert.NotEmpty(t, target.Role)
				assert.Contains(t, []string{users.UserRoleReadOnly, users.UserRoleRegular, users.UserRoleAccountAdmin}, target.Role)
				w.WriteHeader(http.StatusInternalServerError)
			})

			resp, err := underTest.CreateUser(createUser)
			assert.Error(t, err)
			assert.Nil(t, resp)
		}
	}
}

func TestUsers_CreateUserDuplicateUser(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.FullName += "-dup"
		if assert.NoError(t, err) {
			mux.HandleFunc(usersApiBasePath, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				jsonBytes, _ := ioutil.ReadAll(r.Body)
				var target users.CreateUpdateUser
				err = json.Unmarshal(jsonBytes, &target)
				assert.NoError(t, err)
				assert.NotNil(t, target)
				assert.NotEmpty(t, target.UserName)
				assert.NotEmpty(t, target.FullName)
				assert.NotZero(t, target.AccountId)
				assert.NotEmpty(t, target.Role)
				assert.Contains(t, []string{users.UserRoleReadOnly, users.UserRoleRegular, users.UserRoleAccountAdmin}, target.Role)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, fixture("create_duplicate_user.json"))
			})

			resp, err := underTest.CreateUser(createUser)
			assert.Error(t, err)
			assert.Nil(t, resp)
		}
	}
}

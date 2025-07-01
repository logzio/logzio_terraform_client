package users_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/users"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestUsers_UpdateUser(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		userId := int32(123456)
		mux.HandleFunc(usersApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(userId), 10))
			jsonBytes, _ := io.ReadAll(r.Body)
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
			fmt.Fprint(w, fixture("update_user.json"))
		})

		updateUser, err := getCreateUser()
		updateUser.FullName += "-update"
		updated, err := underTest.UpdateUser(userId, updateUser)
		assert.NoError(t, err)
		assert.NotNil(t, updated)
	}
}

func TestUsers_UpdateUserApiFail(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		userId := int32(123456)
		mux.HandleFunc(usersApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(userId), 10))
			jsonBytes, _ := io.ReadAll(r.Body)
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

		updateUser, err := getCreateUser()
		updateUser.FullName += "-update"
		updated, err := underTest.UpdateUser(userId, updateUser)
		assert.Error(t, err)
		assert.Nil(t, updated)
	}
}

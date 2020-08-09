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

const (
	test_fullname = "Test User"
)

func TestUsers_CreateValidUser(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/user-management", func(w http.ResponseWriter, r *http.Request) {
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "username")
		assert.Contains(t, target, "accountID")
		assert.Contains(t, target, "fullName")
		assert.Contains(t, target, "roles")

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_user.json"))
		w.WriteHeader(http.StatusOK)
	})

	u := users.User{
		Username:  "test_create_user@test.co",
		Fullname:  "test create user",
		AccountId: int64(123456),
		Roles:     []int32{users.UserTypeUser},
	}

	user, err := underTest.CreateUser(u)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUsers_CreateDuplicateUser(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	defer teardown()

	mux.HandleFunc("/v1/user-management", func(w http.ResponseWriter, r *http.Request) {
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "username")
		assert.Contains(t, target, "accountID")
		assert.Contains(t, target, "fullName")
		assert.Contains(t, target, "roles")

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_duplicate_user.json"))
		w.WriteHeader(http.StatusOK)
	})

	u := users.User{
		Username:  "test_duplicate_user@test.co",
		Fullname:  "test duplicate user",
		AccountId: int64(123456),
		Roles:     []int32{users.UserTypeUser},
	}

	user, err := underTest.CreateUser(u)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUsers_CreateUserInvalidEmail(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	defer teardown()

	mux.HandleFunc("/v1/user-management", func(w http.ResponseWriter, r *http.Request) {
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "username")
		assert.Contains(t, target, "accountID")
		assert.Contains(t, target, "fullName")
		assert.Contains(t, target, "roles")

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_invalid_email.json"))
		w.WriteHeader(http.StatusOK)
	})

	u := users.User{
		Username:  "AnInvalidEmailAddress",
		Fullname:  "test duplicate user",
		AccountId: int64(123456),
		Roles:     []int32{users.UserTypeUser},
	}

	user, err := underTest.CreateUser(u)
	assert.Error(t, err)
	assert.Nil(t, user)
}

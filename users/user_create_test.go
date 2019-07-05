package users_test

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/jonboydell/logzio_client/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	test_username = "test@massive.co"
	test_fullname = "Test User"
)

func TestUsers_CreateValidUser(t *testing.T) {
	underTest, err := setupUsersTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		u := users.User{
			Username:  "testcreateuser@massive.co",
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
		}

		user, err := underTest.CreateUser(u)
		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			v, err := underTest.GetUser(user.Id)

			if assert.NoError(t, err) && assert.NotNil(t, v) {
				assert.Equal(t, "testcreateuser@massive.co", v.Username)
				assert.Equal(t, test_fullname, v.Fullname)
				assert.Equal(t, accountId, v.AccountId)
				assert.True(t, v.Active)
				assert.Equal(t, user.Id, user.Id)
			}

			err = underTest.DeleteUser(user.Id)
			assert.NoError(t, err)
		}
	}
}

func TestUsers_CreateDeleteDuplicateUser(t *testing.T) {
	underTest, err := setupUsersTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		u := users.User{
			Username:  "testduplicateuser@massive.co",
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
		}

		user, err := underTest.CreateUser(u)
		assert.NoError(t, err)
		_, err = underTest.CreateUser(u)
		assert.Error(t, err)

		err = underTest.DeleteUser(user.Id)
	}
}

func TestUsers_CreateInvalidUser_Email(t *testing.T) {
	underTest, err := setupUsersTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		u := users.User{
			Username:  "InvalidTestUser",
			Fullname:  "Test User",
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
		}

		_, err := underTest.CreateUser(u)
		assert.Error(t, err)
	}
}

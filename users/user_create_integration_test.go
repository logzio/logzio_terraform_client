// +build integration

package users_test

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/jonboydell/logzio_client/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationUsers_CreateValidUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	accountId, _ := test_utils.GetAccountId()

		if assert.NoError(t, err) && assert.NotZero(t, accountId) {
		u := users.User{
			Username:  "test_create_user@test.co",
			Fullname:  "test create user",
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
		}

		user, err := underTest.CreateUser(u)
		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			v, err := underTest.GetUser(user.Id)

			if assert.NoError(t, err) && assert.NotNil(t, v) {
				assert.Equal(t, "test_create_user@test.co", v.Username)
				assert.Equal(t, "test create user", v.Fullname)
				assert.Equal(t, accountId, v.AccountId)
				assert.True(t, v.Active)
				assert.Equal(t, user.Id, user.Id)
			}

			err = underTest.DeleteUser(user.Id)
			assert.NoError(t, err)
		}
	}
}

func TestIntegrationUsers_CreateDeleteDuplicateUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) && assert.NotZero(t, accountId) {
		u := users.User{
			Username:  "test_duplicate_user@test.co",
			Fullname:  "test duplicate user",
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
		}

		user, err := underTest.CreateUser(u)
		if assert.NoError(t, err) {
			_, err = underTest.CreateUser(u)
			assert.Error(t, err)
		}

		err = underTest.DeleteUser(user.Id)
	}
}

func TestIntegrationUsers_CreateInvalidUser_Email(t *testing.T) {
	underTest, err  := setupUsersIntegrationTest()
	accountId, erx := test_utils.GetAccountId()

	if assert.NoError(t, err) && assert.NoError(t, erx) && assert.NotZero(t, accountId) {
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

func TestIntegrationUsers_DeleteNonExistingUser(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	defer teardown()

	if assert.NoError(t, err) {
		err = underTest.DeleteUser(21345)
		assert.Error(t, err)
	}
}

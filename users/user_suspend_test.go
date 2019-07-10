package users_test

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/jonboydell/logzio_client/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers_SuspendUser(t *testing.T) {
	underTest, err := setupUsersTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		user, err := underTest.CreateUser(users.User{
			Username:  test_username,
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
			Active:    true,
		})

		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			suspended, err := underTest.SuspendUser(user.Id)
			assert.True(t, suspended)
			assert.NoError(t, err)
			assert.NotNil(t, user)

			u, err := underTest.GetUser(user.Id)
			assert.NoError(t, err)
			assert.False(t, u.Active)
		}

		err = underTest.DeleteUser(user.Id)
		assert.NoError(t, err)
	}
}

func TestUsers_UnsuspendUser(t *testing.T) {
	underTest, err := setupUsersTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		user, err := underTest.CreateUser(users.User{
			Username:  "testunsuspenduser@massive.co",
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
			Active:    true,
		})
		assert.NoError(t, err)

		if assert.NotNil(t, user) {
			success, err := underTest.SuspendUser(user.Id)
			assert.NoError(t, err)
			assert.True(t, success, "suspend request success should be TRUE")

			u, err := underTest.GetUser(user.Id)
			assert.NoError(t, err)
			assert.False(t, u.Active, "user should be rendred inactive, ACTIVE should be FALSE")

			success, err = underTest.UnSuspendUser(user.Id)
			assert.NoError(t, err)
			assert.True(t, success, "unsuspend request success should be TRUE")

			u, err = underTest.GetUser(user.Id)
			assert.NoError(t, err)
			assert.True(t, u.Active)
		}

		err = underTest.DeleteUser(user.Id)
		assert.NoError(t, err)
	}
}

func TestUsers_SuspendSuspendedUser(t *testing.T) {
	underTest, err := setupUsersTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		user, err := underTest.CreateUser(users.User{
			Username:  "testsuspenduser@massive.co",
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
			Active:    true,
		})
		assert.NoError(t, err)

		if assert.NotNil(t, user) {
			success, err := underTest.SuspendUser(user.Id)
			assert.NoError(t, err)
			assert.True(t, success, "suspend request success should be TRUE")

			success, err = underTest.SuspendUser(user.Id)
			assert.NoError(t, err)
			assert.True(t, success, "suspend request success should be TRUE")
		}

		err = underTest.DeleteUser(user.Id)
		assert.NoError(t, err)
	}
}

func TestUsers_UnsuspendActiveUser(t *testing.T) {
	underTest, err := setupUsersTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		user, err := underTest.CreateUser(users.User{
			Username:  "testunsuspendactiveuser@massive.co",
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
			Active:    true,
		})
		assert.NoError(t, err)

		if assert.NotNil(t, user) {
			success, err := underTest.UnSuspendUser(user.Id)
			assert.NoError(t, err)
			assert.True(t, success, "unsuspend request success should be TRUE")
		}

		err = underTest.DeleteUser(user.Id)
		assert.NoError(t, err)
	}
}

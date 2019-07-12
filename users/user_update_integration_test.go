// +build integration

package users_test

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/jonboydell/logzio_client/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationUsers_UpdateExistingUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()

	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		user, err := underTest.CreateUser(users.User{
			Username:  "updateexistinguser@massive.co",
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
		})

		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			user.Fullname = "test_updatedfullname"
			user.Active = true

			v, err := underTest.UpdateUser(*user)
			assert.NoError(t, err)

			v, err = underTest.GetUser(user.Id)

			if assert.NoError(t, err) && assert.NotNil(t, v) {
				assert.Equal(t, "test_updatedfullname", v.Fullname)
				assert.Equal(t, accountId, v.AccountId)
				assert.True(t, v.Active)
				assert.Equal(t, user.Id, user.Id)
			}

		}

		defer underTest.DeleteUser(user.Id)
	}
}

func TestIntegrationUsers_UpdateNonExistingUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		user := users.User{
			Username:  "some@random.user",
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
			Id:        -1,
		}

		_, err := underTest.UpdateUser(user)
		assert.Error(t, err)
	}
}

func TestIntegrationUsers_UpdateExistingUserInvalidUpdate(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	accountId, _ := test_utils.GetAccountId()

	if assert.NoError(t, err) {
		user, err := underTest.CreateUser(users.User{
			Username:  "updateexistinguser.invalid@massive.co",
			Fullname:  test_fullname,
			AccountId: accountId,
			Roles:     []int32{users.UserTypeUser},
		})

		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			user.Username = "test_invalidusername"
			user.Fullname = "test_updatedfullname"
			user.Active = true

			_, err := underTest.UpdateUser(*user)
			assert.Error(t, err)
		}

		defer underTest.DeleteUser(user.Id)
	}
}

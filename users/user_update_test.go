package users_test

import (
	"github.com/jonboydell/logzio_client/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers_UpdateExistingUser(t *testing.T) {
	underTest, err := setupUsersTest()

	if assert.NoError(t, err) {
		user, err := underTest.CreateUser(users.User{
			Username:  test_username,
			Fullname:  test_fullname,
			AccountId: underTest.AccountId,
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
				assert.Equal(t, underTest.AccountId, v.AccountId)
				assert.True(t, v.Active)
				assert.Equal(t, user.Id, user.Id)
			}

			err = underTest.DeleteUser(user.Id)
			assert.NoError(t, err)
		}
	}
}

func TestUsers_UpdateNonExistingUser(t *testing.T) {
	underTest, err := setupUsersTest()

	if assert.NoError(t, err) {
		user := users.User{
			Username:  "some@random.user",
			Fullname:  test_fullname,
			AccountId: underTest.AccountId,
			Roles:     []int32{users.UserTypeUser},
			Id:        -1,
		}

		_, err := underTest.UpdateUser(user)
		assert.Error(t, err)
	}
}

func TestUsers_UpdateExistingUserInvalidUpdate(t *testing.T) {
	underTest, err := setupUsersTest()

	if assert.NoError(t, err) {
		user, err := underTest.CreateUser(users.User{
			Username:  test_username,
			Fullname:  test_fullname,
			AccountId: underTest.AccountId,
			Roles:     []int32{users.UserTypeUser},
		})

		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			user.Username = "test_invalidusername"
			user.Fullname = "test_updatedfullname"
			user.Active = true

			_, err := underTest.UpdateUser(*user)
			assert.Error(t, err)

			err = underTest.DeleteUser(user.Id)
			assert.NoError(t, err)
		}
	}
}

package users

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpoints_CreateDeleteValidUser(t *testing.T) {
	setupUsersTest()

	if assert.NotNil(t, users) {
		u := User{
			Username:  "test@massive.co",
			Fullname:  "Test User",
			AccountId: accountId,
			Roles:     []int32{userTypeUser},
		}

		err, ok := validateUserRequest(u)
		assert.NoError(t, err)
		assert.True(t, ok)

		req, err := createUserApiRequest(test_utils.GetApiToken(), u)
		if assert.NoError(t, err) && assert.NotNil(t, req) {
			user, err := users.CreateUser(u)
			assert.NoError(t, err)
			if assert.NotNil(t, user) {
				v, err := users.GetUser(user.Id)

				if assert.NoError(t, err) && assert.NotNil(t, v) {
					assert.Equal(t, "test@massive.co", v.Username)
					assert.Equal(t, "Test User", v.Fullname)
					assert.Equal(t, accountId, v.AccountId)
					assert.True(t, v.Active)
					assert.Equal(t, user.Id, user.Id)
				}

				err = users.DeleteUser(user.Id)
				assert.NoError(t, err)
			}
		}
	}
}

func TestEndpoints_CreateInvalidUser_Email(t *testing.T) {
	setupUsersTest()

	if assert.NotNil(t, users) {
		u := User{
			Username:  "SomeTestUser",
			Fullname:  "Test User",
			AccountId: accountId,
			Roles:     []int32{userTypeUser},
		}

		err, ok := validateUserRequest(u)
		assert.NoError(t, err)
		assert.True(t, ok)

		req, err := createUserApiRequest(test_utils.GetApiToken(), u)
		if assert.NoError(t, err) && assert.NotNil(t, req) {
			user, err := users.CreateUser(u)
			assert.Error(t, err)
			assert.Nil(t, user)
		}
	}
}

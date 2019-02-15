package user_management

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpoints_CreateDeleteValidUser(t *testing.T) {
	setupUsersTest()

	if assert.NotNil(t, users) {
		u := User{
			Username : "SomeTestUser",
			Fullname : "Test User",
			AccountId : 00,
			Roles : []int64{2},
		}

		err, ok := validateUserRequest(u)
		assert.NoError(t, err)
		assert.True(t, ok)

		req, err := createUserApiRequest("", u)
		assert.NoError(t, err)
		assert.NotNil(t, req)

		user, err := users.CreateUser(u)
		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			err = users.DeleteUser(user.Id)
			assert.NoError(t, err)
		}
	}
}
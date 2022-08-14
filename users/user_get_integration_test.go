package users_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationUsers_GetUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.FullName += "-get"
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			if assert.NoError(t, err) && assert.NotNil(t, resp) {
				defer underTest.DeleteUser(resp.Id)
				time.Sleep(2 * time.Second)
				getUser, err := underTest.GetUser(resp.Id)
				assert.NoError(t, err)
				assert.NotNil(t, getUser)
				assert.Equal(t, resp.Id, getUser.Id)
				assert.Equal(t, createUser.FullName, getUser.FullName)
				assert.Equal(t, createUser.UserName, getUser.UserName)
				assert.Equal(t, createUser.AccountId, getUser.AccountId)
				assert.Equal(t, createUser.Role, getUser.Role)
				assert.True(t, getUser.Active)
			}
		}
	}
}

func TestIntegrationUsers_GetUserIdNotFound(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		id := int32(123456)
		user, err := underTest.GetUser(id)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "failed with missing user")
	}
}

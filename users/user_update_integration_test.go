package users_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationUsers_UpdateUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		if assert.NoError(t, err) && assert.NotNil(t, createUser) {
			createUser.FullName += "-to-update"
			resp, err := underTest.CreateUser(createUser)
			if assert.NoError(t, err) && assert.NotNil(t, resp) {
				defer underTest.DeleteUser(resp.Id)
				time.Sleep(2 * time.Second)
				createUser.FullName += "-after"
				resp, err = underTest.UpdateUser(resp.Id, createUser)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				time.Sleep(2 * time.Second)
				// Double check that the user was updated
				user, err := underTest.GetUser(resp.Id)
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, createUser.FullName, user.FullName)
			}
		}

	}
}

func TestIntegrationUsers_UpdateUserIdNotExist(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		if assert.NoError(t, err) && assert.NotNil(t, createUser) {
			createUser.FullName += "-to-update-id-not-exist"
			resp, err := underTest.UpdateUser(1234, createUser)
			assert.Error(t, err)
			assert.Nil(t, resp)
		}
	}
}

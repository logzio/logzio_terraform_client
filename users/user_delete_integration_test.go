package users_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationUsers_DeleteUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.FullName += "-delete"
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			if assert.NoError(t, err) && assert.NotNil(t, resp) {
				time.Sleep(2 * time.Second)
				defer func() {
					err = underTest.DeleteUser(resp.Id)
					assert.NoError(t, err)
				}()
			}
		}
	}
}

func TestIntegrationUsers_DeleteUserNotFound(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		err = underTest.DeleteUser(1)
		assert.Error(t, err)
	}
}

package users_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationUsers_UnsuspendUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.FullName += "-unsuspend"
		if assert.NoError(t, err) && assert.NotNil(t, createUser) {
			resp, err := underTest.CreateUser(createUser)
			if assert.NoError(t, err) && assert.NotNil(t, resp) {
				assert.NotEmpty(t, resp.Id)
				time.Sleep(2 * time.Second)
				defer underTest.DeleteUser(resp.Id)
				err = underTest.SuspendUser(resp.Id)
				assert.NoError(t, err)
				time.Sleep(2 * time.Second)
				// double check that the user was suspended
				user, err := underTest.GetUser(resp.Id)
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.False(t, user.Active)
				err = underTest.UnSuspendUser(resp.Id)
				assert.NoError(t, err)
				time.Sleep(2 * time.Second)
				// double check that the user was suspended
				user, err = underTest.GetUser(resp.Id)
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.True(t, user.Active)
			}
		}
	}
}

func TestIntegrationUsers_UnsuspendUserIdNotFound(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		id := int32(1234)
		err := underTest.UnSuspendUser(id)
		assert.Error(t, err)
	}
}

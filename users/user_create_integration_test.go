package users_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationUsers_CreateUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.FullName += "-create"
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			if assert.NoError(t, err) && assert.NotNil(t, resp) {
				defer underTest.DeleteUser(resp.Id)
				assert.NotZero(t, resp.Id)
				time.Sleep(2 * time.Second)
			}
		}
	}
}

func TestIntegrationUsers_CreateUserNoUserName(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.UserName = ""
		createUser.FullName += "-create-no-username"
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			assert.Error(t, err)
			assert.Nil(t, resp)
		}
	}
}

func TestIntegrationUsers_CreateUserNoFullName(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.FullName = ""
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			assert.Error(t, err)
			assert.Nil(t, resp)
		}
	}
}

func TestIntegrationUsers_CreateUserNoAccountId(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.AccountId = 0
		createUser.FullName += "-create-no-account-id"
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			assert.Error(t, err)
			assert.Nil(t, resp)
		}
	}
}

func TestIntegrationUsers_CreateUserNoRole(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.Role = ""
		createUser.FullName += "-create-no-role"
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			assert.Error(t, err)
			assert.Nil(t, resp)
		}
	}
}

func TestIntegrationUsers_CreateUserInvalidRole(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.Role = "SOME_ROLE"
		createUser.FullName += "-create-invalid-role"
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			assert.Error(t, err)
			assert.Nil(t, resp)
		}
	}
}

func TestIntegrationUsers_CreateUserDuplicateUser(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createUser, err := getCreateUser()
		createUser.FullName += "-create-dup"
		if assert.NoError(t, err) {
			resp, err := underTest.CreateUser(createUser)
			if assert.NoError(t, err) && assert.NotNil(t, resp) {
				defer underTest.DeleteUser(resp.Id)
				time.Sleep(2 * time.Second)
				_, err := underTest.CreateUser(createUser)
				assert.Error(t, err)
			}

		}
	}
}

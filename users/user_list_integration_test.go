package users_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationUsers_ListUsers(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		users, err := underTest.ListUsers()
		assert.NoError(t, err)
		assert.NotEmpty(t, users)
	}
}

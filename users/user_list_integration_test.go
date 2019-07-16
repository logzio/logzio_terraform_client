// +build integration

package users_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationUsers_ListUsers(t *testing.T) {
	underTest, err := setupUsersIntegrationTest()

	if assert.NoError(t, err) {
		users, err := underTest.ListUsers()
		assert.NoError(t, err)
		assert.NotEmpty(t, users, "user list shouldn't be empty")
	}
}

package users_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers_ListUsers(t *testing.T) {
	underTest, err := setupUsersTest()

	if assert.NoError(t, err) {
		users, err := underTest.ListUsers()
		assert.NoError(t, err)
		assert.NotEmpty(t, users, "user list shouldn't be empty")
	}
}

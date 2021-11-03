package authentication_groups_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationAuthenticationGroups_Get(t *testing.T) {
	underTest, err := setupAuthenticationGroupsIntegrationTest()

	if assert.NoError(t, err) {
		groups, err := underTest.GetAuthenticationGroups()
		assert.NoError(t, err)
		assert.NotNil(t, groups)
	}
}

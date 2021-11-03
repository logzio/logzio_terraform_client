package authentication_groups_test

import (
	"github.com/logzio/logzio_terraform_client/authentication_groups"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAuthenticationGroups_PostCreate(t *testing.T) {
	underTest, err := setupAuthenticationGroupsIntegrationTest()
	if assert.NoError(t, err) {
		createGroups := getCreateGroups()
		groups, err := underTest.PostAuthenticationGroups(createGroups)
		if assert.NoError(t, err) && assert.NotEmpty(t, groups) {
			time.Sleep(2 * time.Second)
			defer underTest.PostAuthenticationGroups(getEmptyGroup())

			assert.Equal(t, len(createGroups), len(groups))
			for _, group := range createGroups {
				assert.Contains(t, groups, group)
			}
		}
	}
}

func TestIntegrationAuthenticationGroups_PostDeleteGroup(t *testing.T) {
	underTest, err := setupAuthenticationGroupsIntegrationTest()
	if assert.NoError(t, err) {
		createGroups := getCreateGroups()
		groups, err := underTest.PostAuthenticationGroups(createGroups)
		if assert.NoError(t, err) && assert.NotEmpty(t, groups) {
			time.Sleep(2 * time.Second)
			defer underTest.PostAuthenticationGroups(getEmptyGroup())
			updateGroups := createGroups[0:2]
			groups, err = underTest.PostAuthenticationGroups(updateGroups)
			time.Sleep(2 * time.Second)
			if assert.NoError(t, err) && assert.NotEmpty(t, groups) {
				assert.NotEmpty(t, groups)
				assert.Equal(t, len(updateGroups), len(groups))
				for _, group := range updateGroups {
					assert.Contains(t, groups, group)
				}
			}
		}
	}
}

func TestIntegrationAuthenticationGroups_PostUpdateGroup(t *testing.T) {
	underTest, err := setupAuthenticationGroupsIntegrationTest()
	if assert.NoError(t, err) {
		createGroups := getCreateGroups()
		groups, err := underTest.PostAuthenticationGroups(createGroups)
		if assert.NoError(t, err) && assert.NotEmpty(t, groups) {
			time.Sleep(2 * time.Second)
			defer underTest.PostAuthenticationGroups(getEmptyGroup())
			createGroups[1].UserRole = authentication_groups.AuthGroupsUserRoleRegular
			groups, err = underTest.PostAuthenticationGroups(createGroups)
			time.Sleep(2 * time.Second)
			if assert.NoError(t, err) && assert.NotEmpty(t, groups) {
				assert.Equal(t, len(createGroups), len(groups))
				for _, group := range createGroups {
					assert.Contains(t, groups, group)
				}
			}
		}
	}
}

func TestIntegrationAuthenticationGroups_PostEmptyGroupName(t *testing.T) {
	underTest, err := setupAuthenticationGroupsIntegrationTest()
	if assert.NoError(t, err) {
		createGroups := getCreateGroups()
		createGroups[1].Group = ""
		groups, err := underTest.PostAuthenticationGroups(createGroups)
		assert.Error(t, err)
		assert.Nil(t, groups)
	}
}

func TestIntegrationAuthenticationGroups_PostEmptyUserRole(t *testing.T) {
	underTest, err := setupAuthenticationGroupsIntegrationTest()
	if assert.NoError(t, err) {
		createGroups := getCreateGroups()
		createGroups[1].UserRole = ""
		groups, err := underTest.PostAuthenticationGroups(createGroups)
		assert.Error(t, err)
		assert.Nil(t, groups)
	}
}

func TestIntegrationAuthenticationGroups_PostInvalidUserRole(t *testing.T) {
	underTest, err := setupAuthenticationGroupsIntegrationTest()
	if assert.NoError(t, err) {
		createGroups := getCreateGroups()
		createGroups[1].UserRole = "invalid user role"
		groups, err := underTest.PostAuthenticationGroups(createGroups)
		assert.Error(t, err)
		assert.Nil(t, groups)
	}
}

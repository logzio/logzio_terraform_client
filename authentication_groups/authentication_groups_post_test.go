package authentication_groups_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/authentication_groups"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAuthenticationGroups_Post(t *testing.T) {
	underTest, teardown, err := setupAuthenticationGroupsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(authGroupsApiBasePath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target []authentication_groups.AuthenticationGroup
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("authentication_groups.json"))
		})

		createAuthGroups := getCreateGroups()
		groups, err := underTest.PostAuthenticationGroups(createAuthGroups)
		assert.NoError(t, err)
		assert.NotNil(t, groups)
		assert.Equal(t, 3, len(groups))
		assert.Equal(t, "test_group_admin", groups[0].Group)
		assert.Equal(t, authentication_groups.AuthGroupsUserRoleAdmin, groups[0].UserRole)
		assert.Equal(t, "test_group_readonly", groups[1].Group)
		assert.Equal(t, authentication_groups.AuthGroupsUserRoleReadonly, groups[1].UserRole)
		assert.Equal(t, "test_group_regular", groups[2].Group)
		assert.Equal(t, authentication_groups.AuthGroupsUserRoleRegular, groups[2].UserRole)
	}
}

func TestAuthenticationGroups_PostEmptyGroupName(t *testing.T) {
	underTest, teardown, err := setupAuthenticationGroupsTest()
	defer teardown()

	createAuthGroups := getCreateGroups()
	createAuthGroups[1].Group = ""
	groups, err := underTest.PostAuthenticationGroups(createAuthGroups)
	assert.Error(t, err)
	assert.Nil(t, groups)
}

func TestAuthenticationGroups_PostEmptyUserRole(t *testing.T) {
	underTest, teardown, err := setupAuthenticationGroupsTest()
	defer teardown()

	createAuthGroups := getCreateGroups()
	createAuthGroups[1].UserRole = ""
	groups, err := underTest.PostAuthenticationGroups(createAuthGroups)
	assert.Error(t, err)
	assert.Nil(t, groups)
}

func TestAuthenticationGroups_PostInvalidUserRole(t *testing.T) {
	underTest, teardown, err := setupAuthenticationGroupsTest()
	defer teardown()

	createAuthGroups := getCreateGroups()
	createAuthGroups[1].UserRole = "INVALID"
	groups, err := underTest.PostAuthenticationGroups(createAuthGroups)
	assert.Error(t, err)
	assert.Nil(t, groups)
}

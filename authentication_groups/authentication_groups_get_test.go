package authentication_groups_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAuthenticationGroups_Get(t *testing.T) {
	underTest, teardown, err := setupAuthenticationGroupsTest()
	defer teardown()

	mux.HandleFunc(authGroupsApiBasePath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("authentication_groups.json"))
	})

	groups, err := underTest.GetAuthenticationGroups()
	assert.NoError(t, err)
	assert.NotNil(t, groups)
	assert.NotEmpty(t, groups)
	assert.Equal(t, 3, len(groups))
}

func TestAuthenticationGroups_GetApiFail(t *testing.T) {
	underTest, teardown, err := setupAuthenticationGroupsTest()
	defer teardown()

	mux.HandleFunc(authGroupsApiBasePath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	groups, err := underTest.GetAuthenticationGroups()
	assert.Error(t, err)
	assert.Nil(t, groups)
}

func TestAuthenticationGroups_GetNotFound(t *testing.T) {
	underTest, teardown, err := setupAuthenticationGroupsTest()
	defer teardown()

	mux.HandleFunc(authGroupsApiBasePath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusNotFound)
	})

	groups, err := underTest.GetAuthenticationGroups()
	assert.Error(t, err)
	assert.Nil(t, groups)
	assert.Contains(t, err.Error(), "failed with missing authentication groups")
}
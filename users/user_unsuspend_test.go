package users_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestUsers_UnsuspendUser(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		id := int32(123445)
		mux.HandleFunc(usersApiBasePath+"/unsuspend/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusNoContent)
		})

		err = underTest.UnSuspendUser(id)
		assert.NoError(t, err)
	}
}

func TestUsers_UnsuspendUserApiFail(t *testing.T) {
	underTest, teardown, err := setupUsersTest()
	defer teardown()
	if assert.NoError(t, err) {
		id := int32(123445)
		mux.HandleFunc(usersApiBasePath+"/unsuspend/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.WriteHeader(http.StatusInternalServerError)
		})

		err = underTest.UnSuspendUser(id)
		assert.Error(t, err)
	}
}

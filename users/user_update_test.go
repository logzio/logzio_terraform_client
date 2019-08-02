package users_test

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client/users"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestUsers_UpdateUser(t *testing.T) {
	underTest, err, teardown := setupUsersTest()
	defer teardown()

	accountId := int64(1234567)

	mux.HandleFunc("/v1/user-management/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(accountId, 10))

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "username")
		assert.Contains(t, target, "accountID")
		assert.Contains(t, target, "fullName")
		assert.Contains(t, target, "roles")

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_user.json"))
		w.WriteHeader(http.StatusOK)
	})

	u := users.User{
		Id:        accountId,
		Username:  "test_create_user@test.co",
		Fullname:  "test create user",
		AccountId: accountId,
		Roles:     []int32{users.UserTypeUser},
	}

	updated, err := underTest.UpdateUser(u)
	assert.NoError(t, err)
	assert.NotNil(t, updated)
}

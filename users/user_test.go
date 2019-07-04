package users_test

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/jonboydell/logzio_client/users"
)

var underTest *users.Users
var apiToken string
var accountId int32

func setupUsersTest() {
	apiToken = test_utils.GetApiToken()
	underTest, _ = users.New(apiToken)
	accountId = test_utils.GetAccountId()
}

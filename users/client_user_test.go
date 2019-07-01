package users

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"strconv"
)

var users *Users
var apiToken string
var accountId int32

func setupUsersTest() {
	apiToken = test_utils.GetApiToken()
	users, _ = New(apiToken)
	accId, _ := strconv.ParseInt(test_utils.GetAccountId(), 10, 32)
	accountId = int32(accId)
}

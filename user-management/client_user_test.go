package user_management

import "github.com/jonboydell/logzio_client/test_utils"

var users *Users
var apiToken string
var createdUsers []int64

func setupUsersTest() {
	apiToken = test_utils.GetApiToken()
	users, _ = New(apiToken)
	createdUsers = []int64{}
}

func tearDownUsersTest() {
	for x := 0; x < len(createdUsers); x++ {
		users.DeleteUser(createdUsers[x])
	}
}

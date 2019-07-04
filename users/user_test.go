package users_test

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/jonboydell/logzio_client/users"
)


func setupUsersTest() (*users.Users, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	accountId, err := test_utils.GetAccountId()
	if err != nil {
		return nil, err
	}

	underTest, err := users.New(apiToken, accountId)
	return underTest, nil
}

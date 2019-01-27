package user_management

import (
	"fmt"
	"github.com/jonboydell/logzio_client/client"
)

const (
	fldUserId string = "id"
	fldUserUsername            string = "username"
	fldUserFullname          string = "fullName"
	fldAccountId         string = "accountId"
	fldRoles   string = "roles"
)

const (
	userTypeUser int = 2
	userTypeAdmin int = 3
)

type User struct {
	Id int64
	Username string
	Fullname string
	AccountId int64
	Roles []int64
}

type Users struct {
	client.Client
}

func New(apiToken string) (*Users, error) {
	var c Users
	c.ApiToken = apiToken
	if len(apiToken) > 0 {
		return &c, nil
	} else {
		return nil, fmt.Errorf("API token not defined")
	}
}
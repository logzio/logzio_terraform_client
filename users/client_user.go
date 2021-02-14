package users

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	userServiceEndpoint = "%s/v1/user-management"
)

const (
	fldUserId        string = "id"
	fldUserUsername  string = "username"
	fldUserFullname  string = "fullName"
	fldUserAccountId string = "accountID"
	fldUserRoles     string = "roles"
	fldUserActive    string = "active"
)

const (
	UserTypeUser  int32 = 2
	UserTypeAdmin int32 = 3
)

type User struct {
	Id        int64
	Username  string
	Fullname  string
	AccountId int64
	Roles     []int32
	Active    bool
}

type UserError struct {
	errorCode  string
	message    string
	requestId  string
	parameters map[string]interface{}
}

type UsersClient struct {
	*client.Client
}

// Creates a new entry point into the users functions, accepts the user's logz.io API token and account Id
func New(apiToken, baseUrl string) (*UsersClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("base URL not defined")
	}
	c := &UsersClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

func jsonToUser(json map[string]interface{}) User {
	user := User{
		Id:        int64(json[fldUserId].(float64)),
		Username:  json[fldUserUsername].(string),
		Fullname:  json[fldUserFullname].(string),
		AccountId: int64(json[fldUserAccountId].(float64)),
		Active:    json[fldUserActive].(bool),
	}
	roles := json[fldUserRoles].([]interface{})
	var rs []int32
	for _, num := range roles {
		rs = append(rs, int32(num.(float64)))
	}
	user.Roles = rs
	return user
}

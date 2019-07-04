package users

import (
	"fmt"
	"github.com/jonboydell/logzio_client/client"
)

const (
	userServiceEndpoint = "%s/v1/user-management"
)

const (
	fldUserId       string = "id"
	fldUserUsername string = "username"
	fldUserFullname string = "fullName"
	fldUserAccountId    string = "accountID"
	fldUserRoles        string = "roles"
	fldUserActive string = "active"
)

const (
	UserTypeUser  int32 = 2
	UserTypeAdmin int32 = 3
)

type User struct {
	Id        int32
	Username  string
	Fullname  string
	AccountId int32
	Roles     []int32
	Active    bool
}

type UserError struct {
	errorCode string
	message string
	requestId string
	parameters map[string]interface{}
}

type Users struct {
	client.Client
}

func New(apiToken string, accountId int32) (*Users, error) {
	if len(apiToken) > 0 {
		var c Users
		c.ApiToken = apiToken
		c.AccountId = accountId
		return &c, nil
	} else {
		return nil, fmt.Errorf("API token not defined")
	}
}

func jsonToUser(json map[string]interface{}) (User) {
	user := User {
		Id: int32(json[fldUserId].(float64)),
		Username: json[fldUserUsername].(string),
		Fullname: json[fldUserFullname].(string),
		AccountId: int32(json[fldUserAccountId].(float64)),
		Active: json[fldUserActive].(bool),
	}
	roles := json[fldUserRoles].([]interface{})
	var rs []int32
	for _, num := range roles {
		rs = append(rs, int32(num.(float64)))
	}
	user.Roles = rs
	return user
}

func jsonToError(json map[string]interface{}) (UserError) {
	userError := UserError{
		errorCode: json["errorCode"].(string),
		message: json["message"].(string),
		requestId: json["requestId"].(string),
		parameters: json["parameters"].(map[string]interface{}),
	}
	return userError
}
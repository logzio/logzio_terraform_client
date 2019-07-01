package users

import (
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const (
	userServiceEndpoint = "%s/v1/user-management"
)

const (
	fldUserId       string = "id"
	fldUserUsername string = "username"
	fldUserFullname string = "fullName"
	fldAccountId    string = "accountId"
	fldRoles        string = "roles"
)

const (
	userTypeUser  int32 = 2
	userTypeAdmin int32 = 3
)

type User struct {
	Id        int32
	Username  string
	Fullname  string
	AccountId int32
	Roles     []int32
	Active    bool
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

type userValidator = func(u User) (error, bool)
type userBuilder = func(apiToken string, u User) (*http.Request, error)
type userChecker = func(b []byte) error

func (c *Users) makeUserRequest(user interface{}, validator userValidator, builder userBuilder, checker userChecker) ([]byte, error, bool) {
	u := user.(User)
	if err, ok := validator(u); !ok {
		return nil, err, false
	}
	req, _ := builder(c.ApiToken, u)

	httpClient := client.GetHttpClient(req)
	resp, _ := httpClient.Do(req)
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{200}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes), false
	}
	err = checker(jsonBytes)
	if err != nil {
		return nil, err, false
	}
	return jsonBytes, nil, true
}

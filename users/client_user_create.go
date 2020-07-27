package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	CreateUserServiceUrl     string = userServiceEndpoint
	createUserServiceMethod  string = http.MethodPost
	createUserServiceSuccess int    = http.StatusOK
)

func validateUserRequest(u User) (error, bool) {
	if len(u.Username) <= 0 {
		return fmt.Errorf("Not implemented"), false
	}
	if len(u.Fullname) <= 0 {
		return fmt.Errorf("Not implemented"), false
	}
	return nil, true
}

func (c *UsersClient) createUserApiRequest(apiToken string, u User) (*http.Request, error) {
	var (
		createUser = map[string]interface{}{
			fldUserUsername:  u.Username,
			fldUserFullname:  u.Fullname,
			fldUserAccountId: u.AccountId,
			fldUserRoles:     u.Roles,
		}
	)

	jsonBytes, err := json.Marshal(createUser)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	url := fmt.Sprintf(CreateUserServiceUrl, baseUrl)
	req, err := http.NewRequest(createUserServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func createUserHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{createUserServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	var target map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func checkCreateUserResponse(response map[string]interface{}) error {
	if _, ok := response["errorCode"]; ok {
		return fmt.Errorf("Error creating user; %v", response)
	}

	if _, ok := response["errorMessage"]; ok {
		return fmt.Errorf("Error creating user; %v", response)
	}

	if _, ok := response["id"]; ok {
		return nil
	}

	return fmt.Errorf("Error creating user; %v", response)
}

// Creates a new logz.io user, given a new User object
// Returns the new user (and nil) and (nil and) any error that occurred during the creation of the user
func (c *UsersClient) CreateUser(user User) (*User, error) {
	if err, ok := validateUserRequest(user); !ok {
		return nil, err
	}
	req, _ := c.createUserApiRequest(c.ApiToken, user)

	target, err := createUserHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = checkCreateUserResponse(target)
	if err != nil {
		return nil, err
	}

	user.Id = int64(target[fldUserId].(float64))
	return &user, nil
}

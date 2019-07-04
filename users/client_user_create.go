package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const (
	createUserServiceUrl     string = userServiceEndpoint
	createUserServiceMethod  string = http.MethodPost
	createUserServiceSuccess int    = 200
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

func createUserApiRequest(apiToken string, u User) (*http.Request, error) {
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

	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(createUserServiceUrl, baseUrl)
	req, err := http.NewRequest(createUserServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func createUserHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
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
	if ok, message := client.IsErrorResponse(response); ok {
		return fmt.Errorf("Error creating user; %s", message)
	}

	return nil
}

func (c *Users) CreateUser(user User) (*User, error) {
	if err, ok := validateUserRequest(user); !ok {
		return nil, err
	}
	req, _ := createUserApiRequest(c.ApiToken, user)

	target, err := createUserHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = checkCreateUserResponse(target)
	if err != nil {
		return nil, err
	}

	user.Id = int32(target[fldUserId].(float64))
	return &user, nil
}

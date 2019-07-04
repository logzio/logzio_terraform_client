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
	updateUserServiceUrl     string = userServiceEndpoint + "/%d"
	updateUserServiceMethod  string = http.MethodPut
	updateUserServiceSuccess int    = 200
)


func validateUserUpdateRequest(u User) (error, bool) {
	if len(u.Username) <= 0 {
		return fmt.Errorf("Not implemented"), false
	}
	if len(u.Fullname) <= 0 {
		return fmt.Errorf("Not implemented"), false
	}
	return nil, true
}

func updateUserApiRequest(apiToken string, u User) (*http.Request, error) {
	var (
		updateUser = map[string]interface{}{
			fldUserUsername:  u.Username,
			fldUserFullname:  u.Fullname,
			fldUserAccountId: u.AccountId,
			fldUserRoles:     u.Roles,
		}
	)

	jsonBytes, err := json.Marshal(updateUser)
	if err != nil {
		return nil, err
	}

	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(updateUserServiceUrl, baseUrl, u.Id)
	req, err := http.NewRequest(updateUserServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func updateUserHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, _ := httpClient.Do(req)
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

func checkUpdateUserResponse(response map[string]interface{}) error {
	if ok, message := client.IsErrorResponse(response); ok {
		return fmt.Errorf("Error updating user; %s", message)
	}

	return nil
}

func (c *Users) UpdateUser(user User) (*User, error) {
	if err, ok := validateUserUpdateRequest(user); !ok {
		return nil, err
	}
	req, _ := updateUserApiRequest(c.ApiToken, user)

	target, err := updateUserHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = checkUpdateUserResponse(target)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

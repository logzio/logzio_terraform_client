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
	updateUserServiceUrl     string = userServiceEndpoint + "/%d"
	updateUserServiceMethod  string = http.MethodPut
	updateUserServiceSuccess int    = http.StatusOK
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

func (c *UsersClient) updateUserApiRequest(apiToken string, u User) (*http.Request, error) {
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

	baseUrl := c.BaseUrl
	url := fmt.Sprintf(updateUserServiceUrl, baseUrl, u.Id)
	req, err := http.NewRequest(updateUserServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func updateUserHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{updateUserServiceSuccess}) {
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

// Updates an existing user, the supplied user object must specify their unique id
// Returns the updated user if successful, an error otherwise
// NOTE: The logz.io API user update API function will not update the username field (for an unknown reason)
func (c *UsersClient) UpdateUser(user User) (*User, error) {
	if err, ok := validateUserUpdateRequest(user); !ok {
		return nil, err
	}
	req, _ := c.updateUserApiRequest(c.ApiToken, user)

	target, err := logzio_client.CreateHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = checkUpdateUserResponse(target)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

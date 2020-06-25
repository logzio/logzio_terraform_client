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
	deleteUserServiceUrl         = userServiceEndpoint + "/%d"
	deleteUserServiceMethod      = "DELETE"
	deleteUserServiceSuccess int = 200
)

func validateDeleteUserRequest(u User) (error, bool) {
	return nil, true
}

func (c *UsersClient) deleteUserApiRequest(apiToken string, u User) (*http.Request, error) {
	var deleteUser = map[string]interface{}{
		fldUserId: u.Id,
	}

	jsonBytes, err := json.Marshal(deleteUser)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	url := fmt.Sprintf(deleteUserServiceUrl, baseUrl, u.Id)
	req, err := http.NewRequest(deleteUserServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func checkDeleteUserRequest(b []byte) error {
	return nil
}

func deleteUserHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if !logzio_client.CheckValidStatus(resp, []int{deleteUserServiceSuccess}) {
		return nil, fmt.Errorf("%d", resp.StatusCode)
	}
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	var target map[string]interface{}

	// a successful delete requests returns no body, just the 200 status code,
	// other errors can return a 200 and an error message...
	if len(jsonBytes) == 0 {
		return target, nil
	}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

func checkDeleteUserResponse(response map[string]interface{}) error {
	if _, ok := response["errorCode"]; ok {
		return fmt.Errorf("Error creating user; %v", response)
	}
	return nil
}

// Deletes a user from logz.io given their unique ID (an integer)
// Returns either nil (success) or an error if the user couldn't be deleted
func (c *UsersClient) DeleteUser(id int64) error {

	user := User{Id: id}
	if err, ok := validateDeleteUserRequest(user); !ok {
		return err
	}
	req, _ := c.deleteUserApiRequest(c.ApiToken, user)

	target, err := deleteUserHttpRequest(req)
	if err != nil {
		return err
	}

	err = checkDeleteUserResponse(target)
	if err != nil {
		return err
	}

	return nil
}

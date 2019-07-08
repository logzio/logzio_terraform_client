package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
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

func deleteUserApiRequest(apiToken string, u User) (*http.Request, error) {
	var deleteUser = map[string]interface{}{
		fldUserId: u.Id,
	}

	jsonBytes, err := json.Marshal(deleteUser)
	if err != nil {
		return nil, err
	}

	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(deleteUserServiceUrl, baseUrl, u.Id)
	req, err := http.NewRequest(deleteUserServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func checkDeleteUserRequest(b []byte) error {
	return nil
}

func deleteUserHttpRequest(req *http.Request) error {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !logzio_client.CheckValidStatus(resp, []int{deleteUserServiceSuccess}) {
		return fmt.Errorf("%d", resp.StatusCode)
	}
	return err
}

// Deletes a user from logz.io given their unique ID (an integer)
// Returns either nil (success) or an error if the user couldn't be deleted
func (c *UsersClient) DeleteUser(id int32) error {

	user := User{Id: id}
	if err, ok := validateDeleteUserRequest(user); !ok {
		return err
	}
	req, _ := deleteUserApiRequest(c.ApiToken, user)

	err := deleteUserHttpRequest(req)
	if err != nil {
		return err
	}

	return nil
}

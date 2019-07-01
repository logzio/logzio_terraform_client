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
	deleteUserServiceUrl = userServiceEndpoint + "/%d"
	deleteUserServiceMethod = "DELETE"
)

func validateDeleteUserRequest(u User) (error, bool) {
	return nil, true
}

func deleteUserApiRequest(apiToken string, u User) (*http.Request, error) {
	var deleteUser = map[string]interface{}{
		"id": u.Id,
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

func (c *Users) DeleteUser(id int32) error {

	if _, err, ok := c.makeUserRequest(User{Id:id}, validateDeleteUserRequest, deleteUserApiRequest, func(b []byte) error {
		var data map[string]interface{}
		json.Unmarshal(b, &data)

		if val, ok := data["validationErrors"]; ok {
			return fmt.Errorf("%v", val)
		}

		return nil
	}); !ok {
		return err
	} else {
		return nil
	}
	return fmt.Errorf("Not implemented")
}
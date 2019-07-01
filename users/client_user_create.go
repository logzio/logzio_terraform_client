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
	createUserServiceUrl string = userServiceEndpoint
	createUserServiceMethod string = http.MethodPost
	createUserServiceSuccess int = 200
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
		"username": u.Username,
		"fullName": u.Fullname,
		"accountID": u.AccountId,
		"roles": u.Roles,
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

func (c *Users) CreateUser(user User) (*User, error) {
	if jsonBytes, err, ok := c.makeUserRequest(user, validateUserRequest, createUserApiRequest, func(b []byte) error {
		var data map[string]interface{}
		json.Unmarshal(b, &data)
		if val, ok := data["validationErrors"]; ok {
			return fmt.Errorf("%v", val)
		}
		return nil
	}); !ok {
		return nil, err
	} else {
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		if err != nil {
			return nil, err
		}

		if _, ok := target["errorCode"]; ok {
			return nil, fmt.Errorf("Error creating user; %s", target["message"])
		}

		user.Id = int32(target["id"].(float64))

		return &user, nil
	}
}
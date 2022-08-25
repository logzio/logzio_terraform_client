package users

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listUserServiceUrl         = userServiceEndpoint
	listUserServiceMethod      = http.MethodGet
	listUserServiceSuccess int = http.StatusOK
)

func (c *UsersClient) ListUsers() ([]User, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listUserServiceMethod,
		Url:          fmt.Sprintf(listUserServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listUserServiceSuccess},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   nil,
		ApiAction:    listUserAction,
		ResourceName: userResourceName,
	})

	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(res, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

package users

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getUserServiceUrl         = userServiceEndpoint + "/%d"
	getUserServiceMethod      = http.MethodGet
	getUserServiceSuccess int = http.StatusOK
)

func (c *UsersClient) GetUser(userId int32) (*User, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getUserServiceMethod,
		Url:          fmt.Sprintf(getUserServiceUrl, c.BaseUrl, userId),
		Body:         nil,
		SuccessCodes: []int{getUserServiceSuccess},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   userId,
		ApiAction:    getUserAction,
		ResourceName: userResourceName,
	})

	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(res, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

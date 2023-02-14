package users

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createUserServiceUrl     string = userServiceEndpoint
	createUserServiceMethod  string = http.MethodPost
	createUserServiceSuccess int    = http.StatusOK
)

func (c *UsersClient) CreateUser(createUser CreateUpdateUser) (*ResponseId, error) {
	err := validateCreateUpdateUser(createUser)
	if err != nil {
		return nil, err
	}

	createUserJson, err := json.Marshal(createUser)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createUserServiceMethod,
		Url:          fmt.Sprintf(createUserServiceUrl, c.BaseUrl),
		Body:         createUserJson,
		SuccessCodes: []int{createUserServiceSuccess},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   nil,
		ApiAction:    createUserAction,
		ResourceName: userResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal ResponseId
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

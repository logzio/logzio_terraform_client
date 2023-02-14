package users

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateUserServiceUrl     string = userServiceEndpoint + "/%d"
	updateUserServiceMethod  string = http.MethodPut
	updateUserServiceSuccess int    = http.StatusOK
)

// NOTE: The logz.io API user update API function will not update the username field
func (c *UsersClient) UpdateUser(userId int32, updateUser CreateUpdateUser) (*ResponseId, error) {
	err := validateCreateUpdateUser(updateUser)
	if err != nil {
		return nil, err
	}

	updateUserJson, err := json.Marshal(updateUser)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateUserServiceMethod,
		Url:          fmt.Sprintf(updateUserServiceUrl, c.BaseUrl, userId),
		Body:         updateUserJson,
		SuccessCodes: []int{updateUserServiceSuccess},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   userId,
		ApiAction:    updateUserAction,
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

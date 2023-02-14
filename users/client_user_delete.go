package users

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteUserServiceUrl       = userServiceEndpoint + "/%d"
	deleteUserServiceMethod    = http.MethodDelete
	deleteUserSuccess          = http.StatusOK
	deleteUserSuccessNoContent = http.StatusNoContent
)

func (c *UsersClient) DeleteUser(userId int32) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteUserServiceMethod,
		Url:          fmt.Sprintf(deleteUserServiceUrl, c.BaseUrl, userId),
		Body:         nil,
		SuccessCodes: []int{deleteUserSuccess, deleteUserSuccessNoContent},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   userId,
		ApiAction:    deleteUserAction,
		ResourceName: userResourceName,
	})

	return err
}

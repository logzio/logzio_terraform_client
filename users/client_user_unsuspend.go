package users

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	unsuspendUserServiceUrl              = userServiceEndpoint + "/unsuspend/%d"
	unsuspendUserServiceMethod           = http.MethodPost
	unsuspendUserServiceSuccess          = http.StatusOK
	unsuspendUserServiceSuccessNoContent = http.StatusNoContent
)

func (c *UsersClient) UnSuspendUser(userId int32) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   unsuspendUserServiceMethod,
		Url:          fmt.Sprintf(unsuspendUserServiceUrl, c.BaseUrl, userId),
		Body:         nil,
		SuccessCodes: []int{unsuspendUserServiceSuccess, unsuspendUserServiceSuccessNoContent},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   userId,
		ApiAction:    unsuspendUserAction,
		ResourceName: userResourceName,
	})

	return err
}

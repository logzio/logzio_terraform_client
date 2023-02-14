package users

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	suspendUserServiceUrl              string = userServiceEndpoint + "/suspend/%d"
	suspendUserServiceMethod           string = http.MethodPost
	suspendUserServiceSuccess          int    = http.StatusOK
	suspendUserServiceSuccessNoContent        = http.StatusNoContent
)

func (c *UsersClient) SuspendUser(userId int32) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   suspendUserServiceMethod,
		Url:          fmt.Sprintf(suspendUserServiceUrl, c.BaseUrl, userId),
		Body:         nil,
		SuccessCodes: []int{suspendUserServiceSuccess, suspendUserServiceSuccessNoContent},
		NotFoundCode: http.StatusNotFound,
		ResourceId:   userId,
		ApiAction:    suspendUserAction,
		ResourceName: userResourceName,
	})

	return err
}

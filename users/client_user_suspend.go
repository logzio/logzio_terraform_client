package users

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"net/http"
)

const (
	suspendUserServiceUrl     string = userServiceEndpoint + "/suspend/%d"
	unsuspendUserServiceUrl   string = userServiceEndpoint + "/unsuspend/%d"
	suspendUserServiceMethod  string = http.MethodPost
	suspendUserServiceSuccess int    = 200
)

func validateSuspendUserRequest(userId int64) (error, bool) {
	return nil, true
}

func suspendUserHttpRequest(req *http.Request) error {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !logzio_client.CheckValidStatus(resp, []int{suspendUserServiceSuccess}) {
		return fmt.Errorf("%d", resp.StatusCode)
	}
	return nil
}

func (c *UsersClient) suspendUserApiRequest(apiToken string, userId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	url := fmt.Sprintf(suspendUserServiceUrl, baseUrl, userId)
	req, err := http.NewRequest(suspendUserServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *UsersClient) unsuspendUserApiRequest(apiToken string, userId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	url := fmt.Sprintf(unsuspendUserServiceUrl, baseUrl, userId)
	req, err := http.NewRequest(suspendUserServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Suspends a user (sets the ACTIVE flag to false)
// If the call was successful, true is returned (the status of the call) or an error if otherwise
func (c *UsersClient) SuspendUser(userId int64) (bool, error) {
	if err, ok := validateSuspendUserRequest(userId); !ok {
		return false, err
	}
	req, _ := c.suspendUserApiRequest(c.ApiToken, userId)

	err := suspendUserHttpRequest(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Unsuspends a user (sets the ACTIVE flag to true)
// If the call was successful, true is returned (the status of the call) or an error if otherwise
func (c *UsersClient) UnSuspendUser(userId int64) (bool, error) {
	if err, ok := validateSuspendUserRequest(userId); !ok {
		return false, err
	}
	req, _ := c.unsuspendUserApiRequest(c.ApiToken, userId)

	err := suspendUserHttpRequest(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

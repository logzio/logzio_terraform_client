package users

import (
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const (
	suspendUserServiceUrl     string = userServiceEndpoint + "/suspend/%d"
unsuspendUserServiceUrl     string = userServiceEndpoint + "/unsuspend/%d"
	suspendUserServiceMethod  string = http.MethodPost
	suspendUserServiceSuccess int    = 200
)

func validateSuspendUserRequest(userId int32) (error, bool) {
	return nil, true
}

func suspendUserHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, _ := httpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{suspendUserServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	var target map[string]interface{}
	return target, nil
}

func checkSuspendUserResponse(response map[string]interface{}) error {
	return nil
}

func suspendUserApiRequest(apiToken string, userId int32) (*http.Request, error) {
	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(suspendUserServiceUrl, baseUrl, userId)
	req, err := http.NewRequest(suspendUserServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func unsuspendUserApiRequest(apiToken string, userId int32) (*http.Request, error) {
	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(unsuspendUserServiceUrl, baseUrl, userId)
	req, err := http.NewRequest(suspendUserServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *Users) SuspendUser(userId int32) (bool, error) {
	if err, ok := validateSuspendUserRequest(userId); !ok {
		return false, err
	}
	req, _ := suspendUserApiRequest(c.ApiToken, userId)

	target, err := suspendUserHttpRequest(req)
	if err != nil {
		return false, err
	}

	err = checkSuspendUserResponse(target)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Users) UnSuspendUser(userId int32) (bool, error) {
	if err, ok := validateSuspendUserRequest(userId); !ok {
		return false, err
	}
	req, _ := unsuspendUserApiRequest(c.ApiToken, userId)

	target, err := suspendUserHttpRequest(req)
	if err != nil {
		return false, err
	}

	err = checkSuspendUserResponse(target)
	if err != nil {
		return false, err
	}

	return true, nil
}
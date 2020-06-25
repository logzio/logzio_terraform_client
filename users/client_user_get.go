package users

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	getUserServiceUrl         = userServiceEndpoint + "/%d"
	getUserServiceMethod      = http.MethodGet
	getUserServiceSuccess int = 200
)

func validateGetUserRequest(u User) (error, bool) {
	return nil, true
}

func (c *UsersClient) getUserApiRequest(apiToken string, u User) (*http.Request, error) {

	baseUrl := c.BaseUrl
	url := fmt.Sprintf(getUserServiceUrl, baseUrl, u.Id)
	req, err := http.NewRequest(getUserServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func getUserHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{getUserServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}

	var target map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

// Returns a user given their unique ID (an integer), the user ID supplied must belong to the account that the supplied
// API token belongs to, returns an error otherwise
func (c *UsersClient) GetUser(id int64) (*User, error) {

	u := User{Id: id}
	if err, ok := validateGetUserRequest(u); !ok {
		return nil, err
	}
	req, _ := c.getUserApiRequest(c.ApiToken, u)

	target, err := getUserHttpRequest(req)
	if err != nil {
		return nil, err
	}

	if errorCode, ok := target[client.ERROR_CODE]; ok {
		return nil, fmt.Errorf("%s", errorCode)
	}

	user := jsonToUser(target)

	return &user, nil
}

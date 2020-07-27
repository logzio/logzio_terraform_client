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
	listUserServiceUrl         = userServiceEndpoint
	listUserServiceMethod      = "GET"
	listUserServiceSuccess int = http.StatusOK
)

func (c *UsersClient) listUsersApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	url := fmt.Sprintf(listUserServiceUrl, baseUrl)
	req, err := http.NewRequest(listUserServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func listUsersHttpRequest(req *http.Request) ([]map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{listUserServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}

	var target []map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

// Lists all the users in an array belonging to the account that the supplied API token belongs to, returns an error otherwise
func (c *UsersClient) ListUsers() ([]User, error) {
	req, _ := c.listUsersApiRequest(c.ApiToken)

	target, err := listUsersHttpRequest(req)
	if err != nil {
		return nil, err
	}

	var users []User
	for _, json := range target {
		user := jsonToUser(json)
		users = append(users, user)
	}

	return users, nil
}

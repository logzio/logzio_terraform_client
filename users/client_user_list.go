package users

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const (
	listUserServiceUrl = userServiceEndpoint
	listUserServiceMethod = "GET"
	listUserServiceSuccess int    = 200

)

func listUsersApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(listUserServiceUrl, baseUrl)
	req, err := http.NewRequest(listUserServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func listUsersHttpRequest(req *http.Request) ([]map[string]interface{}, error){
	httpClient := client.GetHttpClient(req)
	resp, _ := httpClient.Do(req)
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{listUserServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}

	var target []map[string]interface{}
	err := json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

func (c *Users) ListUsers() ([]User, error) {
	req, _ := listUsersApiRequest(c.ApiToken)

	target, err := listUsersHttpRequest(req)
	if (err != nil) {
		return nil, err
	}

	var users []User
	for _, json := range target {
		user := jsonToUser(json)
		users = append(users, user)
	}

	return users, nil
}
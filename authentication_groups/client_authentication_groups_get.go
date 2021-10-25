package authentication_groups

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	authGroupsGetServiceUrl            = authGroupsServiceEndpoint
	authGroupsGetServiceMethod  string = http.MethodGet
	authGroupsGetMethodSuccess  int    = http.StatusOK
	authGroupsGetMethodNotFound int    = http.StatusNotFound
)

// GetAuthenticationGroups returns all authentication groups, or error if occurred
func (c *AuthenticationGroupsClient) GetAuthenticationGroups() (*AuthenticationGroups, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   authGroupsGetServiceMethod,
		Url:          fmt.Sprintf(authGroupsGetServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{authGroupsGetMethodSuccess},
		NotFoundCode: authGroupsGetMethodNotFound,
		ResourceId:   nil,
		ApiAction:    authGroupsGetOperation,
	})

	if err != nil {
		return nil, err
	}

	var archive AuthenticationGroups
	err = json.Unmarshal(res, &archive)
	if err != nil {
		return nil, err
	}

	return &archive, nil
}
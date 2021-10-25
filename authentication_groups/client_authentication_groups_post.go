package authentication_groups

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	authGroupsPostUrl           = authGroupsServiceEndpoint
	authGroupsPostServiceMethod string = http.MethodPost
	authGroupsPostStatusSuccess int    = http.StatusOK
	authGroupsPostStatusNotFound int    = http.StatusNotFound
)

// PostAuthenticationGroups receives a list of authentication groups and generates a request to Logz.io
func (c *AuthenticationGroupsClient) PostAuthenticationGroups(groups AuthenticationGroups) (*AuthenticationGroups, error) {
	groupsJson, err := json.Marshal(groups)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   authGroupsPostServiceMethod,
		Url:          fmt.Sprintf(authGroupsPostUrl, c.BaseUrl),
		Body:         groupsJson,
		SuccessCodes: []int{authGroupsPostStatusSuccess},
		NotFoundCode: authGroupsPostStatusNotFound,
		ResourceId:   nil,
		ApiAction:    authGroupsPostOperation,
	})

	if err != nil {
		return nil, err
	}

	var retVal AuthenticationGroups
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}
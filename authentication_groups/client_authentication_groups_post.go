package authentication_groups

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	authGroupsPostUrl                   = authGroupsServiceEndpoint
	authGroupsPostServiceMethod  string = http.MethodPost
	authGroupsPostStatusSuccess  int    = http.StatusOK
	authGroupsPostStatusNotFound int    = http.StatusNotFound
)

// PostAuthenticationGroups receives a list of authentication groups and generates a request to Logz.io
func (c *AuthenticationGroupsClient) PostAuthenticationGroups(groups []AuthenticationGroup) ([]AuthenticationGroup, error) {
	err := validateAuthenticationGroups(groups)
	if err != nil {
		return nil, err
	}

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

	var retVal []AuthenticationGroup
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return retVal, nil
}

func validateAuthenticationGroups(groups []AuthenticationGroup) error {
	validUserRoles := []string{
		AuthGroupsUserRoleRegular,
		AuthGroupsUserRoleReadonly,
		AuthGroupsUserRoleAdmin,
	}

	for _, group := range groups {
		if len(group.Group) == 0 {
			return fmt.Errorf("field Group must be set")
		}

		if len(group.UserRole) == 0 {
			return fmt.Errorf("field UserRole must be set")
		}

		if !logzio_client.Contains(validUserRoles, group.UserRole) {
			return fmt.Errorf("invalid user role. User role should be one of: %s", validUserRoles)
		}
	}

	return nil
}

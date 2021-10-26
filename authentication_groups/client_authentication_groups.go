package authentication_groups

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	authGroupsServiceEndpoint = "%s/v1/authentication/groups"

	authGroupsPostOperation = "PostOperation"
	authGroupsGetOperation  = "GetOperation"

	AuthGroupsUserRoleReadonly = "USER_ROLE_READONLY"
	AuthGroupsUserRoleRegular  = "USER_ROLE_REGULAR"
	AuthGroupsUserRoleAdmin    = "USER_ROLE_ACCOUNT_ADMIN"
)

type AuthenticationGroupsClient struct {
	*client.Client
}

type AuthenticationGroup struct {
	Group    string `json:"group"`
	UserRole string `json:"userRole"`
}

// New Creates a new entry point into the authentication groups functions, accepts the user's logz.io API token and base url
func New(apiToken string, baseUrl string) (*AuthenticationGroupsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	c := &AuthenticationGroupsClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

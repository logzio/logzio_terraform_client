package users

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	userServiceEndpoint  = "%s/v1/user-management"
	UserRoleReadOnly     = "USER_ROLE_READONLY"
	UserRoleRegular      = "USER_ROLE_REGULAR"
	UserRoleAccountAdmin = "USER_ROLE_ACCOUNT_ADMIN"

	createUserAction    = "CreateUpdateUser"
	deleteUserAction    = "DeleteUser"
	getUserAction       = "GetUser"
	listUserAction      = "ListUser"
	suspendUserAction   = "SuspendUser"
	unsuspendUserAction = "UnsuspendUser"
	updateUserAction    = "UpdateUser"

	userResourceName = "user"
)

type CreateUpdateUser struct {
	UserName  string `json:"username"`
	FullName  string `json:"fullName"`
	AccountId int32  `json:"accountID"`
	Role      string `json:"role"`
}

type ResponseId struct {
	Id int32 `json:"id"`
}

type User struct {
	Id        int32  `json:"id"`
	UserName  string `json:"username"`
	FullName  string `json:"fullName"`
	AccountId int32  `json:"accountID"`
	Role      string `json:"role"`
	Active    bool   `json:"active"`
}

type UsersClient struct {
	*client.Client
}

// New Creates a new entry point into the users functions, accepts the user's logz.io API token and base url
func New(apiToken string, baseUrl string) (*UsersClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	c := &UsersClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

func validateCreateUpdateUser(createUser CreateUpdateUser) error {
	if len(createUser.UserName) == 0 {
		return fmt.Errorf("UserName must be set")
	}

	if len(createUser.FullName) == 0 {
		return fmt.Errorf("FullName must be set")
	}

	if createUser.AccountId == 0 {
		return fmt.Errorf("AccountId must be set")
	}

	if len(createUser.Role) == 0 {
		return fmt.Errorf("Role must be set")
	}

	validRoles := []string{
		UserRoleReadOnly,
		UserRoleRegular,
		UserRoleAccountAdmin,
	}

	if !logzio_client.Contains(validRoles, createUser.Role) {
		return fmt.Errorf("Role must be one of %s", validRoles)
	}

	return nil
}

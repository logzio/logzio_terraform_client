# Authentication Groups

Compatible with Logz.io's [Authentication Groups API](https://docs.logz.io/api/#tag/Authentication-groups).

This is an example to a request that creates a new authentication group:

```go
client, _ := authentication_groups.New(apiToken, apiServerAddress)
groups, err := client.PostAuthenticationGroups(
	    []authentication_groups.AuthenticationGroup{
            {
            Group: "test_group_admin",
            UserRole: authentication_groups.AuthGroupsUserRoleAdmin,
            }})
```

|function|func name|
|---|---|
| post request to create, update or delete user groups | `func (c *AuthenticationGroupsClient) PostAuthenticationGroups(groups []AuthenticationGroup) ([]AuthenticationGroup, error)` |
| get authentication groups | `func (c *AuthenticationGroupsClient) GetAuthenticationGroups() ([]AuthenticationGroup, error)` |
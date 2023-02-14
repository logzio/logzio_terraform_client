# Users

Compatible with Logz.io's [manage users API](https://docs.logz.io/api/#tag/Manage-users).

To create a new user, on a specific account or sub-account. you'll need [your account Id](https://docs.logz.io/user-guide/accounts/finding-your-account-id.html).

```go
client, _ := users.New(apiToken, apiServerAddress)
createUser := users.CreateUpdateUser{
                UserName:  "some_test@test.test",
                FullName:  "user test",
                AccountId: 0,
                Role:      users.UserRoleReadOnly,
            }
resp, err := client.CreateUser(createUser)
```

| function       | func name                                                                                          |
|----------------|----------------------------------------------------------------------------------------------------|
| create user    | `func (c *UsersClient) CreateUser(createUser CreateUpdateUser) (*ResponseId, error)`               |
| update user    | `func (c *UsersClient) UpdateUser(userId int32, updateUser CreateUpdateUser) (*ResponseId, error)` |
| delete user    | `func (c *UsersClient) DeleteUser(userId int32) error`                                             |
| get user       | `func (c *UsersClient) GetUser(userId int32) (*User, error)`                                       |
| list users     | `func (c *UsersClient) ListUsers() ([]User, error)`                                                |
| suspend user   | `func (c *UsersClient) SuspendUser(userId int32) error`                                            |
| unsuspend user | `func (c *UsersClient) UnSuspendUser(userId int32) error`                                          |

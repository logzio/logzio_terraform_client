# Users

Compatible with Logz.io's [manage users API](https://docs.logz.io/api/#tag/Manage-users).

To create a new user, on a specific account or sub-account. you'll need [your account Id](https://docs.logz.io/user-guide/accounts/finding-your-account-id.html).

```go
client, _ := users.New(apiToken, apiServerAddress)
user := client.User{
    Username:  "createa@test.user",
    Fullname:  "my username",
    AccountId: 123456,
    Roles:     []int32{users.UserTypeUser},
}
```

|function|func name|
|---|---|
|create user|`func (c *UsersClient) CreateUser(user User) (*User, error)`|
|update user|`func (c *UsersClient) UpdateUser(user User) (*User, error)`|
|delete user|`func (c *UsersClient) DeleteUser(id int32) error`|
|get user|`func (c *UsersClient) GetUser(id int32) (*User, error)`|
|list users|`func (c *UsersClient) ListUsers() ([]User, error)`|
|suspend user|`func (c *UsersClient) SuspendUser(userId int32) (bool, error)`|
|unsuspend user|`func (c *UsersClient) UnSuspendUser(userId int32) (bool, error)`|

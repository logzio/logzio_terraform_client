# Logz.io client library

DEVELOP - [![Build Status](https://travis-ci.org/jonboydell/logzio_client.svg?branch=develop)](https://travis-ci.org/jonboydell/logzio_client) [![Coverage Status](https://coveralls.io/repos/github/jonboydell/logzio_client/badge.svg?branch=develop)](https://coveralls.io/github/jonboydell/logzio_client?branch=develop)

MASTER - [![Build Status](https://travis-ci.org/jonboydell/logzio_client.svg?branch=master)](https://travis-ci.org/jonboydell/logzio_client)

Client library for logz.io API, see below for supported endpoints.

The primary purpose of this library is to act as the API interface for the logz.io Terraform provider.

Logz.io have not written an especially consistent API.  Sometimes, JSON will be presented back from an API call, sometimes not.  Sometimes just a status code, sometimes a 200 status code, but with an error message in the body.  I have attempted to shield the user of this library from those inconsistencies, but as they are laregely not documented, it's 
pretty diffcult to know if I've got them all.

Some of my API methods are inconsistent.  Particularly `New` methods, where I only discovered that I needed the acount identifier as well as the API token when I was dealing with user accounts.

##### Usage



##### Alerts



|function|logz.io api doc|
|---|---|
|create alert||
|update alert||
|delete alert||
|get alert (by id)||
|get alert (by name)||
|list alerts||


##### Users

```go
// Creates a new user (a normal type user, as opposed to an admin user)
import ( 
	"github.com/jonboydell/logzio_client/users"
)

func someFunction() {
	client, _ := users.New(YOUR_API_TOKEN)
    client.CreateUser(users.User{
        Username:  "my.email@domain.com",
        Fullname:  "My Real Name",
        AccountId: accountId,
        Roles:     []int32{users.UserTypeUser},
    })
}

```

|function|method|logz.io api doc|
|---|---|---|
|create user|`users.CreateUser(user User)`|https://docs.logz.io/api/#operation/createUser|
|update user|`users.UpdateUser(user User)`|https://docs.logz.io/api/#operation/updateUser|
|delete user|`users.DeleteUser(userId int32)`|https://docs.logz.io/api/#operation/deleteUser|
|get user|`users.GetUser(userId int32)`|https://docs.logz.io/api/#operation/getUser|
|list users|`users.ListUsers()`|https://docs.logz.io/api/#operation/listUsers|
|suspend user|`users.SuspendUser(userId int32)`|https://docs.logz.io/api/#operation/suspendUser|
|unsuspend user|`users.UnSuspendUser(userId int32)`|https://docs.logz.io/api/#operation/unsuspendUser|

##### Endpoints
There's no 1-1 mapping between this library and the logz.io API functions, logz.io provide one API endpoint per *type* of notification endpoint being created.  I have abstracted this so that depending on how you create your `Endpoints` variable that you pass to `CreateEndpoint` the `CreateEndpoint` function will work out which API call to make. 

For more info, see: https://docs.logz.io/api/#tag/Manage-notification-endpoints





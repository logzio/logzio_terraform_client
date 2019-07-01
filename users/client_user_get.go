package users

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"net/http"
)

const (
	getUserServiceUrl = userServiceEndpoint + "/%d"
	getUserServiceMethod = "GET"
)

func validateGetUserRequest(u User) (error, bool) {
	return nil, true
}

func getUserApiRequest(apiToken string, u User) (*http.Request, error) {

	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(getUserServiceUrl, baseUrl, u.Id)
	req, err := http.NewRequest(getUserServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *Users) GetUser(id int32) (*User, error) {

	if jsonBytes, err, ok := c.makeUserRequest(User {Id: id}, validateGetUserRequest, getUserApiRequest, func(b []byte) error {
		var data map[string]interface{}
		json.Unmarshal(b, &data)

		if val, ok := data["validationErrors"]; ok {
			return fmt.Errorf("%v", val)
		}

		return nil
	}); !ok {
		return nil, err
	} else {
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		if err != nil {
			return nil, err
		}

		userId := int32(target["id"].(float64))
		accountId := int32(target["accountID"].(float64))
		roles := target["roles"].([]interface{})
		x := []int32{}
		for _, role := range roles {
			r := int32(role.(float64))
			x = append(x, r)
		}

		user := User {
			Id: userId,
			AccountId: accountId,
			Username: target["username"].(string),
			Fullname :target["fullName"].(string),
			Active :target["active"].(bool),
			Roles : x,
		}

		return &user, nil
	}
	return nil, fmt.Errorf("Not implemented")
}
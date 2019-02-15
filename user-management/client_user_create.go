package user_management

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func validateUserRequest(u User) (error, bool) {
	if len(u.Username) <= 0 {
		return fmt.Errorf(""), false
	}
	if len(u.Fullname) <= 0 {
		return fmt.Errorf(""), false
	}
	return nil, true
}

func createUserApiRequest(a string, u User) (*http.Request, error) {
	return nil, fmt.Errorf("")
}

func (c *Users) CreateUser(user User) (*User, error) {
	if jsonBytes, err, ok := c.makeUserRequest(user, validateUserRequest, createUserApiRequest, func(b []byte) error {
		var data map[string]interface{}
		json.Unmarshal(b, &data)

		if val, ok := data["errorCode"]; ok {
			return fmt.Errorf("%v", val)
		}

		if val, ok := data["message"]; ok {
			return fmt.Errorf("%v", val)
		}

		if strings.Contains(fmt.Sprintf("%s", b), "") {
			return fmt.Errorf("%d %s", 200, "")
		}
		return nil
	}); !ok {
		return nil, err
	} else {
		var target User
		err = json.Unmarshal(jsonBytes, &target)
		if err != nil {
			return nil, err
		}

		return &target, nil
	}
}
package sub_accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
	"strconv"
)

const (
	createSubAccountServiceUrl    string = subAccountServiceEndpoint
	createSubAccountServiceMethod string = http.MethodPost
	serviceSuccess                int    = http.StatusOK
)

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// Create a sub account, return account's id & token if successful, an error otherwise
func (c *SubAccountClient) CreateSubAccount(createSubAccount CreateOrUpdateSubAccount) (*SubAccountCreateResponse, error) {
	err := validateCreateSubAccount(createSubAccount)
	if err != nil {
		return nil, err
	}

	SubAccountJson, err := json.Marshal(createSubAccount)
	if err != nil {
		return nil, err
	}

	req, err := c.buildCreateApiRequest(c.ApiToken, SubAccountJson)
	if err != nil {
		return nil, err
	}

	jsonResponse, err := logzio_client.CreateHttpRequestBytesResponse(req)
	if err != nil {
		return nil, err
	}

	var reVal SubAccountCreateResponse
	err = json.Unmarshal(jsonResponse, &reVal)
	if err != nil {
		return nil, err
	}

	return &reVal, nil
}

func validateCreateSubAccount(createSubAccount CreateOrUpdateSubAccount) error {
	if len(createSubAccount.Email) == 0 {
		return fmt.Errorf("email must be set")
	}

	if len(createSubAccount.AccountName) == 0 {
		return fmt.Errorf("account name must be set")
	}

	if createSubAccount.SharingObjectsAccounts == nil {
		return fmt.Errorf("sharing object accounts must be initialized, even without any object")
	}

	if createSubAccount.RetentionDays <= 0 {
		return fmt.Errorf("retention days should be >= 1")
	}

	if len(createSubAccount.Flexible) > 0 {
		_, err := strconv.ParseBool(createSubAccount.Flexible)
		if err != nil {
			return fmt.Errorf("flexible field is not set to boolean value")
		}
	} else {
		if createSubAccount.ReservedDailyGB != 0 {
			return fmt.Errorf("when isFlexible=false reservedDailyGB should be 0, empty, or emitted")
		}
	}

	if len(createSubAccount.UtilizationSettings.UtilizationEnabled) > 0 {
		_, err := strconv.ParseBool(createSubAccount.UtilizationSettings.UtilizationEnabled)
		if err != nil {
			return fmt.Errorf("utilizationEnabled field is not set to boolean value")
		}
	}
	return nil
}

func (c *SubAccountClient) buildCreateApiRequest(apiToken string, jsonBytes []byte) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(createSubAccountServiceMethod, fmt.Sprintf(createSubAccountServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

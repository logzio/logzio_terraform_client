package sub_accounts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	createSubAccountServiceUrl     = subAccountServiceEndpoint
	createSubAccountServiceMethod  = http.MethodPost
	createSubAccountMethodSuccess  = http.StatusOK
	createSubAccountStatusNotFound = http.StatusNotFound
)

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// CreateSubAccount creates sub account, return account's id & token if successful, an error otherwise
func (c *SubAccountClient) CreateSubAccount(createSubAccount CreateOrUpdateSubAccount) (*SubAccountCreateResponse, error) {
	err := validateCreateSubAccount(createSubAccount)
	if err != nil {
		return nil, err
	}

	SubAccountJson, err := json.Marshal(createSubAccount)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createSubAccountServiceMethod,
		Url:          fmt.Sprintf(createSubAccountServiceUrl, c.BaseUrl),
		Body:         SubAccountJson,
		SuccessCodes: []int{createSubAccountMethodSuccess},
		NotFoundCode: createSubAccountStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateSubAccount,
		ResourceName: subAccountResourceName,
	})

	if err != nil {
		return nil, err
	}

	var reVal SubAccountCreateResponse
	err = json.Unmarshal(res, &reVal)
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
		if createSubAccount.SoftLimitGB != nil {
			return fmt.Errorf("when isFlexible=true SoftLimitGB should be empty or omitted")
		}
	} else {
		if createSubAccount.ReservedDailyGB != nil {
			return fmt.Errorf("when isFlexible=false reservedDailyGB should be 0, empty, or emitted")
		}
		if createSubAccount.SoftLimitGB != nil && *createSubAccount.SoftLimitGB <= 0 {
			return fmt.Errorf("SoftLimitGB should be > 0 when set")
		}
	}

	if len(createSubAccount.UtilizationSettings.UtilizationEnabled) > 0 {
		_, err := strconv.ParseBool(createSubAccount.UtilizationSettings.UtilizationEnabled)
		if err != nil {
			return fmt.Errorf("utilizationEnabled field is not set to boolean value")
		}
	}

	if createSubAccount.SnapSearchRetentionDays != nil {
		if *createSubAccount.SnapSearchRetentionDays < 1 {
			return fmt.Errorf("snapSearchRetentionDays should be >= 1")
		}
		if createSubAccount.RetentionDays < 4 {
			return fmt.Errorf("SnapSearchRetentionDays cannot be set if retentionDays is less than 4")
		}
	}

	return nil
}

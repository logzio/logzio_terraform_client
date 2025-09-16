package sub_accounts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	updateSubAccountServiceUrl      = subAccountServiceEndpoint + "/%d"
	updateSubAccountServiceMethod   = http.MethodPut
	updateSubAccountServiceSuccess  = http.StatusNoContent
	updateSubAccountServiceNotFound = http.StatusNotFound
)

func (c *SubAccountClient) UpdateSubAccount(subAccountId int64, updateSubAccount CreateOrUpdateSubAccount) error {
	err := validateUpdateSubAccount(updateSubAccount)
	if err != nil {
		return err
	}

	updateSubAccountJson, err := json.Marshal(updateSubAccount)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateSubAccountServiceMethod,
		Url:          fmt.Sprintf(updateSubAccountServiceUrl, c.BaseUrl, subAccountId),
		Body:         updateSubAccountJson,
		SuccessCodes: []int{updateSubAccountServiceSuccess},
		NotFoundCode: updateSubAccountServiceNotFound,
		ResourceId:   subAccountId,
		ApiAction:    operationUpdateSubAccount,
		ResourceName: subAccountResourceName,
	})

	return err
}

func validateUpdateSubAccount(updateSubAccount CreateOrUpdateSubAccount) error {
	if len(updateSubAccount.AccountName) == 0 {
		return fmt.Errorf("account name must be set")
	}

	isFlexible := mapStringToBool(updateSubAccount.Flexible)
	if isFlexible {
		_, err := strconv.ParseBool(updateSubAccount.Flexible)
		if err != nil {
			return fmt.Errorf("flexible field is not set to boolean value")
		}
		if updateSubAccount.SoftLimitGB != nil {
			return fmt.Errorf("when isFlexible=true SoftLimitGB should be empty or omitted")
		}
	} else {
		if updateSubAccount.ReservedDailyGB != nil {
			return fmt.Errorf("when isFlexible=false reservedDailyGB should be 0, empty, or emitted")
		}
		if updateSubAccount.SoftLimitGB != nil && *updateSubAccount.SoftLimitGB <= 0 {
			return fmt.Errorf("SoftLimitGB should be > 0 when set")
		}
	}

	if updateSubAccount.SnapSearchRetentionDays != nil && *updateSubAccount.SnapSearchRetentionDays < 1 {
		return fmt.Errorf("snapSearchRetentionDays should be >= 1")
	}

	return nil
}

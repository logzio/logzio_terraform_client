package sub_accounts

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
	"strconv"
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
	})

	return err
}

func validateUpdateSubAccount(updateSubAccount CreateOrUpdateSubAccount) error {
	if len(updateSubAccount.AccountName) == 0 {
		return fmt.Errorf("account name must be set")
	}

	if len(updateSubAccount.Flexible) > 0 {
		_, err := strconv.ParseBool(updateSubAccount.Flexible)
		if err != nil {
			return fmt.Errorf("flexible field is not set to boolean value")
		}
	} else {
		if updateSubAccount.ReservedDailyGB != 0 {
			return fmt.Errorf("when isFlexible=false reservedDailyGB should be 0, empty, or emitted")
		}
	}

	return nil
}

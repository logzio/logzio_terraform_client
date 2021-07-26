package sub_accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	updateSubAccountServiceUrl      string = subAccountServiceEndpoint + "/%d"
	updateSubAccountServiceMethod   string = http.MethodPut
	updateSubAccountServiceSuccess  int    = http.StatusNoContent
	updateSubAccountServiceNotFound int    = http.StatusNotFound
)

func (c *SubAccountClient) UpdateSubAccount(subAccountId int64, updateSubAccount CreateOrUpdateSubAccount) error {
	err := validateUpdateSubAccount(updateSubAccount)
	if err != nil {
		return err
	}

	req, err := c.buildUpdateApiRequest(c.ApiToken, subAccountId, updateSubAccount)
	if err != nil {
		return err
	}
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{updateSubAccountServiceSuccess}) {
		if resp.StatusCode == updateSubAccountServiceNotFound {
			return fmt.Errorf("API call %s failed with missing sub account %d, data: %s", operationUpdateSubAccount, subAccountId, jsonBytes)
		}

		return fmt.Errorf("API call %s failed with status code %d, data: %s", operationUpdateSubAccount, resp.StatusCode, jsonBytes)
	}

	return nil
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

func (c *SubAccountClient) buildUpdateApiRequest(apiToken string, subAccountId int64, updateSubAccount CreateOrUpdateSubAccount) (*http.Request, error) {
	jsonBytes, err := json.Marshal(updateSubAccount)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(updateSubAccountServiceMethod, fmt.Sprintf(updateSubAccountServiceUrl, baseUrl, subAccountId), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

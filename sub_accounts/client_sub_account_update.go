package sub_accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	updateServiceUrl     string = subAccountServiceEndpoint + "/%d"
	updateServiceMethod  string = http.MethodPut
	updateServiceSuccess int    = http.StatusNoContent
)

func (c *SubAccountClient) updateValidateRequest(id int64) (error, bool) {
	return nil, true
}

func (c *SubAccountClient) updateApiRequest(apiToken string, id int64, subAccount SubAccount) (*http.Request, error) {
	var (
		updateUser = map[string]interface{}{
			//"email":                  subAccount.Email,
			"accountName":            subAccount.AccountName,
			"maxDailyGB":             subAccount.MaxDailyGB,
			"retentionDays":          subAccount.RetentionDays,
			"searchable":             subAccount.Searchable,
			"accessible":             subAccount.Accessible,
			"sharingObjectsAccounts": subAccount.SharingObjectAccounts,
			"docSizeSetting":         subAccount.DocSizeSetting,
			"utilizationSettings":    subAccount.UtilizationSettings,
		}
	)

	jsonBytes, err := json.Marshal(updateUser)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(updateServiceUrl, c.BaseUrl, id)
	req, err := http.NewRequest(updateServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func (c *SubAccountClient) updateHttpRequest(req *http.Request) error {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !logzio_client.CheckValidStatus(resp, []int{updateServiceSuccess}) {
		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	return nil
}

func (c *SubAccountClient) UpdateSubAccount(id int64, subAccount SubAccount) error {
	if err, ok := c.getValidateRequest(id); !ok {
		return err
	}
	req, _ := c.updateApiRequest(c.ApiToken, id, subAccount)

	err := c.updateHttpRequest(req)
	if err != nil {
		return err
	}

	return nil
}

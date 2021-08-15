package sub_accounts

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	deleteSubAccountServiceUrl     string = subAccountServiceEndpoint + "/%d"
	deleteSubAccountServiceMethod  string = http.MethodDelete
	deleteSubAccountServiceSuccess int    = http.StatusNoContent
	deleteSubAccountMethodNotFound int    = http.StatusNotFound
)

// Deletes sub account specified by it's unique id, returns an error if a problem is encountered
func (c *SubAccountClient) DeleteSubAccount(subAccountId int64) error {
	req, err := c.buildDeleteApiRequest(c.ApiToken, subAccountId)
	if err != nil {
		return err
	}
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{deleteSubAccountServiceSuccess}) {
		if resp.StatusCode == deleteSubAccountMethodNotFound {
			return fmt.Errorf("API call %s failed with missing sub account %d, data: %s", operationDeleteSubAccount, subAccountId, jsonBytes)
		}

		return fmt.Errorf("API call %s failed with status code %d, data: %s", operationDeleteSubAccount, resp.StatusCode, jsonBytes)
	}

	return nil
}

func (c *SubAccountClient) buildDeleteApiRequest(apiToken string, subAccountId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteSubAccountServiceMethod, fmt.Sprintf(deleteSubAccountServiceUrl, baseUrl, subAccountId), nil)
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

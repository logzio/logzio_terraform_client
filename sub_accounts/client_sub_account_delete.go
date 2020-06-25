package sub_accounts

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	deleteServiceUrl     string = subAccountServiceEndpoint + "/%d"
	deleteServiceMethod  string = http.MethodDelete
	deleteServiceSuccess int    = http.StatusNoContent
)

func (c *SubAccountClient) deleteValidateRequest(id int64) (error, bool) {
	return nil, true
}

func (c *SubAccountClient) deleteApiRequest(apiToken string, id int64) (*http.Request, error) {
	url := fmt.Sprintf(deleteServiceUrl, c.BaseUrl, id)
	req, err := http.NewRequest(deleteServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *SubAccountClient) deleteHttpRequest(req *http.Request) error {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !logzio_client.CheckValidStatus(resp, []int{deleteServiceSuccess}) {
		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	return nil
}

func (c *SubAccountClient) DeleteSubAccount(id int64) (error) {
	if err, ok := c.deleteValidateRequest(id); !ok {
		return err
	}
	req, _ := c.deleteApiRequest(c.ApiToken, id)

	err := c.deleteHttpRequest(req)
	if err != nil {
		return err
	}

	return nil
}

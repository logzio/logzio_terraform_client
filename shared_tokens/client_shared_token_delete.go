package shared_tokens

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	deleteServiceUrl     string = sharedTokenServiceEndpoint + "/%d"
	deleteServiceMethod  string = http.MethodDelete
	deleteServiceSuccess int    = http.StatusNoContent
)

func (c *SharedTokenClient) deleteApiRequest(apiToken string, id int64) (*http.Request, error) {
	url := fmt.Sprintf(deleteServiceUrl, c.BaseUrl, id)
	req, err := http.NewRequest(deleteServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *SharedTokenClient) deleteHttpRequest(req *http.Request) error {
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

func (c *SharedTokenClient) DeleteSharedToken(id int64) error {
	req, _ := c.deleteApiRequest(c.ApiToken, id)

	err := c.deleteHttpRequest(req)
	if err != nil {
		return err
	}

	return nil
}

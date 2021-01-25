package shared_tokens

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
	updateServiceUrl     string = sharedTokenServiceEndpoint + "/%d"
	updateServiceMethod  string = http.MethodPut
	updateServiceSuccess int    = http.StatusNoContent
)

func (c *SharedTokenClient) updateApiRequest(apiToken string, id int64, sharedToken SharedToken) (*http.Request, error) {
	var (
		updateSharedToken = map[string]interface{}{
			"name":                  sharedToken.Name,
			"filters":               sharedToken.FilterIds,
		}
	)
	jsonBytes, err := json.Marshal(updateSharedToken)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(updateServiceUrl, c.BaseUrl, id)
	req, err := http.NewRequest(updateServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func (c *SharedTokenClient) updateHttpRequest(req *http.Request) error {
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

func (c *SharedTokenClient) UpdateSharedToken(id int64, sharedToken SharedToken) error {
	req, _ := c.updateApiRequest(c.ApiToken, id, sharedToken)

	err := c.updateHttpRequest(req)
	if err != nil {
		return err
	}

	return nil
}
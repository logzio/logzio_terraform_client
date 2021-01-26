package shared_tokens

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	getServiceUrl     string = sharedTokenServiceEndpoint + "/%d"
	getServiceMethod  string = http.MethodGet
	getServiceSuccess int    = http.StatusOK
)

func (c *SharedTokenClient) getApiRequest(apiToken string, id int64) (*http.Request, error) {

	url := fmt.Sprintf(getServiceUrl, c.BaseUrl, id)
	req, err := http.NewRequest(getServiceMethod, url, nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

func (c *SharedTokenClient) getHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{getServiceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	var target map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (c *SharedTokenClient) GetSharedToken(id int32) (*SharedToken, error) {
	req, _ := c.getApiRequest(c.ApiToken, int64(id))

	target, err := c.getHttpRequest(req)
	if err != nil {
		return nil, err
	}

	pretty, err := json.MarshalIndent(target, "", "	")
	if err != nil {
		c.logger.Error("Error parsing shared token: ", err)
	} else {
		c.logger.Trace("shared token:", string(pretty))
	}

	sharedToken := jsonToSharedToken(target)
	return &sharedToken, nil
}

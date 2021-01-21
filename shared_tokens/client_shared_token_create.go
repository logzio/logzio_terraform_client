package shared_tokens

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	serviceUrl     string = sharedTokenServiceEndpoint
	serviceMethod  string = http.MethodPost
	serviceSuccess int    = http.StatusOK
)

func (c *SharedTokenClient) createApiRequest(apiToken string, s SharedToken) (*http.Request, error) {
	var (
		createSharedToken = map[string]interface{}{
			"name":                  s.Name,
			"token":            	 s.Token,
			"filters":               s.FilterIds,
		}
	)

	jsonBytes, err := json.Marshal(createSharedToken)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	url := fmt.Sprintf(serviceUrl, baseUrl)
	req, err := http.NewRequest(serviceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *SharedTokenClient) CreateSharedToken(sharedToken SharedToken) (*SharedToken, error) {
	req, _ := c.createApiRequest(c.ApiToken, sharedToken)
	target, err := logzio_client.CreateHttpRequest(req)
	if err != nil {
		return nil, err
	}

	createdSharedToken := SharedToken{
		Id:                    int32(target["id"].(float64)),
		Token:            sharedToken.Token,
		Name:           sharedToken.Name,
		FilterIds:   sharedToken.FilterIds,
	}
	return &createdSharedToken, nil
}
package sub_accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const (
	serviceUrl     string = subAccountServiceEndpoint
	serviceMethod  string = http.MethodPost
	serviceSuccess int    = http.StatusOK
)

func (c *SubAccountClient) validateRequest(s SubAccount) (error, bool) {
	return nil, true
}

func (c *SubAccountClient) createApiRequest(apiToken string, s SubAccount) (*http.Request, error) {
	var (
		createUser = map[string]interface{}{
			"email":                  s.Email,
			"accountName":            s.AccountName,
			"maxDailyGB":             s.MaxDailyGB,
			"retentionDays":          s.RetentionDays,
			"searchable":             s.Searchable,
			"accessible":             s.Accessible,
			"sharingObjectsAccounts": s.SharingObjectAccounts,
			"docSizeSetting":         s.DocSizeSetting,
			"utilizationSettings":    s.UtilizationSettings,
		}
	)

	jsonBytes, err := json.Marshal(createUser)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	url := fmt.Sprintf(serviceUrl, baseUrl)
	req, err := http.NewRequest(serviceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *SubAccountClient) createHttpRequest(req *http.Request) (map[string]interface{}, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{serviceSuccess}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}
	var target map[string]interface{}
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (c *SubAccountClient) checkResponse(response map[string]interface{}) error {
	return nil
}

func (c *SubAccountClient) CreateSubAccount(subAccount SubAccount) (*SubAccount, error) {
	if err, ok := c.validateRequest(subAccount); !ok {
		return nil, err
	}
	req, _ := c.createApiRequest(c.ApiToken, subAccount)

	target, err := c.createHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = c.checkResponse(target)
	if err != nil {
		return nil, err
	}

	subAccount.Id = int32(target["accountId"].(float64))
	return &subAccount, nil
}

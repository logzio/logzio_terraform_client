package sub_accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	serviceUrl     string = subAccountServiceEndpoint
	serviceMethod  string = http.MethodPost
	serviceSuccess int    = http.StatusOK
)

func (c *SubAccountClient) createApiRequest(apiToken string, s SubAccountCreate) (*http.Request, error) {
	if s.SharingObjectAccounts == nil {
		s.SharingObjectAccounts = make([]int32, 0)
	}

	var (
		createSubAccount = map[string]interface{}{
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

	jsonBytes, err := json.Marshal(createSubAccount)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	url := fmt.Sprintf(serviceUrl, baseUrl)
	req, err := http.NewRequest(serviceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *SubAccountClient) CreateSubAccount(subAccount SubAccountCreate) (*SubAccount, error) {
	req, _ := c.createApiRequest(c.ApiToken, subAccount)
	target, err := logzio_client.CreateHttpRequest(req)
	if err != nil {
		return nil, err
	}

	var sharingAccounts []interface{}
	for _, obj := range subAccount.SharingObjectAccounts {
		sharingAccounts = append(sharingAccounts, obj)
	}

	createdSubAccount := SubAccount{
		Id:                    int64(target["accountId"].(float64)),
		MaxDailyGB:            subAccount.MaxDailyGB,
		AccountName:           subAccount.AccountName,
		UtilizationSettings:   subAccount.UtilizationSettings,
		DocSizeSetting:        subAccount.DocSizeSetting,
		Accessible:            subAccount.Accessible,
		Searchable:            subAccount.Searchable,
		RetentionDays:         subAccount.RetentionDays,
		SharingObjectAccounts: sharingAccounts,
		Token:                 target[fldAccountToken].(string),
		AccountId:             int64(target["accountId"].(float64)),
	}
	return &createdSubAccount, nil
}

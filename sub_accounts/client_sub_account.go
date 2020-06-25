package sub_accounts

import (
	"fmt"
	"github.com/jonboydell/logzio_client/client"
)

const (
	subAccountServiceEndpoint = "%s/v1/account-management/time-based-accounts"
)

const (
	fldAccountId                    string = "accountId"          //required
	fldEmail                 string = "email"       //required
	fldAccountName           string = "accountName" //required
	fldMaxDailyGB            string = "maxDailyGB"
	fldRetentionDays         string = "retentionDays" //required
	fldAccessible            string = "accessible"
	fldSearchable            string = "searchable"
	fldSharingAccountObjects string = "sharingObjectsAccounts" //required
	fldDocSizeSetting        string = "docSizeSetting"
	fldUtilizationSettings   string = "utilizationSettings"
	fldFrequencyMinutes      string = "frequencyMinutes"
	fldUtilizationEnabled    string = "utilizationEnabled"
	fldAccountToken          string = "accountToken"
	fldDailyUsagesList       string = "dailyUsagesList"
)

type SubAccount struct {
	Id                    int64
	Email                 string
	AccountName           string
	MaxDailyGB            float32
	RetentionDays         int32
	Searchable            bool
	Accessible            bool
	SharingObjectAccounts []interface{}
	DocSizeSetting        bool
	UtilizationSettings   map[string]interface{}
	AccountToken          string
	DailyUsagesList       interface{}
}

type SubAccountCreate struct {
	Email                 string
	AccountName           string
	MaxDailyGB            float32
	RetentionDays         int32
	Searchable            bool
	Accessible            bool
	SharingObjectAccounts []int32
	DocSizeSetting        bool
	UtilizationSettings   map[string]interface{}
	AccountToken          string
	DailyUsagesList       interface{}
}

type SubAccountClient struct {
	*client.Client
}

// Creates a new entry point into the sub-account functions, accepts the user's logz.io API token and account Id
func New(apiToken string, baseUrl string) (*SubAccountClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	c := &SubAccountClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

func jsonToSubAccount(json map[string]interface{}) SubAccount {
	subAccount := SubAccount{
 		Id:                    int64(json[fldAccountId].(float64)),
		Email:                 json[fldEmail].(string),
		AccountName:           json[fldAccountName].(string),
		AccountToken:          json[fldAccountToken].(string),
		MaxDailyGB:            float32(json[fldMaxDailyGB].(float64)),
		RetentionDays:         int32(json[fldRetentionDays].(float64)),
		Searchable:            json[fldSearchable].(bool),
		Accessible:            json[fldAccessible].(bool),
		DocSizeSetting:        json[fldDocSizeSetting].(bool),
		SharingObjectAccounts: json[fldSharingAccountObjects].([]interface{}),
		UtilizationSettings:   json[fldUtilizationSettings].(map[string]interface{}),
		DailyUsagesList:       json[fldDailyUsagesList],
	}
	return subAccount
}

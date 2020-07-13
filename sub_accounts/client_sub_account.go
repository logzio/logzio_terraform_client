package sub_accounts

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	subAccountServiceEndpoint = "%s/v1/account-management/time-based-accounts"
	loggerName = "logzio-client"
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
	//Email                 string
	AccountName           string
	MaxDailyGB            float32
	RetentionDays         int32
	Searchable            bool
	Accessible            bool
	SharingObjectAccounts []interface{}
	DocSizeSetting        bool
	UtilizationSettings   map[string]interface{}
	//AccountToken          string
}

type SubAccountRelation struct {
	OwnerAccountId		int64	`json:"ownerAccountId"`
	SubAccountId		int64	`json:"subAccountId"`
	Searchable			bool	`json:"searchable"`
	Accessible			bool	`json:"accessible"`
	CreatedDate			int64	`json:"createdDate"`
	LastUpdatedDate		int64	`json:"lastUpdatedDate"`
	LastUpdaterUserId	int64	`json:"lastUpdaterUserId"`
	Type				string	`json:"type"`
}

type Account struct {
	AccountId 			int64	`json:"accountId"`
	AccountToken 		string	`json:"accountToken"`
	AccountName 		string	`json:"accountName"`
	Active 				bool	`json:"active"`
	EsIndexPrefix 		string	`json:"esIndexPrefix"`
	MaxDailyGB 			int64	`json:"maxDailyGB"`
	RetentionDays 		int64	`json:"retentionDays"`
}

type SubAccountDetailed struct {
	SubAccountRelation		SubAccountRelation		`json:"subAccountRelation"`
	Account					Account					`json:"account"`
	SharingObjectAccounts 	[]interface{}			`json:"sharingObjectsAccounts"`
	UtilizationSettings  	map[string]interface{}	`json:"utilizationSettings"`
	DailyUsagesList			map[string]interface{}	`json:"dailyUsagesList"`
	DocSizeSetting        	bool					`json:"docSizeSetting"`}

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
	logger                  hclog.Logger
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
		logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Debug,
			Name:       loggerName,
			JSONFormat: true,
		}),
	}
	return c, nil
}

func jsonToSubAccount(json map[string]interface{}) SubAccount {
	subAccount := SubAccount{
		Id:                    int64(json[fldAccountId].(float64)),
		//Email:                 email.(string),
		AccountName:           json[fldAccountName].(string),
		//AccountToken:          token.(string),
		MaxDailyGB:            float32(json[fldMaxDailyGB].(float64)),
		RetentionDays:         int32(json[fldRetentionDays].(float64)),
		Searchable:            json[fldSearchable].(bool),
		Accessible:            json[fldAccessible].(bool),
		DocSizeSetting:        json[fldDocSizeSetting].(bool),
		SharingObjectAccounts: json[fldSharingAccountObjects].([]interface{}),
		UtilizationSettings:   json[fldUtilizationSettings].(map[string]interface{}),
	}

		if json[fldUtilizationSettings] != nil {
			subAccount.UtilizationSettings = json[fldUtilizationSettings].(map[string]interface{})
		}
	return subAccount
}


func jsonToDetailedSubAccount(jsonMap map[string]interface{}) (*SubAccountDetailed, error) {
	jsonBytes, err := json.Marshal(jsonMap)
	if err != nil {
		return nil, err
	}

	var subAccount SubAccountDetailed
	if err := json.Unmarshal(jsonBytes, &subAccount); err != nil {
		return nil, err
	}
	return &subAccount, nil
}
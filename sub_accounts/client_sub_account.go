package sub_accounts

import (
	"fmt"
	"github.com/jonboydell/logzio_client/client"
)

const (
	subAccountServiceEndpoint = "%s/v1/account-management/time-based-accounts"
)

const (
	fldId                    string = "id"          //required
	fldEmail                 string = "email"       //required
	fldAccountName           string = "accountName" //required
	fldMaxDailyGB            string = "maxDailyGB"
	fldRetentionDays         string = "retentionDays" //required
	fldAccessible            string = "accessible"
	fldSharingAccountObjects string = "sharingObjectsAccounts" //required
	fldDocSizeSetting        string = "docSizeSetting"
	fldUtilizationSettings   string = "utilizationSettings"
	fldFrequencyMinutes      string = "frequencyMinutes"
	fldUtilizationEnabled    string = "utilizationEnabled"
)

type SubAccount struct {
	Id                    int32
	Email                 string
	AccountName           string
	MaxDailyGB            float32
	RetentionDays         int32
	Searchable            bool
	Accessible            bool
	SharingObjectAccounts []int32
	DocSizeSetting        bool
	UtilizationSettings   map[string]string
}

type SubAccountClient struct {
	client.Client
}

// Creates a new entry point into the sub-account functions, accepts the user's logz.io API token and account Id
func New(apiToken string) (*SubAccountClient, error) {
	if len(apiToken) > 0 {
		var c SubAccountClient
		c.ApiToken = apiToken
		c.BaseUrl = client.GetLogzIoBaseUrl()
		return &c, nil
	} else {
		return nil, fmt.Errorf("API token not defined")
	}
}

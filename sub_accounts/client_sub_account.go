package sub_accounts

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	subAccountServiceEndpoint               = "%s/v1/account-management/time-based-accounts"
	loggerName                              = "logzio-client"
	operationGetSubAccount           string = "GetSubAccount"
	operationDeleteSubAccount        string = "DeleteSubAccount"
	operationGetDetailedSubAccount   string = "GetDetailedSubAccount"
	operationListSubAccounts         string = "ListSubAccounts"
	operationListDetailedSubAccounts string = "ListDetailedSubAccounts"
	operationUpdateSubAccount        string = "UpdateSubAccount"
)

type SubAccountClient struct {
	*client.Client
	logger hclog.Logger
}

type CreateOrUpdateSubAccount struct {
	Email                  string                                   `json:"email,omitempty"`
	AccountName            string                                   `json:"accountName"`
	Flexible               string                                   `json:"isFlexible,omitempty"` // boolean
	ReservedDailyGB        float32                                  `json:"reservedDailyGB,omitempty"`
	MaxDailyGB             float32                                  `json:"maxDailyGB"`
	RetentionDays          int32                                    `json:"retentionDays"`
	Searchable             string                                   `json:"searchable,omitempty"` // boolean
	Accessible             string                                   `json:"accessible,omitempty"` // boolean
	SharingObjectsAccounts []int32                                  `json:"sharingObjectsAccounts"`
	DocSizeSetting         string                                   `json:"docSizeSetting"` // boolean
	UtilizationSettings    AccountUtilizationSettingsCreateOrUpdate `json:"utilizationSettings"`
}

type AccountUtilizationSettingsCreateOrUpdate struct {
	FrequencyMinutes   int32  `json:"frequencyMinutes"`
	UtilizationEnabled string `json:"utilizationEnabled"` // boolean
}

type AccountUtilizationSettings struct {
	FrequencyMinutes   int32 `json:"frequencyMinutes"`
	UtilizationEnabled bool  `json:"utilizationEnabled"`
}

type SubAccount struct {
	AccountId              int32                      `json:"accountId"`
	Email                  string                     `json:"email"`
	AccountName            string                     `json:"accountName"`
	Flexible               bool                       `json:"isFlexible"`
	ReservedDailyGB        float32                    `json:"reservedDailyGB"`
	MaxDailyGB             float32                    `json:"maxDailyGB"`
	RetentionDays          int32                      `json:"retentionDays"`
	Searchable             bool                       `json:"searchable"`
	Accessible             bool                       `json:"accessible"`
	DocSizeSetting         bool                       `json:"docSizeSetting"`
	SharingObjectsAccounts []SharingAccount           `json:"sharingObjectsAccounts"`
	UtilizationSettings    AccountUtilizationSettings `json:"utilizationSettings"`
}

type SharingAccount struct {
	AccountId   int32  `json:"accountId"`
	AccountName string `json:"accountName"`
}

type DetailedSubAccount struct {
	SubAccountRelation     SubAccountRelationObject   `json:"subAccountRelation"`
	Account                AccountView                `json:"account"`
	SharingObjectsAccounts []AccountView              `json:"sharingObjectsAccounts"`
	UtilizationSettings    AccountUtilizationSettings `json:"utilizationSettings"`
	DailyUsagesList        DailyUsagesListObject      `json:"dailyUsagesList"`
	DocSizeSetting         bool                       `json:"docSizeSetting"`
}

type SubAccountRelationObject struct {
	OwnerAccountId    int32  `json:"ownerAccountId"`
	SubAccountId      int32  `json:"subAccountId"`
	Searchable        bool   `json:"searchable"`
	Accessible        bool   `json:"accessible"`
	CreatedDate       int64  `json:"createdDate"`
	LastUpdatedDate   int64  `json:"lastUpdatedDate"`
	LastUpdaterUserId int32  `json:"lastUpdaterUserId"`
	Type              string `json:"type"`
}

type AccountView struct {
	AccountId       int32   `json:"accountId"`
	AccountName     string  `json:"accountName"`
	AccountToken    string  `json:"accountToken"`
	Active          bool    `json:"active"`
	EsIndexPrefix   string  `json:"esIndexPrefix"`
	Flexible        bool    `json:"isFlexible"`
	ReservedDailyGB float32 `json:"reservedDailyGB"`
	MaxDailyGB      float32 `json:"maxDailyGB"`
	RetentionDays   int32   `json:"retentionDays"`
}

type DailyUsagesListObject struct {
	Usage []LHDailyCount `json:"usage"`
}

type LHDailyCount struct {
	Date  int64 `json:"date"`
	Bytes int64 `json:"bytes"`
}

type SubAccountCreateResponse struct {
	AccountId    int32  `json:"accountId"`
	AccountToken string `json:"accountToken"`
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

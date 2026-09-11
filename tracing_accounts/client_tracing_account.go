package tracing_accounts

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	tracingAccountServiceEndpoint = "%s/v1/account-management/consumption/tracing-accounts"
	loggerName                    = "logzio-client"

	operationGetTracingAccount             = "GetTracingAccount"
	operationDeleteTracingAccount          = "DeleteTracingAccount"
	operationListTracingAccounts           = "ListTracingAccounts"
	operationUpdateTracingAccount          = "UpdateTracingAccount"
	operationCreateTracingAccount          = "CreateTracingAccount"
	operationGetTracingAccountSoftLimit    = "GetTracingAccountSoftLimit"
	operationUpdateTracingAccountSoftLimit = "UpdateTracingAccountSoftLimit"

	tracingAccountResourceName = "tracing account"
)

type TracingAccountClient struct {
	*client.Client
	logger hclog.Logger
}

// CreateOrUpdateTracingAccount is the request body for creating or updating a consumption
// tracing account.
type CreateOrUpdateTracingAccount struct {
	AccountName          string  `json:"accountName"`
	AuthorizedAccountIds []int32 `json:"authorizedAccountIds"`
	MaxDailyGB           float32 `json:"maxDailyGB"`
	Email                string  `json:"email"`
}

type AuthorizedAccount struct {
	AccountId   int32  `json:"accountId"`
	AccountName string `json:"accountName"`
}

type TracingAccount struct {
	AccountId          int32               `json:"accountId"`
	AccountName        string              `json:"accountName"`
	MaxDailyGB         *float32            `json:"maxDailyGB"`
	Retention          int32               `json:"retention"`
	CreatedAt          string              `json:"createdAt"`
	Token              string              `json:"token,omitempty"`
	AuthorizedAccounts []AuthorizedAccount `json:"authorizedAccounts"`
	SuspensionState    *string             `json:"suspensionState,omitempty"`
}

// TracingAccountSoftLimit is the soft limit, in GB, of a consumption tracing account.
// SoftLimitGB is nil when no soft limit is set.
//
// For tracing the soft limit and MaxDailyGB are the same underlying value: the API reads
// SoftLimitGB from the account's maxDailyGB and writes maxDailyGB from SoftLimitGB. These endpoints
// therefore duplicate GetTracingAccount().MaxDailyGB, and are provided for API-surface completeness
// - MaxDailyGB on create or update sets the same thing.
type TracingAccountSoftLimit struct {
	AccountId   int32    `json:"accountId"`
	SoftLimitGB *float32 `json:"softLimitGB"`
}

// UpdateTracingAccountSoftLimit is the request body for setting the soft limit of a consumption
// tracing account. TracingAccountId must match the id in the path, the API rejects a mismatch.
type UpdateTracingAccountSoftLimit struct {
	TracingAccountId int32   `json:"tracingAccountId"`
	SoftLimitGB      float32 `json:"softLimitGB"`
}

// New creates a new entry point into the tracing-account functions, accepts the user's logz.io API
// token and base url
func New(apiToken string, baseUrl string) (*TracingAccountClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("base URL not defined")
	}

	c := &TracingAccountClient{
		Client: client.New(apiToken, baseUrl),
		logger: hclog.New(&hclog.LoggerOptions{
			Name:  loggerName,
			Level: hclog.LevelFromString("INFO"),
		}),
	}

	return c, nil
}

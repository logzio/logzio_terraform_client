# Sub-accounts

Compatible with Logz.io's [sub-accounts API](https://api-docs.logz.io/docs/logz/manage-time-based-log-accounts).

To create a new sub-account, on a main account.
```go
client, _ := sub_accounts.New(apiToken, apiServerAddress)
maxDailyGb := new(float32)
*maxDailyGb = 1
subaccount := sub_accounts.CreateOrUpdateSubAccount{
                Email:                  "some@email.test",
                AccountName:            "tf_client_test",
                MaxDailyGB:             maxDailyGb,
                RetentionDays:          1,
                Searchable:             strconv.FormatBool(false),
                Accessible:             strconv.FormatBool(true),
                SharingObjectsAccounts: []int32{},
                DocSizeSetting:         strconv.FormatBool(false),
            }
```

|function|func name|
|---|---|
|create sub-account|`func (c *SubAccountClient) CreateSubAccount(createSubAccount CreateOrUpdateSubAccount) (*SubAccountCreateResponse, error)`|
|update sub-account|`func (c *SubAccountClient) UpdateSubAccount(subAccountId int64, updateSubAccount CreateOrUpdateSubAccount) error`|
|delete sub-account|`func (c *SubAccountClient) DeleteSubAccount(subAccountId int64) error`|
|get sub-account|`func (c *SubAccountClient) GetSubAccount(subAccountId int64) (*SubAccount, error)`|
|get detailed sub-account|`func (c *SubAccountClient) GetDetailedSubAccount(subAccountId int64) (*DetailedSubAccount, error)`|
|list sub-accounts|`func (c *SubAccountClient) ListSubAccounts() ([]SubAccount, error)`|
|list detailed sub-accounts|`func (c *SubAccountClient) ListDetailedSubAccounts() ([]DetailedSubAccount, error)`|
|update warm tier retention|`func (c *SubAccountClient) UpdateWarmRetention(accountId int64, updateWarmRetention UpdateWarmRetention) ([]AccountRetentionDetails, error)`|

To set the warm tier retention of the main account or of one of its sub-accounts:
```go
client, _ := sub_accounts.New(apiToken, apiServerAddress)
retentionDetails, err := client.UpdateWarmRetention(accountId, sub_accounts.UpdateWarmRetention{SnapSearchRetentionDays: 3})
```
The main account must already have warm tier and keep at least 1 day of it. A sub-account can be set from 0 (no warm tier) up to the main account's warm retention. The response is the hot and warm retention of the main account and all its sub-accounts after the update.

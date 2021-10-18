# Restore logs

Compatible with Logz.io's [restore logs API](https://docs.logz.io/api/#tag/Restore-logs).

You can restore data from your active archiving account, whether an AWS S3 bucket or Azure Blob Storage. Restoring data gives you the option to query logs after they have expired from your time-based account.
To initiate a restore operation:
```go
client, err := restore_logs.New(apiToken, apiServerAddress)
restore, err := client.InitiateRestoreOperation(restore_logs.InitiateRestore{
                                                AccountName: "test_account",
                                                StartTime:   1634437185,
                                                EndTime:     1634444385,
                                                })
```

|function|func name|
|---|---|
| initiate a restore operation | `func (c *RestoreClient) InitiateRestoreOperation(initiateRestore InitiateRestore) (*RestoreOperation, error)` |
| get details of a restore operation | `func (c *RestoreClient) GetRestoreOperation(restoreId int32) (*RestoreOperation, error)` |
| list restore operations | `func (c *RestoreClient) ListRestoreOperations() ([]RestoreOperation, error)` |
| delete a restore operation | `func (c *RestoreClient) DeleteRestoreOperation(restoreId int32) (*RestoreOperation, error)` |

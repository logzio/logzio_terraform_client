# Archive logs

Compatible with Logz.io's [archive logs API](https://docs.logz.io/api/#tag/Archive-logs).
You can archive logs to an AWS S3 bucket or Azure Blob Storage. Archiving gives you the option to restore logs and query them after they have expired from your time-based account.
To create a new S3 archive:
```go
client, _ := archive_logs.New(apiToken, apiServerAddress)
secretCredentials := archive_logs.S3SecretCredentialsObject{
                        AccessKey: "some-access-key",
                        SecretKey: "some-secret-key",
                    }
storageSettings := archive_logs.S3StorageSettings{
                            CredentialsType:     archive_logs.CredentialsTypeKeys,
                            Path:                "some-path",
                            S3SecretCredentials: &secretCredentials,
                        }
var createArchive archive_logs.CreateOrUpdateArchiving
createArchive.Compressed = new(bool)
*createArchive.Compressed = true
createArchive.Enabled = new(bool)
*createArchive.Enabled = true
createArchive.StorageType = archive_logs.StorageTypeS3
archive, err := client.SetupArchive(createArchive)
```

|function|func name|
|---|---|
| setup logs archive | `func (c *ArchiveLogsClient) SetupArchive(createArchive CreateOrUpdateArchiving) (*ArchiveLogs, error)` |
| delete logs archive | `func (c *ArchiveLogsClient) DeleteArchiveLogs(archiveId int32) error` |
| list logs archives | `func (c *ArchiveLogsClient) ListArchiveLog() ([]ArchiveLogs, error)` |
| retrieve an archive | `func (c *ArchiveLogsClient) RetrieveArchiveLogsSetting(archiveId int32) (*ArchiveLogs, error)` |
| update archives settings | `func (c *ArchiveLogsClient) UpdateArchiveLogs(archiveId int32, updateArchive CreateOrUpdateArchiving) (*ArchiveLogs, error)` |

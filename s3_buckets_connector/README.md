# S3 Bucket Connector

Compatible with Logz.io's [S3 Bucket Connector API](https://docs.logz.io/api/#tag/Connect-to-S3-Buckets).

To create a new S3 bucket connector:

```go
client, _ := s3_buckets_connector.New(apiToken, test_utils.GetLogzIoBaseUrl())
addS3Buckt := false
active := true
connector, err := underTest.CreateS3BucketConnector(s3_buckets_connector.S3BucketConnectorRequest{
                        Bucket:                   "my_bucket",
                        AddS3ObjectKeyAsLogField: &addS3Buckt,
                        Active:                   &active,
                        Region:                   s3_buckets_connector.RegionUsEast1,
                        LogsType:                 s3_buckets_connector.LogsTypeElb,
                    })
```

| Function          | Function Name                                                                                                                     |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| Create connector  | `func (c *S3BucketsConnectorClient) CreateS3BucketConnector(create S3BucketConnectorRequest) (*S3BucketConnectorResponse, error)` |
| Delete connector  | `func (c *S3BucketsConnectorClient) DeleteS3BucketConnector(s3BucketConnectorId int64) error`                                     |
| Get connector     | `func (c *S3BucketsConnectorClient) GetS3BucketConnector(s3BucketConnectorId int64) (*S3BucketConnectorResponse, error)`          |
| List connectors   | `func (c *S3BucketsConnectorClient) ListS3BucketConnectors() ([]S3BucketConnectorResponse, error)`                                |
| Update connectors | `func (c *S3BucketsConnectorClient) UpdateS3BucketConnector(s3BucketConnectorId int64, update S3BucketConnectorRequest) error`    |
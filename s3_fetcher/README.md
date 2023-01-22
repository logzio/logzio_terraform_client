# S3 Fetcher

Compatible with Logz.io's [S3 Fetcher API](https://docs.logz.io/api/#tag/Connect-to-S3-Buckets).

To create a new S3 fetcher:

```go
client, _ := s3_fetcher.New(apiToken, test_utils.GetLogzIoBaseUrl())
addS3Buckt := false
active := true
connector, err := underTest.CreateS3Fetcher(s3_fetcher.S3FetcherRequest{
                        Bucket:                   "my_bucket",
                        AddS3ObjectKeyAsLogField: &addS3Buckt,
                        Active:                   &active,
                        Region:                   s3_fetcher.RegionUsEast1,
                        LogsType:                 s3_fetcher.LogsTypeElb,
                    })
```

| Function | Function Name                                                                                    |
|----------|--------------------------------------------------------------------------------------------------|
| Create   | `func (c *S3FetcherClient) CreateS3Fetcher(create S3FetcherRequest) (*S3FetcherResponse, error)` |
| Delete   | `func (c *S3FetcherClient) DeleteS3Fetcher(s3FetcherId int64) error`                             |
| Get      | `func (c *S3FetcherClient) GetS3Fetcher(s3FetcherId int64) (*S3FetcherResponse, error)`          |
| List     | `func (c *S3FetcherClient) ListS3Fetchers() ([]S3FetcherResponse, error)`                        |
| Update   | `func (c *S3FetcherClient) UpdateS3Fetcher(s3FetcherId int64, update S3FetcherRequest) error`    |
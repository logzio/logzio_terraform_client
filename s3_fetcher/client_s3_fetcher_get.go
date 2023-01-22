package s3_fetcher

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getS3FetcherServiceUrl      = s3FetcherServiceEndpoint + "/%d"
	getS3FetcherServiceMethod   = http.MethodGet
	getS3FetcherServiceSuccess  = http.StatusOK
	getS3FetcherServiceNotFound = http.StatusNotFound
)

// GetS3Fetcher returns a s3 fetcher given its unique identifier, an error otherwise
func (c *S3FetcherClient) GetS3Fetcher(s3FetcherId int64) (*S3FetcherResponse, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getS3FetcherServiceMethod,
		Url:          fmt.Sprintf(getS3FetcherServiceUrl, c.BaseUrl, s3FetcherId),
		Body:         nil,
		SuccessCodes: []int{getS3FetcherServiceSuccess},
		NotFoundCode: getS3FetcherServiceNotFound,
		ResourceId:   s3FetcherId,
		ApiAction:    operationGetS3Fetcher,
		ResourceName: s3FetcherResourceName,
	})

	if err != nil {
		return nil, err
	}

	var s3Fetcher S3FetcherResponse
	err = json.Unmarshal(res, &s3Fetcher)
	if err != nil {
		return nil, err
	}

	return &s3Fetcher, nil
}

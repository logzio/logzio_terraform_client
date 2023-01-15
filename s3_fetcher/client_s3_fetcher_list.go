package s3_fetcher

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listS3FetcherServiceUrl     = s3FetcherServiceEndpoint
	listS3FetcherServiceMethod  = http.MethodGet
	listS3FetcherServiceSuccess = http.StatusOK
	listS3FetcherStatusNotFound = http.StatusNotFound
)

// ListS3Fetchers returns all the s3 fetchers in an array, returns an error if any problem occurs during the API call
func (c *S3FetcherClient) ListS3Fetchers() ([]S3FetcherResponse, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listS3FetcherServiceMethod,
		Url:          fmt.Sprintf(listS3FetcherServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listS3FetcherServiceSuccess},
		NotFoundCode: listS3FetcherStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationListS3Fetcher,
		ResourceName: s3FetcherResourceName,
	})

	if err != nil {
		return nil, err
	}

	var s3Fetchers []S3FetcherResponse
	err = json.Unmarshal(res, &s3Fetchers)
	if err != nil {
		return nil, err
	}

	return s3Fetchers, nil
}

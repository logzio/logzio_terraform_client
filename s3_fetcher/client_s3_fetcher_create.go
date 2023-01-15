package s3_fetcher

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createS3FetcherServiceUrl     = s3FetcherServiceEndpoint
	createS3FetcherServiceMethod  = http.MethodPost
	createS3FetcherMethodCreated  = http.StatusCreated
	createS3FetcherStatusNotFound = http.StatusNotFound
)

// CreateS3Fetcher creates a new S3 fetcher if successful, an error otherwise
func (c *S3FetcherClient) CreateS3Fetcher(create S3FetcherRequest) (*S3FetcherResponse, error) {
	err := validateCreateUpdateS3FetcherRequest(create)
	if err != nil {
		return nil, err
	}

	createS3FetcherJson, err := json.Marshal(create)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createS3FetcherServiceMethod,
		Url:          fmt.Sprintf(createS3FetcherServiceUrl, c.BaseUrl),
		Body:         createS3FetcherJson,
		SuccessCodes: []int{createS3FetcherMethodCreated},
		NotFoundCode: createS3FetcherStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateS3Fetcher,
		ResourceName: s3FetcherResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal S3FetcherResponse
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

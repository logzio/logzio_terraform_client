package s3_buckets_connector

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createS3BucketsConnectorServiceUrl     = s3BucketsServiceEndpoint
	createS3BucketsConnectorServiceMethod  = http.MethodPost
	createS3BucketsConnectorMethodCreated  = http.StatusCreated
	createS3BucketsConnectorStatusNotFound = http.StatusNotFound
)

// CreateS3BucketConnector creates a new S3 bucket connector if successful, an error otherwise
func (c *S3BucketsConnectorClient) CreateS3BucketConnector(create S3BucketConnectorRequest) (*S3BucketConnectorResponse, error) {
	err := validateCreateUpdateS3BucketRequest(create)
	if err != nil {
		return nil, err
	}

	createS3ConnectorJson, err := json.Marshal(create)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createS3BucketsConnectorServiceMethod,
		Url:          fmt.Sprintf(createS3BucketsConnectorServiceUrl, c.BaseUrl),
		Body:         createS3ConnectorJson,
		SuccessCodes: []int{createS3BucketsConnectorMethodCreated},
		NotFoundCode: createS3BucketsConnectorStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationCreateS3BucketConnector,
		ResourceName: s3BucketConnectorResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal S3BucketConnectorResponse
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

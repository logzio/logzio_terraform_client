package s3_buckets_connector

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listS3BucketConnectorServiceUrl     = s3BucketsServiceEndpoint
	listS3BucketConnectorServiceMethod  = http.MethodGet
	listS3BucketConnectorServiceSuccess = http.StatusOK
	listS3BucketConnectorStatusNotFound = http.StatusNotFound
)

// ListS3BucketConnectors returns all the s3 bucket connectors in an array, returns an error if any problem occurs during the API call
func (c *S3BucketsConnectorClient) ListS3BucketConnectors() ([]S3BucketConnector, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listS3BucketConnectorServiceMethod,
		Url:          fmt.Sprintf(listS3BucketConnectorServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listS3BucketConnectorServiceSuccess},
		NotFoundCode: listS3BucketConnectorStatusNotFound,
		ResourceId:   nil,
		ApiAction:    operationListS3BucketConnector,
		ResourceName: s3BucketConnectorResourceName,
	})

	if err != nil {
		return nil, err
	}

	var s3BucketConnectors []S3BucketConnector
	err = json.Unmarshal(res, &s3BucketConnectors)
	if err != nil {
		return nil, err
	}

	return s3BucketConnectors, nil
}

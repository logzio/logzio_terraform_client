package s3_buckets_connector

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getS3BucketConnectorServiceUrl      = s3BucketsServiceEndpoint + "/%d"
	getS3BucketConnectorServiceMethod   = http.MethodGet
	getS3BucketConnectorServiceSuccess  = http.StatusOK
	getS3BucketConnectorServiceNotFound = http.StatusNotFound
)

// GetS3BucketConnector returns a sub account given its unique identifier, an error otherwise
func (c *S3BucketsConnectorClient) GetS3BucketConnector(s3BucketConnectorId int64) (*S3BucketConnector, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getS3BucketConnectorServiceMethod,
		Url:          fmt.Sprintf(getS3BucketConnectorServiceUrl, c.BaseUrl, s3BucketConnectorId),
		Body:         nil,
		SuccessCodes: []int{getS3BucketConnectorServiceSuccess},
		NotFoundCode: getS3BucketConnectorServiceNotFound,
		ResourceId:   s3BucketConnectorId,
		ApiAction:    operationGetS3BucketConnector,
		ResourceName: s3BucketConnectorResourceName,
	})

	if err != nil {
		return nil, err
	}

	var s3BucketConnector S3BucketConnector
	err = json.Unmarshal(res, &s3BucketConnector)
	if err != nil {
		return nil, err
	}

	return &s3BucketConnector, nil
}

package s3_buckets_connector

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteS3BucketConnectorServiceUrl     = s3BucketsServiceEndpoint + "/%d"
	deleteS3BucketConnectorServiceMethod  = http.MethodDelete
	deleteS3BucketConnectorServiceSuccess = http.StatusOK
	deleteS3BucketConnectorNotFound       = http.StatusNotFound
)

// DeleteS3BucketConnector deletes a s3 bucket connector specified by its unique id, returns an error if a problem is encountered
func (c *S3BucketsConnectorClient) DeleteS3BucketConnector(s3BucketConnectorId int64) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteS3BucketConnectorServiceMethod,
		Url:          fmt.Sprintf(deleteS3BucketConnectorServiceUrl, c.BaseUrl, s3BucketConnectorId),
		Body:         nil,
		SuccessCodes: []int{deleteS3BucketConnectorServiceSuccess},
		NotFoundCode: deleteS3BucketConnectorNotFound,
		ResourceId:   s3BucketConnectorId,
		ApiAction:    operationDeleteS3BucketConnector,
		ResourceName: s3BucketConnectorResourceName,
	})

	return err
}

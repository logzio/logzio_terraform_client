package s3_buckets_connector

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateS3BucketConnectorServiceUrl      = s3BucketsServiceEndpoint + "/%d"
	updateS3BucketConnectorServiceMethod   = http.MethodPut
	updateS3BucketConnectorServiceSuccess  = http.StatusOK
	updateS3BucketConnectorServiceNotFound = http.StatusNotFound
)

func (c *S3BucketsConnectorClient) UpdateSubAccount(s3BucketConnectorId int64, update S3BucketConnectorRequest) error {
	err := validateCreateUpdateS3BucketRequest(update)
	if err != nil {
		return err
	}

	updateS3BucketConnectorJson, err := json.Marshal(update)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateS3BucketConnectorServiceMethod,
		Url:          fmt.Sprintf(updateS3BucketConnectorServiceUrl, c.BaseUrl, s3BucketConnectorId),
		Body:         updateS3BucketConnectorJson,
		SuccessCodes: []int{updateS3BucketConnectorServiceSuccess},
		NotFoundCode: updateS3BucketConnectorServiceNotFound,
		ResourceId:   s3BucketConnectorId,
		ApiAction:    operationUpdateS3BucketConnector,
		ResourceName: s3BucketConnectorResourceName,
	})

	return err
}

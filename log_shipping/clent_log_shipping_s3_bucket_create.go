package log_shipping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const createS3BucketServiceMethod string = http.MethodPost
const createS3BucketMethodSuccess int = http.StatusOK

// type FieldError struct {
// 	Field   string
// 	Message string
// }

// func (e FieldError) Error() string {
// 	return fmt.Sprintf("%v: %v", e.Field, e.Message)
// }

func validateCreateS3BucketRequest(bucket S3Bucket) error {

	if len(bucket.AccessKey) == 0 {
		return fmt.Errorf("accessKey must be set")
	}

	// TODO: complete validation

	return nil
}

func buildCreateS3BucketRequest(bucket S3Bucket) map[string]interface{} {
	var createBucket = map[string]interface{}{}

	createBucket[fldAwsAccessKey] = bucket.AccessKey
	createBucket[fldAwsSecretKey] = bucket.SecretKey
	createBucket[fldAwsArn] = bucket.Arn
	createBucket[fldAwsBucket] = bucket.Bucket
	createBucket[fldAwsPrefix] = bucket.Prefix
	createBucket[fldAwsActive] = bucket.Active
	createBucket[fldAwsAddS3ObjectKeyAsLogField] = bucket.AddS3ObjectKeyAsLogField
	createBucket[fldAwsRegion] = bucket.Region
	createBucket[fldLogsType] = bucket.LogsType

	return createBucket
}

func (c *LogShippingClient) buildCreateApiRequest(apiToken string, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(createS3BucketServiceMethod, fmt.Sprintf(createS3BucketServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Create an alert, return the created alert if successful, an error otherwise
func (c *LogShippingClient) CreateAlert(bucket S3Bucket) (*S3Bucket, error) {
	err := validateCreateS3BucketRequest(bucket)
	if err != nil {
		return nil, err
	}

	createAlert := buildCreateS3BucketRequest(bucket)
	req, _ := c.buildCreateApiRequest(c.ApiToken, createAlert)
	jsonResponse, err := logzio_client.CreateHttpRequest(req)
	if err != nil {
		return nil, err
	}

	retVal := jsonS3BucketToS3Bucket(jsonResponse)

	return &retVal, nil
}

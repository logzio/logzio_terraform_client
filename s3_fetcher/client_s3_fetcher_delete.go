package s3_fetcher

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteS3FetcherServiceUrl     = s3FetcherServiceEndpoint + "/%d"
	deleteS3FetcherServiceMethod  = http.MethodDelete
	deleteS3FetcherServiceSuccess = http.StatusOK
	deleteS3FetcherNotFound       = http.StatusNotFound
)

// DeleteS3Fetcher deletes a s3 fetcher specified by its unique id, returns an error if a problem is encountered
func (c *S3FetcherClient) DeleteS3Fetcher(s3FetcherId int64) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteS3FetcherServiceMethod,
		Url:          fmt.Sprintf(deleteS3FetcherServiceUrl, c.BaseUrl, s3FetcherId),
		Body:         nil,
		SuccessCodes: []int{deleteS3FetcherServiceSuccess},
		NotFoundCode: deleteS3FetcherNotFound,
		ResourceId:   s3FetcherId,
		ApiAction:    operationDeleteS3Fetcher,
		ResourceName: s3FetcherResourceName,
	})

	return err
}

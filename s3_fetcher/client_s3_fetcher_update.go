package s3_fetcher

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateS3FetcherServiceUrl      = s3FetcherServiceEndpoint + "/%d"
	updateS3FetcherServiceMethod   = http.MethodPut
	updateS3FetcherServiceSuccess  = http.StatusOK
	updateS3FetcherServiceNotFound = http.StatusNotFound
)

func (c *S3FetcherClient) UpdateS3Fetcher(s3FetcherId int64, update S3FetcherRequest) error {
	err := validateCreateUpdateS3FetcherRequest(update)
	if err != nil {
		return err
	}

	updateS3FetcherJson, err := json.Marshal(update)
	if err != nil {
		return err
	}

	_, err = logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateS3FetcherServiceMethod,
		Url:          fmt.Sprintf(updateS3FetcherServiceUrl, c.BaseUrl, s3FetcherId),
		Body:         updateS3FetcherJson,
		SuccessCodes: []int{updateS3FetcherServiceSuccess},
		NotFoundCode: updateS3FetcherServiceNotFound,
		ResourceId:   s3FetcherId,
		ApiAction:    operationUpdateS3Fetcher,
		ResourceName: s3FetcherResourceName,
	})

	return err
}

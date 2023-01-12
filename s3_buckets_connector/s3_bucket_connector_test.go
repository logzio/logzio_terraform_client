package s3_buckets_connector_test

import (
	"github.com/logzio/logzio_terraform_client/s3_buckets_connector"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

const (
	keys = "keys"
	arn  = "arn"

	envArn       = "AWS_ARN_S3_CONNECTOR"
	envAccessKey = "AWS_ACCESS_KEY"
	envSecretKey = "AWS_SECRET_KEY"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupS3BucketConnectorTest() (*s3_buckets_connector.S3BucketsConnectorClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := s3_buckets_connector.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupS3BucketConnectorIntegrationTest() (*s3_buckets_connector.S3BucketsConnectorClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := s3_buckets_connector.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func getCreateOrUpdateS3BucketConnector(authType string, isLocalTest bool) s3_buckets_connector.S3BucketConnectorRequest {
	addS3ObjectKeyAsLogField := true
	active := false
	request := s3_buckets_connector.S3BucketConnectorRequest{
		Bucket:                   "miri-test-elb-tf", // TODO - change to a different bucket
		AddS3ObjectKeyAsLogField: &addS3ObjectKeyAsLogField,
		Active:                   &active,
		Region:                   s3_buckets_connector.RegionUsEast1,
		LogsType:                 s3_buckets_connector.LogsTypeElb,
	}

	if authType == keys {
		if isLocalTest {
			request.AccessKey = "my_access_key"
			request.SecretKey = "my_secret_key"
		} else {
			request.AccessKey = os.Getenv(envAccessKey)
			request.SecretKey = os.Getenv(envSecretKey)
		}
	} else {
		if isLocalTest {
			request.Arn = "my_arn"
		} else {
			request.Arn = os.Getenv(envArn)
		}
	}

	return request
}

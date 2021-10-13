package archive_logs_test

import (
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setupArchiveLogsTest() (*archive_logs.ArchiveLogsClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := archive_logs.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupArchiveLogsIntegrationTest() (*archive_logs.ArchiveLogsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := archive_logs.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func getCreateOrUpdateArchiveLogs(storageType string) (archive_logs.CreateOrUpdateArchiving, error) {
	var createOrUpdate archive_logs.CreateOrUpdateArchiving
	createOrUpdate.Compressed = new(bool)
	*createOrUpdate.Compressed = true
	createOrUpdate.Enabled = new(bool)
	*createOrUpdate.Enabled = true
	createOrUpdate.StorageType = storageType
	switch storageType {
	case archive_logs.StorageTypeS3:
		// default is KEYS credentials
		storageSettings, err := getS3StorageSettings()
		if err != nil {
			return createOrUpdate, err
		}

		createOrUpdate.AmazonS3StorageSettings = storageSettings
	case archive_logs.StorageTypeBlob:
		storageSettings, err := getBlobStorageSettings()
		if err != nil {
			return createOrUpdate, err
		}

		createOrUpdate.AzureBlobStorageSettings = storageSettings
	}

	return createOrUpdate, nil
}

func getS3StorageSettings() (*archive_logs.S3StorageSettings, error) {
	s3Path, err := test_utils.GetS3Path()
	if err != nil {
		return nil, err
	}

	accessKey, err := test_utils.GetAwsAccessKey()
	if err != nil {
		return nil, err
	}

	secretKey, err := test_utils.GetAwsSecretKey()
	if err != nil {
		return nil, err
	}

	secretCredentials := archive_logs.S3SecretCredentialsObject{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	storageSettings := archive_logs.S3StorageSettings{
		CredentialsType:     archive_logs.CredentialsTypeKeys,
		Path:                s3Path,
		S3SecretCredentials: &secretCredentials,
	}

	return &storageSettings, nil
}

func getS3IamCredentials() (*archive_logs.S3IamCredentials, error) {
	arn, err := test_utils.GetAwsIamCredentials()
	if err != nil {
		return nil, err
	}

	iamCredentials := archive_logs.S3IamCredentials{Arn: arn}
	return &iamCredentials, nil
}

func getBlobStorageSettings() (*archive_logs.BlobSettings, error) {
	tenantId, err := test_utils.GetAzureTenantId()
	if err != nil {
		return nil, err
	}

	clientId, err := test_utils.GetAzureClientId()
	if err != nil {
		return nil, err
	}

	clientSecret, err := test_utils.GetAzureClientSecret()
	if err != nil {
		return nil, err
	}

	accountName, err := test_utils.GetAzureAccountName()
	if err != nil {
		return nil, err
	}

	containerName, err := test_utils.GetAzureContainerName()
	if err != nil {
		return nil, err
	}

	storageSettings := archive_logs.BlobSettings{
		TenantId:      tenantId,
		ClientId:      clientId,
		ClientSecret:  clientSecret,
		AccountName:   accountName,
		ContainerName: containerName,
	}

	return &storageSettings, nil
}
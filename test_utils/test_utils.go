package test_utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	EnvLogzioBaseUrl             = "LOGZIO_BASE_URL"
	EnvLogzioApiToken     string = "LOGZIO_API_TOKEN"
	EnvLogzioWarmApiToken string = "LOGZIO_WARM_API_TOKEN"
	EnvLogzioAccountId    string = "LOGZIO_ACCOUNT_ID"
	EnvMetricsAccountId   string = "LOGZIO_METRICS_ACCOUNT_ID"
	LogzioBaseUrl         string = "https://api.logz.io"
	EnvLogzioEmail        string = "LOGZIO_EMAIL"
	EnvS3Path             string = "S3_PATH"
	EnvAwsAccessKey       string = "AWS_ACCESS_KEY"
	EnvAwsSecretKey       string = "AWS_SECRET_KEY"
	EnvAwsArn             string = "AWS_ARN"
	EnvAzureTenantId      string = "AZURE_TENANT_ID"
	EnvAzureClientId      string = "AZURE_CLIENT_ID"
	EnvAzureClientSecret  string = "AZURE_CLIENT_SECRET"
	EnvAzureAccountName   string = "AZURE_ACCOUNT_NAME"
	EnvAzureContainerName string = "AZURE_CONTAINER_NAME"
	EnvMetricsFolderId    string = "METRICS_FOLDER_ID"

	TestTimeSeparator = time.Millisecond * 500
)

func GetApiToken() (string, error) {
	apiToken := os.Getenv(EnvLogzioApiToken)
	if len(apiToken) > 0 {
		return apiToken, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvLogzioApiToken)
}

func GetWarmApiToken() (string, error) {
	apiToken := os.Getenv(EnvLogzioWarmApiToken)
	if len(apiToken) > 0 {
		return apiToken, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvLogzioWarmApiToken)
}

func GetAccountId() (int64, error) {
	account_id_string := os.Getenv(EnvLogzioAccountId)
	if len(account_id_string) == 0 {
		return -1, fmt.Errorf("%s env var not specified", EnvLogzioAccountId)
	}
	account_id, _ := strconv.ParseInt(account_id_string, 10, 32)
	return int64(account_id), nil
}

func GetLogzIoBaseUrl() string {
	if len(os.Getenv(EnvLogzioBaseUrl)) > 0 {
		return os.Getenv(EnvLogzioBaseUrl)
	}
	return LogzioBaseUrl
}

func GetLogzioEmail() (string, error) {
	email := os.Getenv(EnvLogzioEmail)
	if len(email) > 0 {
		return email, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvLogzioEmail)
}

func GetS3Path() (string, error) {
	s3Path := os.Getenv(EnvS3Path)
	if len(s3Path) > 0 {
		return s3Path, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvS3Path)
}

func GetAwsAccessKey() (string, error) {
	accessKey := os.Getenv(EnvAwsAccessKey)
	if len(accessKey) > 0 {
		return accessKey, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvAwsAccessKey)
}

func GetAwsSecretKey() (string, error) {
	secretKey := os.Getenv(EnvAwsSecretKey)
	if len(secretKey) > 0 {
		return secretKey, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvAwsSecretKey)
}

func GetAwsIamCredentials() (string, error) {
	arn := os.Getenv(EnvAwsArn)
	if len(arn) > 0 {
		return arn, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvAwsArn)
}

func GetAzureTenantId() (string, error) {
	tenantId := os.Getenv(EnvAzureTenantId)
	if len(tenantId) > 0 {
		return tenantId, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvAzureTenantId)
}

func GetAzureClientId() (string, error) {
	tenantId := os.Getenv(EnvAzureClientId)
	if len(tenantId) > 0 {
		return tenantId, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvAzureClientId)
}

func GetAzureClientSecret() (string, error) {
	tenantId := os.Getenv(EnvAzureClientSecret)
	if len(tenantId) > 0 {
		return tenantId, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvAzureClientSecret)
}

func GetAzureAccountName() (string, error) {
	tenantId := os.Getenv(EnvAzureAccountName)
	if len(tenantId) > 0 {
		return tenantId, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvAzureAccountName)
}

func GetAzureContainerName() (string, error) {
	tenantId := os.Getenv(EnvAzureContainerName)
	if len(tenantId) > 0 {
		return tenantId, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvAzureContainerName)
}

func TestDoneTimeBuffer() {
	time.Sleep(TestTimeSeparator)
}

func GetMetricsFolderId() (string, error) {
	folderId := os.Getenv(EnvMetricsFolderId)
	if len(folderId) > 0 {
		return folderId, nil
	}
	return "", fmt.Errorf("%s env var not specified", EnvMetricsFolderId)
}

func GetMetricsAccountId() (int64, error) {
	account_id_string := os.Getenv(EnvMetricsAccountId)
	if len(account_id_string) == 0 {
		return -1, fmt.Errorf("%s env var not specified", EnvMetricsAccountId)
	}
	account_id, _ := strconv.ParseInt(account_id_string, 10, 32)
	return int64(account_id), nil
}

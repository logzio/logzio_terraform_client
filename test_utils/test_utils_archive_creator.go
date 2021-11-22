package test_utils

import "github.com/logzio/logzio_terraform_client/archive_logs"

func GetCreateOrUpdateArchiveLogs(storageType string) (archive_logs.CreateOrUpdateArchiving, error) {
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
	s3Path, err := GetS3Path()
	if err != nil {
		return nil, err
	}

	accessKey, err := GetAwsAccessKey()
	if err != nil {
		return nil, err
	}

	secretKey, err := GetAwsSecretKey()
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

func GetS3IamCredentials() (*archive_logs.S3IamCredentials, error) {
	arn, err := GetAwsIamCredentials()
	if err != nil {
		return nil, err
	}

	iamCredentials := archive_logs.S3IamCredentials{Arn: arn}
	return &iamCredentials, nil
}

func getBlobStorageSettings() (*archive_logs.BlobSettings, error) {
	tenantId, err := GetAzureTenantId()
	if err != nil {
		return nil, err
	}

	clientId, err := GetAzureClientId()
	if err != nil {
		return nil, err
	}

	clientSecret, err := GetAzureClientSecret()
	if err != nil {
		return nil, err
	}

	accountName, err := GetAzureAccountName()
	if err != nil {
		return nil, err
	}

	containerName, err := GetAzureContainerName()
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

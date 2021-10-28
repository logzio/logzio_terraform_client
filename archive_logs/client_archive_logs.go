package archive_logs

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const archiveLogsServiceEndpoint string = "%s/v2/archive/settings"

const (
	StorageTypeS3       = "S3"
	StorageTypeBlob     = "BLOB"
	CredentialsTypeIam  = "IAM"
	CredentialsTypeKeys = "KEYS"

	setupArchiveApi         = "SetupArchiveLogs"
	retrieveArchiveSettings = "RetrieveArchiveSettings"
	updateArchiveSettings   = "UpdateArchiveSettings"
	deleteArchiveSettings   = "DeleteArchiveSettings"
	testArchiveSettings     = "TestArchiveSettings"
	listArchivesSettings    = "ListArchiveSettings"

	archiveResourceName = "archive"
)

type ArchiveLogsClient struct {
	*client.Client
}

type CreateOrUpdateArchiving struct {
	StorageType              string             `json:"storageType"`                        // required
	Enabled                  *bool              `json:"enabled,omitempty"`                  // boolean - defined as a pointer because omitempty automatically omits false value
	Compressed               *bool              `json:"compressed,omitempty"`               // boolean - defined as a pointer because omitempty automatically omits false value
	AmazonS3StorageSettings  *S3StorageSettings `json:"amazonS3StorageSettings,omitempty"`  // Set as pointer so that on marshalling will omit empty
	AzureBlobStorageSettings *BlobSettings      `json:"azureBlobStorageSettings,omitempty"` // Set as pointer so that on marshalling will omit empty
}

// S3StorageSettings - use when StorageType is S3
type S3StorageSettings struct {
	CredentialsType     string                     `json:"credentialsType,omitempty"`     // required
	Path                string                     `json:"path,omitempty"`                // required
	S3SecretCredentials *S3SecretCredentialsObject `json:"s3SecretCredentials,omitempty"` // Set as pointer so that on marshalling will omit empty
	S3IamCredentials    *S3IamCredentials          `json:"s3IamCredentials,omitempty"`    // Set as pointer so that on marshalling will omit empty
}

// S3SecretCredentialsObject - use when CredentialsType is KEYS
type S3SecretCredentialsObject struct {
	AccessKey string `json:"accessKey,omitempty"` // required
	SecretKey string `json:"secretKey,omitempty"` // required
}

// S3IamCredentials - use when CredentialsType is IAM
type S3IamCredentials struct {
	Arn        string `json:"arn,omitempty"`        // required
	ExternalId string `json:"externalId,omitempty"` // only in response
}

// BlobSettings - use when StorageType is BLOB
type BlobSettings struct {
	TenantId      string `json:"tenantId,omitempty"`      // required
	ClientId      string `json:"clientId,omitempty"`      // required
	ClientSecret  string `json:"clientSecret,omitempty"`  // required
	AccountName   string `json:"accountName,omitempty"`   // required
	ContainerName string `json:"containerName,omitempty"` // required
	Path          string `json:"path,omitempty"`
}

type ArchiveLogs struct {
	Id       int32           `json:"id"`
	Settings StorageSettings `json:"settings"`
}

type StorageSettings struct {
	StorageType              string            `json:"storageType"` // required
	Enabled                  bool              `json:"enabled"`
	Compressed               bool              `json:"compressed"`
	AmazonS3StorageSettings  S3StorageSettings `json:"amazonS3StorageSettings"`
	AzureBlobStorageSettings BlobSettings      `json:"azureBlobStorageSettings"`
}

// New Creates a new entry point into the archive logs functions, accepts the user's logz.io API token and base url
func New(apiToken string, baseUrl string) (*ArchiveLogsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	c := &ArchiveLogsClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

func validateCreateOrUpdateArchiveRequest(createOrUpdateArchive CreateOrUpdateArchiving) error {
	switch createOrUpdateArchive.StorageType {
	case StorageTypeS3:
		return validateS3Settings(createOrUpdateArchive.AmazonS3StorageSettings)
	case StorageTypeBlob:
		return validateBlobSettings(createOrUpdateArchive.AzureBlobStorageSettings)
	default:
		return fmt.Errorf("storage type must be one of %s", []string{StorageTypeS3, StorageTypeBlob})
	}
}

func validateS3Settings(s3Settings *S3StorageSettings) error {
	if len(s3Settings.Path) == 0 {
		return fmt.Errorf("s3 storage path must be set for storage type %s", StorageTypeS3)
	}

	switch s3Settings.CredentialsType {
	case CredentialsTypeIam:
		if len(s3Settings.S3IamCredentials.Arn) == 0 {
			return fmt.Errorf("arn must be set for credentials type %s", CredentialsTypeIam)
		}
	case CredentialsTypeKeys:
		if len(s3Settings.S3SecretCredentials.AccessKey) == 0 {
			return fmt.Errorf("access key must be set for credentials type %s", CredentialsTypeKeys)
		}
		if len(s3Settings.S3SecretCredentials.SecretKey) == 0 {
			return fmt.Errorf("secret key must be set for credentials type %s", CredentialsTypeKeys)
		}
	default:
		return fmt.Errorf("credentials type must be one of %s", []string{CredentialsTypeIam, CredentialsTypeKeys})
	}

	return nil
}

func validateBlobSettings(blobSettings *BlobSettings) error {
	if len(blobSettings.TenantId) == 0 {
		return fmt.Errorf("tenant id must be set for storage type %s", StorageTypeBlob)
	}

	if len(blobSettings.ClientId) == 0 {
		return fmt.Errorf("client id must be set for storage type %s", StorageTypeBlob)
	}

	if len(blobSettings.ClientSecret) == 0 {
		return fmt.Errorf("client secret must be set for storage type %s", StorageTypeBlob)
	}

	if len(blobSettings.AccountName) == 0 {
		return fmt.Errorf("account name must be set for storage type %s", StorageTypeBlob)
	}

	if len(blobSettings.ContainerName) == 0 {
		return fmt.Errorf("container name must be set for storage type %s", StorageTypeBlob)
	}

	return nil
}

package s3_fetcher

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	s3FetcherServiceEndpoint string = "%s/v1/log-shipping/s3-buckets"

	s3FetcherResourceName = "s3 fetcher"

	operationCreateS3Fetcher = "CreateS3Fetcher"
	operationGetS3Fetcher    = "GetS3Fetcher"
	operationListS3Fetcher   = "ListS3Fetcher"
	operationUpdateS3Fetcher = "UpdateS3Fetcher"
	operationDeleteS3Fetcher = "DeleteS3Fetcher"
)

type S3FetcherClient struct {
	*client.Client
}

type S3FetcherRequest struct {
	AccessKey                string      `json:"accessKey,omitempty"`
	SecretKey                string      `json:"secretKey,omitempty"`
	Arn                      string      `json:"arn,omitempty"`
	Bucket                   string      `json:"bucket"`
	Prefix                   string      `json:"prefix,omitempty"`
	Active                   *bool       `json:"active"`
	AddS3ObjectKeyAsLogField *bool       `json:"addS3ObjectKeyAsLogField,omitempty"`
	Region                   AwsRegion   `json:"region"`
	LogsType                 AwsLogsType `json:"logsType"`
}

type S3FetcherResponse struct {
	AccessKey                string      `json:"accessKey,omitempty"`
	Arn                      string      `json:"arn,omitempty"`
	Bucket                   string      `json:"bucket"`
	Prefix                   string      `json:"prefix,omitempty"`
	Active                   bool        `json:"active"`
	AddS3ObjectKeyAsLogField bool        `json:"addS3ObjectKeyAsLogField,omitempty"`
	Region                   AwsRegion   `json:"region"`
	LogsType                 AwsLogsType `json:"logsType"`
	Id                       int64       `json:"id,omitempty"`
}

type AwsRegion string
type AwsLogsType string

const (
	RegionUsEast1      AwsRegion = "US_EAST_1"
	RegionUsEast2      AwsRegion = "US_EAST_2"
	RegionUsWest1      AwsRegion = "US_WEST_1"
	RegionUsWest2      AwsRegion = "US_WEST_2"
	RegionEuWest1      AwsRegion = "EU_WEST_1"
	RegionEuWest2      AwsRegion = "EU_WEST_2"
	RegionEuWest3      AwsRegion = "EU_WEST_3"
	RegionEuCentral1   AwsRegion = "EU_CENTRAL_1"
	RegionApNortheast1 AwsRegion = "AP_NORTHEAST_1"
	RegionApNortheast2 AwsRegion = "AP_NORTHEAST_2"
	RegionApSoutheast1 AwsRegion = "AP_SOUTHEAST_1"
	RegionApSoutheast2 AwsRegion = "AP_SOUTHEAST_2"
	RegionSaEast1      AwsRegion = "SA_EAST_1"
	RegionApSouth1     AwsRegion = "AP_SOUTH_1"
	RegionCaCentral1   AwsRegion = "CA_CENTRAL_1"

	LogsTypeElb        AwsLogsType = "elb"
	LogsTypeVpcFlow    AwsLogsType = "vpcflow"
	LogsTypeS3Access   AwsLogsType = "S3Access"
	LogsTypeCloudfront AwsLogsType = "cloudfront"
)

func (r AwsRegion) String() string {
	return string(r)
}

func (t AwsLogsType) String() string {
	return string(t)
}

func New(apiToken, baseUrl string) (*S3FetcherClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &S3FetcherClient{
		Client: client.New(apiToken, baseUrl),
	}

	return c, nil
}

func validateCreateUpdateS3FetcherRequest(req S3FetcherRequest) error {
	if len(req.Bucket) == 0 {
		return fmt.Errorf("field bucket must be set")
	}

	if len(req.Region) == 0 {
		return fmt.Errorf("field region must be set")
	}

	if len(req.LogsType) == 0 {
		return fmt.Errorf("field logsType must be set")
	}

	if req.Active == nil {
		return fmt.Errorf("field active must be set")
	}

	if (req.AccessKey != "" && req.SecretKey == "") || (req.AccessKey == "" && req.SecretKey != "") {
		return fmt.Errorf("both aws keys must be set")
	}

	if req.AccessKey == "" && req.SecretKey == "" && req.Arn == "" {
		return fmt.Errorf("either keys or arn must be set")
	}

	err := isValidRegion(req.Region)
	if err != nil {
		return err
	}

	err = isValidLogsType(req.LogsType)
	if err != nil {
		return err
	}

	return nil
}

func GetValidRegions() []AwsRegion {
	return []AwsRegion{
		RegionUsEast1,
		RegionUsEast2,
		RegionUsWest1,
		RegionUsWest2,
		RegionEuWest1,
		RegionEuWest2,
		RegionEuWest3,
		RegionEuCentral1,
		RegionApNortheast1,
		RegionApNortheast2,
		RegionApSoutheast1,
		RegionApSoutheast2,
		RegionSaEast1,
		RegionApSouth1,
		RegionCaCentral1,
	}
}

func isValidRegion(region AwsRegion) error {
	validRegions := GetValidRegions()
	for _, validRegion := range validRegions {
		if validRegion == region {
			return nil
		}
	}

	return fmt.Errorf("invalid region. region must be one of: %s", validRegions)
}

func GetValidLogsType() []AwsLogsType {
	return []AwsLogsType{
		LogsTypeElb,
		LogsTypeVpcFlow,
		LogsTypeS3Access,
		LogsTypeCloudfront,
	}
}

func isValidLogsType(logsType AwsLogsType) error {
	validLogsTypes := GetValidLogsType()

	for _, validType := range validLogsTypes {
		if validType == logsType {
			return nil
		}
	}

	return fmt.Errorf("invalid logs type. logs type must be one of: %s", validLogsTypes)
}

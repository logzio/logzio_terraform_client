package s3_buckets_connector

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	s3BucketsServiceEndpoint string = "%s/v1/log-shipping/s3-buckets"

	s3BucketConnectorResourceName = "s3 bucket connector"

	operationCreateS3BucketConnector = "CreateS3BucketConnector"
	operationGetS3BucketConnector    = "GetS3BucketConnector"
	operationListS3BucketConnector   = "ListS3BucketConnector"
	operationUpdateS3BucketConnector = "UpdateS3BucketConnector"
	operationDeleteS3BucketConnector = "DeleteS3BucketConnector"
)

type S3BucketsConnectorClient struct {
	*client.Client
}

type S3BucketConnectorRequest struct {
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

type S3BucketConnectorResponse struct {
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

func New(apiToken, baseUrl string) (*S3BucketsConnectorClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &S3BucketsConnectorClient{
		Client: client.New(apiToken, baseUrl),
	}

	return c, nil
}

func validateCreateUpdateS3BucketRequest(req S3BucketConnectorRequest) error {
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

func isValidRegion(region AwsRegion) error {
	validRegions := []string{
		RegionUsEast1.String(),
		RegionUsEast2.String(),
		RegionUsWest1.String(),
		RegionUsWest2.String(),
		RegionEuWest1.String(),
		RegionEuWest2.String(),
		RegionEuWest3.String(),
		RegionEuCentral1.String(),
		RegionApNortheast1.String(),
		RegionApNortheast2.String(),
		RegionApSoutheast1.String(),
		RegionApSoutheast2.String(),
		RegionSaEast1.String(),
		RegionApSouth1.String(),
		RegionCaCentral1.String(),
	}

	if !logzio_client.Contains(validRegions, region.String()) {
		return fmt.Errorf("invalid region. region must be one of: %s", validRegions)
	}

	return nil
}

func isValidLogsType(logsType AwsLogsType) error {
	validLogsTypes := []string{
		LogsTypeElb.String(),
		LogsTypeVpcFlow.String(),
		LogsTypeS3Access.String(),
		LogsTypeCloudfront.String(),
	}

	if !logzio_client.Contains(validLogsTypes, logsType.String()) {
		return fmt.Errorf("invalid logs type. logs type must be one of: %s", validLogsTypes)
	}

	return nil
}

package s3_buckets_connector

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	s3BucketsServiceEndpoint string = "%s/v1/log-shipping/s3-buckets"

	s3BucketConnectorResourceName = "s3 bucket connector"

	operationCreateS3BucketConnector = "CreateS3BucketConnector"
	operationGetS3BucketConnector    = "GetS3BucketConnector"
	operationListS3BucketConnector   = "ListS3BucketConnector"
)

type S3BucketsConnectorClient struct {
	*client.Client
}

type S3BucketConnector struct {
	AccessKey                string      `json:"accessKey,omitempty"`
	SecretKey                string      `json:"secretKey,omitempty"`
	Arn                      string      `json:"arn,omitempty"`
	Bucket                   string      `json:"bucket"`
	Prefix                   string      `json:"prefix,omitempty"`
	Active                   *bool       `json:"active,omitempty"`
	AddS3ObjectKeyAsLogField *bool       `json:"addS3ObjectKeyAsLogField,omitempty"`
	Region                   AwsRegion   `json:"region"`
	LogsType                 AwsLogsType `json:"logsType"`
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

func validateCreateS3BucketRequest(req S3BucketConnector) error {
	if len(req.Bucket) == 0 {
		return fmt.Errorf("field bucket must be set")
	}

	return nil
}

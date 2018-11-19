package logzio_client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Client struct {
	apiToken string
	log      log.Logger
}

const ENV_LOGZIO_BASE_URL = "LOGZIO_BASE_URL"
const LOGZIO_BASE_URL string = "https://api.logz.io"

var logzioBaseUrl string = LOGZIO_BASE_URL

func getLogzioBaseUrl() string {
	if len(os.Getenv(ENV_LOGZIO_BASE_URL)) > 0 {
		logzioBaseUrl = os.Getenv(ENV_LOGZIO_BASE_URL)
	}
	return logzioBaseUrl
}

func New(apiToken string) *Client {
	var c Client
	c.apiToken = apiToken
	return &c
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func logSomething(function string, logStr string)
{
	log.Printf("%s::%s", function, logStr)
}

func addHttpHeaders(apiToken string, req *http.Request) {
	req.Header.Add("X-API-TOKEN", apiToken)
	req.Header.Add("Content-Type", "application/json")
}

type SeverityThresholdType struct {
	Severity  string `json:"severity"`
	Threshold int    `json:"threshold"`
}

type CreateAlertType struct {
    AlertNotificationEndpoints  []interface{} //required, can be empty if specified
    Description                 string //optional, can be blank if specified
    Filter                      string //optional, can't be blank if specified
    GroupByAggregationFields    []interface{}
    IsEnabled                   bool
	NotificationEmails          []interface{} //required, can be empty
	Operation                   string
	QueryString                 string //required, can't be blank
	SearchTimeFrameMinutes      int
	SeverityThresholdTiers      []SeverityThresholdType `json:"severityThresholdTiers"`
	SuppressNotificationMinutes int //optional, defaults to 0 if not specified
	Title                       string //required
	ValueAggregationField       interface{}
	ValueAggregationType        string
}

type AlertType struct {
	AlertId                     int64
	AlertNotificationEndpoints  []interface{} //required, can be empty if specified
	CreatedAt                   string
	CreatedBy                   string
	Description                 string
	Filter                      string
	GroupByAggregationFields    []interface{}
	IsEnabled                   bool
	LastTriggeredAt             interface{}
	LastUpdated                 string
	NotificationEmails          []interface{}
	Operation                   string
	QueryString                 string `json:"query_string"`
	SearchTimeFrameMinutes      int
	Severity                    string
	SeverityThresholdTiers      []SeverityThresholdType `json:"severityThresholdTiers"`
	SuppressNotificationMinutes int //optional, defaults to 0 if not specified
	Threshold                   int                     `json:"threshold"` // @todo: why is this a float64?
	Title                       string
	ValueAggregationField       interface{}
	ValueAggregationType        string
}

const (
	OperatorGreaterThanOrEquals string = "GREATER_THAN_OR_EQUALS"
	OperatorLessThanOrEquals    string = "LESS_THAN_OR_EQUALS"
	OperatorGreaterThan         string = "GREATER_THAN"
	OperatorLessThan            string = "LESS_THAN"
	OperatorNotEquals           string = "NOT_EQUALS"
	OperatorEquals              string = "EQUALS"

	AggregationTypeUniqueCount string = "UNIQUE_COUNT"
	AggregationTypeAvg         string = "AVG"
	AggregationTypeMax         string = "MAX"
	AggregationTypeNone        string = "NONE"
	AggregationTypeSum         string = "SUM"
	AggregationTypeCount       string = "COUNT"
	AggregationTypeMin         string = "MIN"

	SeverityHigh   string = "HIGH"
	SeverityLow    string = "LOW"
	SeverityMedium string = "MEDIUM"
)

func contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func checkValidStatus(response *http.Response, status []int) bool {
	for x := 0; x < len(status); x++ {
		if response.StatusCode == status[x] {
			return true
		}
	}
	return false
}

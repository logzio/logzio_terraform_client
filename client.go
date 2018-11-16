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
	log  log.Logger
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

func addHttpHeaders(apiToken string, req *http.Request) {
	req.Header.Add("X-API-TOKEN", apiToken)
	req.Header.Add("Content-Type", "application/json")
}

type SeverityThresholdType struct {
	Severity  string `json:"severity"`
	Threshold int    `json:"threshold"`
}

type CreateAlertType struct {
	Title                       string //required
	Description                 string //optional, can be blank if specified
	QueryString                 string //required, can't be blank
	Filter                      string //optional, can't be blank if specified
	Operation                   string
	SeverityThresholdTiers      []SeverityThresholdType `json:"severityThresholdTiers"`
	SearchTimeFrameMinutes      int
	NotificationEmails          []interface{} //required, can be empty
	IsEnabled                   bool
	SuppressNotificationMinutes int //optional, defaults to 0 if not specified
	ValueAggregationType        string
	ValueAggregationField       interface{}
	GroupByAggregationFields    []interface{}
	AlertNotificationEndpoints  []interface{} //required, can be empty if specified
}

type AlertType struct {
	AlertId int64
	Title string
	Severity string
	LastUpdated string
	CreatedAt string
	CreatedBy string
	Description string
	QueryString string `json:"query_string"`
	Filter string
	Operation string
	SeverityThresholdTiers []SeverityThresholdType `json:"severityThresholdTiers"`
	Threshold int     `json:"threshold"` // @todo: why is this a float64?
	SearchTimeFrameMinutes int
	NotificationEmails          []interface{}
	IsEnabled                   bool
	SuppressNotificationMinutes int //optional, defaults to 0 if not specified
	ValueAggregationType        string
	ValueAggregationField       interface{}
	GroupByAggregationFields    []interface{}
	AlertNotificationEndpoints  []interface{} //required, can be empty if specified
	LastTriggeredAt             interface{}
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

package logzio_client

import (
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

func logSomething(function string, logStr string) {
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
	AlertNotificationEndpoints   []interface{}
	Description                  string
	Filter                       string
	GroupByAggregationFields     []interface{}
	IsEnabled                    bool
	NotificationEmails           []interface{}
	Operation                    string
	QueryString                  string
	SearchTimeFrameMinutes       int
	SeverityThresholdTiers       []SeverityThresholdType `json:"severityThresholdTiers"`
	SuppressNotificationsMinutes int
	Title                        string
	ValueAggregationField        interface{}
	ValueAggregationType         string
}

type AlertType struct {
	AlertId                      int64
	AlertNotificationEndpoints   []interface{}
	CreatedAt                    string
	CreatedBy                    string
	Description                  string
	Filter                       string
	GroupByAggregationFields     []interface{}
	IsEnabled                    bool
	LastTriggeredAt              interface{}
	LastUpdated                  string
	NotificationEmails           []interface{}
	Operation                    string
	QueryString                  string `json:"query_string"`
	SearchTimeFrameMinutes       int
	Severity                     string
	SeverityThresholdTiers       []SeverityThresholdType `json:"severityThresholdTiers"`
	SuppressNotificationsMinutes int
	Threshold                    int `json:"threshold"`
	Title                        string
	ValueAggregationField        interface{}
	ValueAggregationType         string
}

const (
	AggregationTypeUniqueCount string = "UNIQUE_COUNT"
	AggregationTypeAvg         string = "AVG"
	AggregationTypeMax         string = "MAX"
	AggregationTypeNone        string = "NONE"
	AggregationTypeSum         string = "SUM"
	AggregationTypeCount       string = "COUNT"
	AggregationTypeMin         string = "MIN"

	OperatorGreaterThanOrEquals string = "GREATER_THAN_OR_EQUALS"
	OperatorLessThanOrEquals    string = "LESS_THAN_OR_EQUALS"
	OperatorGreaterThan         string = "GREATER_THAN"
	OperatorLessThan            string = "LESS_THAN"
	OperatorNotEquals           string = "NOT_EQUALS"
	OperatorEquals              string = "EQUALS"

	SeverityHigh   string = "HIGH"
	SeverityLow    string = "LOW"
	SeverityMedium string = "MEDIUM"

	alertNotificationEndpoints   string = "alertNotificationEndpoints"
	createdAt                    string = "createdAt"
	createdBy                    string = "createdBy"
	description                  string = "description"
	filter                       string = "filter"
	groupByAggregationFields     string = "groupByAggregationFields"
	isEnabled                    string = "isEnabled"
	queryString                  string = "query_string"
	lastTriggeredAt              string = "lastTriggeredAt"
	lastUpdated                  string = "lastUpdated"
	notificationEmails           string = "notificationEmails"
	operation                    string = "operation"
	searchTimeFrameMinutes       string = "searchTimeFrameMinutes"
	severity                     string = "severity"
	severityThresholdTiers       string = "severityThresholdTiers"
	suppressNotificationsMinutes string = "suppressNotificationsMinutes"
	threshold                    string = "threshold"
	title                        string = "title"
	valueAggregationField        string = "valueAggregationField"
	valueAggregationType         string = "valueAggregationType"
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

func jsonAlertToAlert(jsonAlert map[string]interface{}) AlertType {
	alert := AlertType{
		AlertId:                    int64(jsonAlert["alertId"].(float64)),
		AlertNotificationEndpoints: jsonAlert[alertNotificationEndpoints].([]interface{}),
		CreatedAt:                  jsonAlert[createdAt].(string),
		CreatedBy:                  jsonAlert[createdBy].(string),
		Description:                jsonAlert[description].(string),
		Filter:                     jsonAlert[filter].(string),
		IsEnabled:                  jsonAlert[isEnabled].(bool),
		LastUpdated:                jsonAlert[lastUpdated].(string),
		NotificationEmails:         jsonAlert[notificationEmails].([]interface{}),
		Operation:                  jsonAlert[operation].(string),
		QueryString:                jsonAlert[queryString].(string),
		Severity:                   jsonAlert[severity].(string),
		SearchTimeFrameMinutes:     int(jsonAlert[searchTimeFrameMinutes].(float64)),
		SeverityThresholdTiers:     []SeverityThresholdType{},
		Threshold:                  int(jsonAlert[threshold].(float64)),
		Title:                      jsonAlert[title].(string),
		ValueAggregationType:       jsonAlert[valueAggregationType].(string),
	}

	if jsonAlert[groupByAggregationFields] != nil {
		alert.GroupByAggregationFields = jsonAlert[groupByAggregationFields].([]interface{})
	}

	if jsonAlert[lastTriggeredAt] != nil {
		alert.LastTriggeredAt = jsonAlert[lastTriggeredAt].(interface{})
	}

	tiers := jsonAlert[severityThresholdTiers].([]interface{})
	for x := 0; x < len(tiers); x++ {
		tier := tiers[x].(map[string]interface{})
		threshold := SeverityThresholdType{
			Threshold: int(tier[threshold].(float64)),
			Severity:  tier[severity].(string),
		}
		alert.SeverityThresholdTiers = append(alert.SeverityThresholdTiers, threshold)
	}

	if jsonAlert[suppressNotificationsMinutes] != nil {
		alert.SuppressNotificationsMinutes = int(jsonAlert[suppressNotificationsMinutes].(float64))
	}

	if jsonAlert[valueAggregationField] != nil {
		alert.ValueAggregationField = jsonAlert[valueAggregationField].(interface{})
	}

	return alert
}

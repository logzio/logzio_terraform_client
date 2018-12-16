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

type EndpointType struct {
    Id int64
    EndpointType string
	Title string
    Description string
    Url string
    Message string
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

	fldAlertId                    string = "alertId"
	fldAlertNotificationEndpoints string = "alertNotificationEndpoints"
	fldCreatedAt               string = "createdAt"
	fldCreatedBy               string = "createdBy"
	fldDescription             string = "description"
	fldFilter                  string = "filter"
	fldGroupByAggregationFields   string = "groupByAggregationFields"
	fldIsEnabled                  string = "isEnabled"
	fldQueryString                string = "query_string"
	fldLastTriggeredAt            string = "lastTriggeredAt"
	fldLastUpdated                string = "lastUpdated"
	fldNotificationEmails         string = "notificationEmails"
	fldOperation                  string = "operation"
	fldSearchTimeFrameMinutes     string = "searchTimeFrameMinutes"
	fldSeverity                   string = "severity"
	fldSeverityThresholdTiers       string = "severityThresholdTiers"
	fldSuppressNotificationsMinutes string = "suppressNotificationsMinutes"
	fldThreshold                    string = "threshold"
	fldTitle                        string = "title"
	fldValueAggregationField        string = "valueAggregationField"
	fldValueAggregationType         string = "valueAggregationType"

	fldEndpointId string = "id"
	fldEndpointType string = "endpointType"
	fldEndpointTitle string = "title"
	fldEndpointDescription string = "description"
	fldEndpointUrl string = "url"

	endpointTypeSlack string = "slack"
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

func jsonEndpointToEndpoint(jsonEndpoint map[string]interface{}) EndpointType {
	endpoint := EndpointType{
        Id: int64(jsonEndpoint[fldEndpointId].(float64)),
		EndpointType:                  jsonEndpoint[fldEndpointType].(string),
		Title:                  jsonEndpoint[fldEndpointTitle].(string),
		Description:                  jsonEndpoint[fldEndpointDescription].(string),

	}
    return endpoint
}

func jsonAlertToAlert(jsonAlert map[string]interface{}) AlertType {
	alert := AlertType{
		AlertId:                    int64(jsonAlert[fldAlertId].(float64)),
		AlertNotificationEndpoints: jsonAlert[fldAlertNotificationEndpoints].([]interface{}),
		CreatedAt:                  jsonAlert[fldCreatedAt].(string),
		CreatedBy:                  jsonAlert[fldCreatedBy].(string),
		Description:                jsonAlert[fldDescription].(string),
		Filter:                     jsonAlert[fldFilter].(string),
		IsEnabled:                  jsonAlert[fldIsEnabled].(bool),
		LastUpdated:                jsonAlert[fldLastUpdated].(string),
		NotificationEmails:         jsonAlert[fldNotificationEmails].([]interface{}),
		Operation:                  jsonAlert[fldOperation].(string),
		QueryString:                jsonAlert[fldQueryString].(string),
		Severity:                   jsonAlert[fldSeverity].(string),
		SearchTimeFrameMinutes:     int(jsonAlert[fldSearchTimeFrameMinutes].(float64)),
		SeverityThresholdTiers:     []SeverityThresholdType{},
		Threshold:                  int(jsonAlert[fldThreshold].(float64)),
		Title:                      jsonAlert[fldTitle].(string),
		ValueAggregationType:       jsonAlert[fldValueAggregationType].(string),
	}

	if jsonAlert[fldGroupByAggregationFields] != nil {
		alert.GroupByAggregationFields = jsonAlert[fldGroupByAggregationFields].([]interface{})
	}

	if jsonAlert[fldLastTriggeredAt] != nil {
		alert.LastTriggeredAt = jsonAlert[fldLastTriggeredAt].(interface{})
	}

	tiers := jsonAlert[fldSeverityThresholdTiers].([]interface{})
	for x := 0; x < len(tiers); x++ {
		tier := tiers[x].(map[string]interface{})
		threshold := SeverityThresholdType{
			Threshold: int(tier[fldThreshold].(float64)),
			Severity:  tier[fldSeverity].(string),
		}
		alert.SeverityThresholdTiers = append(alert.SeverityThresholdTiers, threshold)
	}

	if jsonAlert[fldSuppressNotificationsMinutes] != nil {
		alert.SuppressNotificationsMinutes = int(jsonAlert[fldSuppressNotificationsMinutes].(float64))
	}

	if jsonAlert[fldValueAggregationField] != nil {
		alert.ValueAggregationField = jsonAlert[fldValueAggregationField].(interface{})
	}

	return alert
}

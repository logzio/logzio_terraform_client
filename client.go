package logzio_client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	name string
	log  log.Logger
}

func New(name string) *Client {
	var c Client
	c.name = name
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
	severity  string
	threshold int
}

type CreateAlertType struct {
	Title                       string //required
	Description                 string //optional, can be blank if specified
	QueryString                 string //required, can't be blank
	Filter                      string //optional, can't be blank if specified
	Operation                   string
	SeverityThresholdTiers      []SeverityThresholdType
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
}

const (
	GreaterThanOrEquals string = "GREATER_THAN_OR_EQUALS"
	LessThanOrEquals    string = "LESS_THAN_OR_EQUALS"
	GreaterThan         string = "GREATER_THAN"
	LessThan            string = "LESS_THAN"
	NotEquals           string = "NOT_EQUALS"
	Equals              string = "EQUALS"

	UniqueCount string = "UNIQUE_COUNT"
	Avg         string = "AVG"
	Max         string = "MAX"
	None        string = "NONE"
	Sum         string = "SUM"
	Count       string = "COUNT"
	Min         string = "MIN"
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

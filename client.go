package logzio_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

type Client struct {
	name string
	log log.Logger
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

func addHttpHeaders (apiToken string, req *http.Request) {
	req.Header.Add("X-API-TOKEN", apiToken)
	req.Header.Add("Content-Type", "application/json")
}

func buildCreateApiRequest(apiToken string, jsonObject map[string]interface{}) (*http.Request, error) {

	mybytes, err := json.Marshal(jsonObject)

	s,_ := prettyprint(mybytes)
	log.Printf("%s::%s::%s", "some_token", "buildCreateApiRequest", s)

	if err != nil {
		return nil, err
	}
	url := "https://api.logz.io/v1/alerts"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(mybytes))
	addHttpHeaders(apiToken, req)
	return req, err
}

func buildDeleteApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	url := fmt.Sprintf("https://api.logz.io/v1/alerts/%d", alertId)
	log.Printf("%s::%s::%s", "some_token", "DeleteAlert", url)
	req, err := http.NewRequest("DELETE", url, nil)
	addHttpHeaders(apiToken, req)
	return req, err
}

func buildListApiRequest(apiToken string) (*http.Request, error) {
	url := "https://api.logz.io/v1/alerts"
	req, err := http.NewRequest("GET", url, nil)
	addHttpHeaders(apiToken, req)
	return req, err
}

type SeverityThresholdType struct {
	severity string
	threshold int
}

type CreateAlertType struct {
	Title 						string //required
	Description 				string //optional, can be blank if specified
	QueryString 				string //required, can't be blank
	Filter 						string //optional, can't be blank if specified
	Operation 					string
	SeverityThresholdTiers 		[]SeverityThresholdType
	SearchTimeFrameMinutes 		int
	NotificationEmails 			[]interface{} //required, can be empty
	IsEnabled 					bool
	SuppressNotificationMinutes int //optional, defaults to 0 if not specified
	ValueAggregationType 		string
	ValueAggregationField 		interface{}
	GroupByAggregationFields 	[]interface{}
	AlertNotificationEndpoints 	[]interface{} //required, can be empty if specified
}

type AlertType struct {
	AlertId int64
}

const (
	GreaterThanOrEquals string = "GREATER_THAN_OR_EQUALS"
	LessThanOrEquals string = "LESS_THAN_OR_EQUALS"
	GreaterThan string = "GREATER_THAN"
	LessThan string = "LESS_THAN"
	NotEquals string = "NOT_EQUALS"
	Equals string = "EQUALS"

	UniqueCount string = "UNIQUE_COUNT"
	Avg string = "AVG"
	Max string = "MAX"
	None string = "NONE"
	Sum string = "SUM"
	Count string = "COUNT"
	Min string = "MIN"
)

func contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func validateCreateAlertRequest(alert CreateAlertType) (error) {

	if len(alert.Title) == 0 {
		return fmt.Errorf("title must be set")
	}

	if len(alert.QueryString) == 0 {
		return fmt.Errorf("query string must be set")
	}

	if alert.NotificationEmails == nil {
		return fmt.Errorf("notificationEmails must not be nil")
	}

	validAggregationTypes := []string {UniqueCount, Avg, Max, None, Sum, Count, Min}
	if !contains(validAggregationTypes, alert.ValueAggregationType) {
		return fmt.Errorf("valueAggregationType must be one of %s", validAggregationTypes)
	}

	validOperations := []string {GreaterThanOrEquals, LessThanOrEquals, GreaterThan, LessThan, NotEquals, Equals}
	if !contains(validOperations, alert.Operation) {
		return fmt.Errorf("operation must be one of %s", validOperations)
	}

	if None == alert.ValueAggregationType && (alert.ValueAggregationField != nil || alert.GroupByAggregationFields != nil) {
		return fmt.Errorf("if ValueAggregaionType is %s then ValueAggregationField and GroupByAggregationFields must be nil", None)
	}

	return nil
}

func buildCreateAlertRequest(alert CreateAlertType) (map[string]interface{}) {
	var createAlert = map[string]interface{}{}
	createAlert["title"] = alert.Title
	createAlert["description"] = alert.Description
	if len(alert.Filter) > 0 {
		createAlert["filter"] = alert.Filter
	}
	createAlert["query_string"] = alert.QueryString
	createAlert["operation"] = alert.Operation
	createAlert["severityThresholdTiers"] = alert.SeverityThresholdTiers
	createAlert["searchTimeFrameMinutes"] = alert.SearchTimeFrameMinutes
	createAlert["notificationEmails"] = alert.NotificationEmails
	createAlert["isEnabled"] = alert.IsEnabled
	createAlert["suppressNotificationMinutes"] = alert.SuppressNotificationMinutes
	createAlert["valueAggregationType"] = alert.ValueAggregationType
	createAlert["valueAggregationField"] = alert.ValueAggregationField
	createAlert["groupByAggregationFields"] = alert.GroupByAggregationFields
	createAlert["alertNotificationEndpoints"] = alert.AlertNotificationEndpoints
	return createAlert
}

func (n *Client) CreateAlert(alert CreateAlertType) (*AlertType, error) {

	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	createAlert := buildCreateAlertRequest(alert)

	/**
	createAlert := map[string]interface{}{
		"title" : alert.Title,
		"description" : alert.Description,
		"filter" : alert.Filter,
		"query_string" : alert.QueryString,
		"severityThresholdTiers" : alert.,
		"searchTimeFrameMinutes" : alert.SearchTimeFrameMinutes,
		"notificationEmails" : []string{"jon.boydell@massive.co"},
		"isEnabled" : alert.IsEnabled,
		"suppressNotificationMinutes" : alert.SuppressNotificationMinutes,
		"valueAggregationType" : alert.ValueAggregationType,
		"valueAggregationField" : nil,
		"groupByAggregationFields" : nil,
		"alertNotificationEndpoints" : alert.AlertNotificationEndpoints,
	}
	**/

	req, _ := buildCreateApiRequest(n.name, createAlert)

	var client http.Client
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", data)
	}

	var target AlertType
	json.Unmarshal(data, &target)

	return &target, nil
}

func (n *Client) DeleteAlert(alertId int64) error {
	req, _ := buildDeleteApiRequest(n.name, alertId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	data, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s", data)
	}

	return nil
}

func (n *Client) ListAlerts() ([]AlertType, error) {
	req, _ := buildListApiRequest(n.name)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", data)
	}

	var arr []AlertType
	err = json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}

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

func buildDeleteApiRequest(apiToken string, alertId string) (*http.Request, error) {
	url := fmt.Sprintf("https://api.logz.io/v1/alerts/%s", alertId)
	log.Print("%s::%s::%s", "some_token", "DeleteAlert", url)
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
	NotificationEmails 			[]interface{} //required, can't be blank
	IsEnabled 					bool
	SuppressNotificationMinutes int //optional, defaults to 0 if not specified
	ValueAggregationType 		string
	ValueAggregationField 		string
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
)

func (n *Client) CreateAlert(alert CreateAlertType) (*AlertType, error) {

	var createApiRequest map[string]interface{}

	if len(alert.Filter) > 0 {
		createApiRequest["filter"] = alert.Filter
	}

	createAlert := map[string]interface{}{
		"title" : alert.Title,
		"description" : alert.Description,
		"query_string" : alert.QueryString,
		"operation": alert.Operation,
		"severityThresholdTiers" : alert.SeverityThresholdTiers,
		"searchTimeFrameMinutes" : alert.SearchTimeFrameMinutes,
		"notificationEmails" : []string{"jon.boydell@massive.co"},
		"isEnabled" : alert.IsEnabled,
		"suppressNotificationMinutes" : alert.SuppressNotificationMinutes,
		"valueAggregationType" : alert.ValueAggregationType,
		"valueAggregationField" : nil,
		"groupByAggregationFields" : nil,
		"alertNotificationEndpoints" : alert.AlertNotificationEndpoints,
	}

	req, _ := buildCreateApiRequest(n.name, createAlert)

	var client http.Client
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", data)
	}

	var target map[string]interface{}
	json.Unmarshal(data, &target)

	v := new(AlertType)
	v.AlertId = target["alertId"].(int64)

	return v, nil
}

func (n *Client) DeleteAlert(alertId string) error {
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

	var arr []AlertType
	err = json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", data)
	}

	return arr, nil
}

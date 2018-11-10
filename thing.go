package logzio_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Thing struct {
	name string
}

type RequestTemplate struct {
	method string
	url string
}

var endpoints = map[string]RequestTemplate{
	"Alert::DeleteById": RequestTemplate{method: "POST", url:"http://api.logz.io/v1/alert"},
	"Alert::GetById": RequestTemplate{method: "GET", url: "http://api.logz.io/v1/alerts/:id"},
	"Alert::UpdateById": RequestTemplate{method: "PUT", url: "http://api.logz.io/v1/alerts/:id"},
	"Alert::Create": RequestTemplate{method: "POST", url: "https://api.logz.io/v1/alerts"},
}

func New(name string) *Thing {
	var t Thing
	t.name = name
	return &t
}

func (n *Thing) Sure() string {
	return n.name
}

func buildCreateApiRequest(apiToken string, jsonObject map[string]interface{}) (*http.Request, error) {
	mybytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(endpoints["Alert::Create"].method, endpoints["Alert::Create"].url, bytes.NewBuffer(mybytes))
	req.Header.Add("X-API-TOKEN", apiToken)
	req.Header.Add("Content-Type", "application/json")
	return req, err
}

	/**{
		"title": "Error level logs",
		"description": "Capture ERROR level logs in the given time range",
		"query_string": "loglevel:ERROR",
		Optional - "filter": "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
		"operation": "GREATER_THAN",
		"severityThresholdTiers": [{
		"severity": "HIGH",
		"threshold": 2
	},
		{
		"severity": "MEDIUM",
		"threshold": 1
		},
		{
		"severity": "LOW",
		"threshold": 0
		}
	],
		"searchTimeFrameMinutes": 5,
		"notificationEmails": [

		],
		"isEnabled": true,
		"suppressNotificationsMinutes": 0,
		"valueAggregationType": "NONE",
		"valueAggregationField": null,
		"groupByAggregationFields": [],
		Required (can be blank) - "alertNotificationEndpoints": []
	}*/

type SeverityThresholdType struct {
	severity string
	threshold int
}

type CreateAlertType struct {
	title 						string
	description 				string
	queryString 				string
	filter 						string //optional, non-blank if specified
	operation 					string
	severityThresholdTiers 		[]SeverityThresholdType
	searchTimeFrameMinutes 		int
	notificationEmails 			[]string
	isEnabled 					bool
	suppressNotificationMinutes int
	valueAggregationType 		string
	valueAggregationField 		string
	groupByAggregationFields 	[]interface{}
	alertNotificationEndpoints 	[]int //required, can be blank
}

type AlertType struct {}

func (n *Thing) CreateAlert(alertType CreateAlertType) (AlertType, error) {

	req, _ := buildCreateApiRequest(n.name, map[string]interface{}{
		"title" : alertType.title,
		"description" : alertType.description,
		"query_string" : alertType.queryString,
		"notificationEmails" : alertType.notificationEmails,
	})

	var client http.Client
	resp, _ := client.Do(req)

	s, _ := ioutil.ReadAll(resp.Body)

	alert := AlertType {

	}

	return alert, fmt.Errorf("%s", s)
}

func (n *Thing) DeleteAlert(id int) error {
	return nil
}

func (n *Thing) ListAlerts() ([]AlertType, error) {
	return []AlertType{}, nil
}

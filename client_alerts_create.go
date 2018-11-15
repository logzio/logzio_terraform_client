package logzio_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const createServiceUrl string = "%s/v1/alerts"
const createServiceMethod string = "POST"

func validateCreateAlertRequest(alert CreateAlertType) error {

	if len(alert.Title) == 0 {
		return fmt.Errorf("title must be set")
	}

	if len(alert.QueryString) == 0 {
		return fmt.Errorf("query string must be set")
	}

	if alert.NotificationEmails == nil {
		return fmt.Errorf("notificationEmails must not be nil")
	}

	validAggregationTypes := []string{UniqueCount, Avg, Max, None, Sum, Count, Min}
	if !contains(validAggregationTypes, alert.ValueAggregationType) {
		return fmt.Errorf("valueAggregationType must be one of %s", validAggregationTypes)
	}

	validOperations := []string{GreaterThanOrEquals, LessThanOrEquals, GreaterThan, LessThan, NotEquals, Equals}
	if !contains(validOperations, alert.Operation) {
		return fmt.Errorf("operation must be one of %s", validOperations)
	}

	if None == alert.ValueAggregationType && (alert.ValueAggregationField != nil || alert.GroupByAggregationFields != nil) {
		return fmt.Errorf("if ValueAggregaionType is %s then ValueAggregationField and GroupByAggregationFields must be nil", None)
	}

	return nil
}

func buildCreateAlertRequest(alert CreateAlertType) map[string]interface{} {
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

func buildCreateApiRequest(apiToken string, jsonObject map[string]interface{}) (*http.Request, error) {

	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	jsonStr, _ := prettyprint(jsonBytes)
	log.Printf("%s::%s::%s", "some_token", "buildCreateApiRequest", jsonStr)

	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(createServiceMethod, fmt.Sprintf(createServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	addHttpHeaders(apiToken, req)

	return req, err
}

func (c *Client) CreateAlert(alert CreateAlertType) (*AlertType, error) {

	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	createAlert := buildCreateAlertRequest(alert)
	req, _ := buildCreateApiRequest(c.apiToken, createAlert)

	var client http.Client
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)
	s, _ := prettyprint(data)

	if !checkValidStatus(resp, []int { 200 }) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "CreateAlert", resp.StatusCode, s)
	}

	var target AlertType
	json.Unmarshal(data, &target)

	return &target, nil
}
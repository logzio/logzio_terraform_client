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

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

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

	validAggregationTypes := []string{AggregationTypeUniqueCount, AggregationTypeAvg, AggregationTypeMax, AggregationTypeNone, AggregationTypeSum, AggregationTypeCount, AggregationTypeMin}
	if !contains(validAggregationTypes, alert.ValueAggregationType) {
		return fmt.Errorf("valueAggregationType must be one of %s", validAggregationTypes)
	}

	validOperations := []string{OperatorGreaterThanOrEquals, OperatorLessThanOrEquals, OperatorGreaterThan, OperatorLessThan, OperatorNotEquals, OperatorEquals}
	if !contains(validOperations, alert.Operation) {
		return fmt.Errorf("operation must be one of %s", validOperations)
	}

	validSeverities := []string{SeverityHigh, SeverityLow, SeverityMedium}
	for x := 0; x < len(alert.SeverityThresholdTiers); x++ {
		s := alert.SeverityThresholdTiers[x]
		if !contains(validSeverities, s.Severity) {
			return fmt.Errorf("severity must be one of %s", validSeverities)
		}
	}

	if AggregationTypeNone == alert.ValueAggregationType && (alert.ValueAggregationField != nil || alert.GroupByAggregationFields != nil) {
		message := fmt.Sprintf("if ValueAggregaionType is %s then ValueAggregationField and GroupByAggregationFields must be nil", AggregationTypeNone)
		return FieldError{"valueAggregationTypeComposite", message}
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

	//jsonStr, _ := prettyprint(jsonBytes)
	log.Printf("%s::%s", "buildCreateApiRequest", jsonBytes)

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

	log.Printf("%s::%s", "CreateAlert::Response", data)

	if !checkValidStatus(resp, []int{200}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "CreateAlert", resp.StatusCode, s)
	}

	var target AlertType
	json.Unmarshal(data, &target)

	return &target, nil
}

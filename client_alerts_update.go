package logzio_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const updateServiceUrl string = "%s/v1/alerts/%d"
const updateServiceMethod string = http.MethodPut

func buildUpdateAlertRequest(alert CreateAlertType) map[string]interface{} {
	var createAlert = map[string]interface{}{}
	createAlert[alertNotificationEndpoints] = alert.AlertNotificationEndpoints
	createAlert[description] = alert.Description
	if len(alert.Filter) > 0 {
		createAlert[filter] = alert.Filter
	}
	createAlert[groupByAggregationFields] = alert.GroupByAggregationFields
	createAlert[isEnabled] = alert.IsEnabled
	createAlert[queryString] = alert.QueryString
	createAlert[notificationEmails] = alert.NotificationEmails
	createAlert[operation] = alert.Operation
	createAlert[searchTimeFrameMinutes] = alert.SearchTimeFrameMinutes
	createAlert[severityThresholdTiers] = alert.SeverityThresholdTiers
	createAlert[suppressNotificationMinutes] = alert.SuppressNotificationMinutes
	createAlert[title] = alert.Title
	createAlert[valueAggregationField] = alert.ValueAggregationField
	createAlert[valueAggregationType] = alert.ValueAggregationType

	return createAlert
}

func buildUpdateApiRequest(apiToken string, alertId int64, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	logSomething("buildUpdateApiRequest", fmt.Sprintf("%s", jsonBytes))

	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(updateServiceMethod, fmt.Sprintf(updateServiceUrl, baseUrl, alertId), bytes.NewBuffer(jsonBytes))
	addHttpHeaders(apiToken, req)

	return req, err
}

func (c *Client) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error) {
	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	createAlert := buildUpdateAlertRequest(alert)
	req, _ := buildUpdateApiRequest(c.apiToken, alertId, createAlert)

	var client http.Client
	resp, _ := client.Do(req)
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	logSomething("UpdateAlert::Response", fmt.Sprintf("%s", jsonBytes))

	if !checkValidStatus(resp, []int{200}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "UpdateAlert", resp.StatusCode, jsonBytes)
	}

	var target AlertType
	json.Unmarshal(jsonBytes, &target)

	return &target, nil
}

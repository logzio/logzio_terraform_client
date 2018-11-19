package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const getServiceUrl string = "%s/v1/alerts/%d"
const getServiceMethod string = "GET"

func buildGetApiRequest(apiToken string, alertId int64) (*http.Request, error) {

	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(getServiceMethod, fmt.Sprintf(getServiceUrl, baseUrl, alertId), nil)

	addHttpHeaders(apiToken, req)
	return req, err
}

func (c *Client) GetAlert(alertId int64) (*AlertType, error) {
	req, _ := buildGetApiRequest(c.apiToken, alertId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s::%s", "GetAlert", data)

	if !checkValidStatus(resp, []int{200}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "GetAlert", resp.StatusCode, data)
	}

	str := fmt.Sprintf("%s", data)
	if strings.Contains(str, "no alert id") {
		return nil, fmt.Errorf("API call %s failed with missing alert %d, data: %s", "GetAlert", alertId, data)
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonResponse)

	alert := AlertType{
		AlertId:                    int64(jsonResponse["alertId"].(float64)),
		AlertNotificationEndpoints: jsonResponse["alertNotificationEndpoints"].([]interface{}),
		CreatedAt:                  jsonResponse["createdAt"].(string),
		CreatedBy:                  jsonResponse["createdBy"].(string),
		Description:                jsonResponse["description"].(string),
		Filter:                     jsonResponse["filter"].(string),
		IsEnabled:                  jsonResponse["isEnabled"].(bool),
		LastUpdated:                jsonResponse["lastUpdated"].(string),
		NotificationEmails:         jsonResponse["notificationEmails"].([]interface{}),
		Operation:                  jsonResponse["operation"].(string),
		QueryString:                jsonResponse["query_string"].(string),
		SearchTimeFrameMinutes:     int(jsonResponse["searchTimeFrameMinutes"].(float64)),
		Severity:                   jsonResponse["severity"].(string),
		SeverityThresholdTiers:     []SeverityThresholdType{},
		Threshold:                  int(jsonResponse["alertId"].(float64)),
		Title:                      jsonResponse["title"].(string),
		ValueAggregationType:       jsonResponse["valueAggregationType"].(string),
	}

	if jsonResponse["groupByAggregationFields"] != nil {
		alert.GroupByAggregationFields = jsonResponse["groupByAggregationFields"].([]interface{})
	}

	if jsonResponse["lastTriggeredAt"] != nil {
		alert.LastTriggeredAt = jsonResponse["lastTriggeredAt"].(interface{})
	}

	tiers := jsonResponse["severityThresholdTiers"].([]interface{})
	for x := 0; x < len(tiers); x++ {
		tier := tiers[x].(map[string]interface{})
		threshold := SeverityThresholdType{
			Threshold: int(tier["threshold"].(float64)),
			Severity:  tier["severity"].(string),
		}
		alert.SeverityThresholdTiers = append(alert.SeverityThresholdTiers, threshold)
	}

	if jsonResponse["suppressNotificationMinutes"] != nil {
		alert.SuppressNotificationMinutes = jsonResponse["suppressNotificationMinutes"].(int)
	}

	if jsonResponse["valueAggregationField"] != nil {
		alert.ValueAggregationField = jsonResponse["valueAggregationField"].(interface{})
	}

	if err != nil {
		return nil, err
	}

	return &alert, nil
}

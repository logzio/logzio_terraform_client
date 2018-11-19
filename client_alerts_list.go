package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const listServiceUrl string = "https://api.logz.io/v1/alerts"
const listServiceMethod string = http.MethodGet

func buildListApiRequest(apiToken string) (*http.Request, error) {
	req, err := http.NewRequest(listServiceMethod, listServiceUrl, nil)
	addHttpHeaders(apiToken, req)
	return req, err
}

func (c *Client) ListAlerts() ([]AlertType, error) {
	req, _ := buildListApiRequest(c.apiToken)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	s, _ := prettyprint(data)
	
	logSomething("ListAlerts", fmt.Sprintf("%s", s))

	if !checkValidStatus(resp, []int{200}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "ListAlerts", resp.StatusCode, data)
	}

	var arr []AlertType

	var jsonResponse []interface{}
	err = json.Unmarshal([]byte(data), &jsonResponse)

	for x := 0; x < len(jsonResponse); x++ {

		var jsonAlert map[string]interface{}
		jsonAlert = jsonResponse[x].(map[string]interface{})

		alert := AlertType{
			AlertId:                    int64(jsonAlert["alertId"].(float64)),
			AlertNotificationEndpoints: jsonAlert["alertNotificationEndpoints"].([]interface{}),
			CreatedAt:                  jsonAlert["createdAt"].(string),
			CreatedBy:                  jsonAlert["createdBy"].(string),
			Description:                jsonAlert["description"].(string),
			Filter:                     jsonAlert["filter"].(string),
			IsEnabled:                  jsonAlert["isEnabled"].(bool),
			LastUpdated:                jsonAlert["lastUpdated"].(string),
			NotificationEmails:         jsonAlert["notificationEmails"].([]interface{}),
			Operation:                  jsonAlert["operation"].(string),
			QueryString:                jsonAlert["query_string"].(string),
			Severity:                   jsonAlert["severity"].(string),
			SearchTimeFrameMinutes:     int(jsonAlert["searchTimeFrameMinutes"].(float64)),
			SeverityThresholdTiers:     []SeverityThresholdType{},
			Threshold:                  int(jsonAlert["alertId"].(float64)),
			Title:                      jsonAlert["title"].(string),
			ValueAggregationType:       jsonAlert["valueAggregationType"].(string),
		}

		if jsonAlert["groupByAggregationFields"] != nil {
			alert.GroupByAggregationFields = jsonAlert["groupByAggregationFields"].([]interface{})
		}

		if jsonAlert["lastTriggeredAt"] != nil {
			alert.LastTriggeredAt = jsonAlert["lastTriggeredAt"].(interface{})
		}

		tiers := jsonAlert["severityThresholdTiers"].([]interface{})
		for x := 0; x < len(tiers); x++ {
			tier := tiers[x].(map[string]interface{})
			threshold := SeverityThresholdType{
				Threshold: int(tier["threshold"].(float64)),
				Severity:  tier["severity"].(string),
			}
			alert.SeverityThresholdTiers = append(alert.SeverityThresholdTiers, threshold)
		}

		if jsonAlert["suppressNotificationMinutes"] != nil {
			alert.SuppressNotificationMinutes = jsonAlert["suppressNotificationMinutes"].(int)
		}

		if jsonAlert["valueAggregationField"] != nil {
			alert.ValueAggregationField = jsonAlert["valueAggregationField"].(interface{})
		}


		arr = append(arr, alert)
	}

	return arr, nil
}

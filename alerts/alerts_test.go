package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/jonboydell/logzio_client/client"
	"github.com/jonboydell/logzio_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupAlertsTest() (*alerts.AlertsClient, error, func()) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err, nil
	}

	underTest, err := alerts.New(apiToken)

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest.BaseUrl = server.URL
	underTest.Client.BaseUrl = server.URL

	return underTest, nil, func() {
		server.Close()
	}
}

func setupAlertsIntegrationTest() (*alerts.AlertsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := alerts.New(apiToken)
	underTest.BaseUrl = client.GetLogzIoBaseUrl()
	return underTest, nil
}

func createValidAlert() alerts.CreateAlertType {
	return alerts.CreateAlertType{
		Title:       "this is my title",
		Description: "this is my description",
		QueryString: "loglevel:ERROR",
		Filter:      "",
		Operation:   alerts.OperatorGreaterThan,
		SeverityThresholdTiers: []alerts.SeverityThresholdType{
			alerts.SeverityThresholdType{
				alerts.SeverityHigh,
				10,
			},
		},
		SearchTimeFrameMinutes:       0,
		NotificationEmails:           []interface{}{},
		IsEnabled:                    true,
		SuppressNotificationsMinutes: 0,
		ValueAggregationType:         alerts.AggregationTypeCount,
		ValueAggregationField:        nil,
		GroupByAggregationFields:     []interface{}{"my_field"},
		AlertNotificationEndpoints:   []interface{}{},
	}
}

func createUpdateAlert() alerts.CreateAlertType {
	return alerts.CreateAlertType{
		Title:       "this is my updated title",
		Description: "this is my description",
		QueryString: "loglevel:ERROR",
		Filter:      "",
		Operation:   alerts.OperatorGreaterThan,
		SeverityThresholdTiers: []alerts.SeverityThresholdType{
			alerts.SeverityThresholdType{
				alerts.SeverityHigh,
				10,
			},
		},
		SearchTimeFrameMinutes:       0,
		NotificationEmails:           []interface{}{},
		IsEnabled:                    true,
		SuppressNotificationsMinutes: 0,
		ValueAggregationType:         alerts.AggregationTypeCount,
		ValueAggregationField:        nil,
		GroupByAggregationFields:     []interface{}{"my_field"},
		AlertNotificationEndpoints:   []interface{}{},
	}
}

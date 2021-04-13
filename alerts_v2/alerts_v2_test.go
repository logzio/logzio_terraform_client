package alerts_v2_test

import (
	"github.com/logzio/logzio_terraform_client/alerts_v2"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupAlertsTest() (*alerts_v2.AlertsV2Client, error, func()) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, _ := alerts_v2.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func getCreateAlertType() alerts_v2.CreateAlertType {
	alertQuery := alerts_v2.AlertQuery{
		Query:                    "loglevel:ERROR",
		Aggregation:              alerts_v2.AggregationObj{AggregationType: alerts_v2.AggregationTypeCount},
		ShouldQueryOnAllAccounts: true,
	}

	trigger := alerts_v2.AlertTrigger{
		Operator:               alerts_v2.OperatorEquals,
		SeverityThresholdTiers: map[string]float32{alerts_v2.SeverityHigh: 10, alerts_v2.SeverityInfo: 5},
	}
	
	subComponent := alerts_v2.SubAlert{
		QueryDefinition: alertQuery,
		Trigger:         trigger,
		Output:          alerts_v2.SubAlertOutput{},
	}

	createAlertType := alerts_v2.CreateAlertType{
		Title:                  "test create alert",
		Description:            "this is my description",
		Tags:                   []string{"some", "words"},
		Output:                 alerts_v2.AlertOutput{},
		SubComponents:          []alerts_v2.SubAlert{subComponent},
		Correlations:           alerts_v2.SubAlertCorrelation{},
		Enabled:                true,
	}

	return createAlertType
}

func getAlertType() alerts_v2.AlertType {
	alertQuery := alerts_v2.AlertQuery{
		Query:                    "loglevel:ERROR",
		Aggregation:              alerts_v2.AggregationObj{AggregationType: alerts_v2.AggregationTypeCount},
		ShouldQueryOnAllAccounts: true,
	}

	trigger := alerts_v2.AlertTrigger{
		Operator:               alerts_v2.OperatorEquals,
		SeverityThresholdTiers: map[string]float32{alerts_v2.SeverityHigh: 10, alerts_v2.SeverityInfo: 5},
	}

	subComponent := alerts_v2.SubAlert{
		QueryDefinition: alertQuery,
		Trigger:         trigger,
		Output:          alerts_v2.SubAlertOutput{},
	}

	alert := alerts_v2.AlertType{
		AlertId: int64(1234567),
		Title:                  "test alert",
		Description:            "this is my description",
		Tags:                   []string{"some", "words"},
		Output:                 alerts_v2.AlertOutput{},
		SubComponents:          []alerts_v2.SubAlert{subComponent},
		Correlations:           alerts_v2.SubAlertCorrelation{},
		Enabled:                true,
	}

	return alert
}

func TestNewWithEmptyBaseUrl(t *testing.T) {
	_, err := alerts_v2.New("any-api-token", "")
	if err == nil {
		t.Fatal("Expected error when baseUrl is empty")
	}
	if err.Error() != "Base URL not defined" {
		t.Fatalf("The expected error message to be '%s' but was '%s'",
			"Base URL not defined", err.Error())
	}
}

func TestNewWithEmptyApiToken(t *testing.T) {
	_, err := alerts_v2.New("", "any-base-url")

	if err == nil {
		t.Fatal("Expected error when API token is empty")
	}
	if err.Error() != "API token not defined" {
		t.Fatalf("The expected error message to be '%s' but was '%s'",
			"API token not defined", err.Error())
	}
}

func setupAlertsV2IntegrationTest() (*alerts_v2.AlertsV2Client, error){
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := alerts_v2.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}
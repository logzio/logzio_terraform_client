package grafana_notification_policies_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_notification_policies"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
)

const (
	grafanaDefaultReceiver = "grafana-default-email"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupGrafanaNotificationPolicyIntegrationTest() (*grafana_notification_policies.GrafanaNotificationPolicyClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := grafana_notification_policies.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func setupGrafanaNotificationPolicyTest() (*grafana_notification_policies.GrafanaNotificationPolicyClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := grafana_notification_policies.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func getGrafanaNotificationPolicyObject() grafana_notification_policies.GrafanaNotificationPolicyTree {
	return grafana_notification_policies.GrafanaNotificationPolicyTree{
		GroupBy:        []string{"hello-world", "alertname"},
		GroupInterval:  "5m",
		GroupWait:      "10s",
		Receiver:       grafanaDefaultReceiver,
		RepeatInterval: "5h",
		Routes: []grafana_notification_policies.GrafanaNotificationPolicy{
			{
				Receiver:       grafanaDefaultReceiver,
				ObjectMatchers: grafana_notification_policies.MatchersObj{grafana_notification_policies.MatcherObj{"hello", "=", "darkness"}},
				Continue:       true,
			},
			{
				Receiver:       grafanaDefaultReceiver,
				ObjectMatchers: grafana_notification_policies.MatchersObj{grafana_notification_policies.MatcherObj{"my", "=~", "oldfriend.*"}},
				Continue:       false,
			},
		},
	}
}

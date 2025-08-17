package metrics_rollup_rules_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/logzio/logzio_terraform_client/metrics_rollup_rules"
	"github.com/logzio/logzio_terraform_client/test_utils"
)

const metricsRollupRulesPath = "/v1/metrics-management/rollup-rules"

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

func setupMetricsRollupRulesTest() (*metrics_rollup_rules.MetricsRollupRulesClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := metrics_rollup_rules.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupMetricsRollupRulesIntegrationTest() (*metrics_rollup_rules.MetricsRollupRulesClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := metrics_rollup_rules.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

// Integration helpers (kept within this test file per project conventions)
func getCreateRollupRule() (metrics_rollup_rules.CreateUpdateRollupRule, error) {
	accountId, err := test_utils.GetMetricsAccountId()
	if err != nil {
		return metrics_rollup_rules.CreateUpdateRollupRule{}, err
	}

	metricName := "integration_test_rollup_metric_" + time.Now().Format("20060102150405")

	return metrics_rollup_rules.CreateUpdateRollupRule{
		AccountId:               accountId,
		MetricName:              metricName,
		MetricType:              metrics_rollup_rules.MetricTypeGauge,
		RollupFunction:          metrics_rollup_rules.AggLast,
		LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
		Labels:                  []string{"env"},
	}, nil
}

func getBulkCreateRollupRules() ([]metrics_rollup_rules.CreateUpdateRollupRule, error) {
	accountId, err := test_utils.GetMetricsAccountId()
	if err != nil {
		return nil, err
	}
	ts := time.Now().Format("20060102150405")
	return []metrics_rollup_rules.CreateUpdateRollupRule{
		{
			AccountId:               accountId,
			MetricName:              "integration_bulk_rollup_a_" + ts,
			MetricType:              metrics_rollup_rules.MetricTypeGauge,
			RollupFunction:          metrics_rollup_rules.AggLast,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"env"},
		},
		{
			AccountId:               accountId,
			MetricName:              "integration_bulk_rollup_b_" + ts,
			MetricType:              metrics_rollup_rules.MetricTypeGauge,
			RollupFunction:          metrics_rollup_rules.AggCount,
			LabelsEliminationMethod: metrics_rollup_rules.LabelsExcludeBy,
			Labels:                  []string{"env"},
		},
	}, nil
}

func getSearchRollupRulesRequest() (metrics_rollup_rules.SearchRollupRulesRequest, error) {
	accountId, err := test_utils.GetMetricsAccountId()
	if err != nil {
		return metrics_rollup_rules.SearchRollupRulesRequest{}, err
	}
	return metrics_rollup_rules.SearchRollupRulesRequest{
		Filter: &metrics_rollup_rules.SearchFilter{
			AccountIds: []int64{accountId},
		},
		Pagination: &metrics_rollup_rules.Pagination{
			PageNumber: 1,
			PageSize:   100,
		},
	}, nil
}

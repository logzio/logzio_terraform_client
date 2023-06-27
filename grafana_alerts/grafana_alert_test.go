package grafana_alerts_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/grafana_alerts"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

const (
	envGrafanaFolderUid = "GRAFANA_FOLDER_UID"
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

func setupGrafanaAlertRuleTest() (*grafana_alerts.GrafanaAlertClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := grafana_alerts.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupGrafanaFolderIntegrationTest() (*grafana_alerts.GrafanaAlertClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := grafana_alerts.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func getGrafanaAlertRuleObject() grafana_alerts.GrafanaAlertRule {
	data := grafana_alerts.GrafanaAlertQuery{
		DatasourceUid: "__expr__",
		Model:         json.RawMessage(`{"conditions":[{"evaluator":{"params":[0,0],"type":"gt"},"operator":{"type":"and"},"query":{"params":[]},"reducer":{"params":[],"type":"avg"},"type":"query"}],"datasource":{"type":"__expr__","uid":"__expr__"},"expression":"1 == 1","hide":false,"intervalMs":1000,"maxDataPoints":43200,"refId":"A","type":"math"}`),
		RefId:         "A",
		RelativeTimeRange: grafana_alerts.RelativeTimeRangeObj{
			From: 0,
			To:   0,
		},
	}

	alertRule := grafana_alerts.GrafanaAlertRule{
		Annotations: map[string]string{"key_test": "value_test"},
		Condition:   "A",
		Data:        []*grafana_alerts.GrafanaAlertQuery{&data},
		FolderUID:   os.Getenv(envGrafanaFolderUid),
		For:         "",
		Id:          0,
		Labels:      nil,
		NoDataState: "",
		OrgID:       0,
		Provenance:  "",
		RuleGroup:   "",
		Title:       "",
		Uid:         "",
		Updated:     time.Time{},
	}
}

func getModel() interface{} {
	modelJson := "{\"conditions\":[{\"evaluator\":{\"params\":[0,0],\"type\":\"gt\"},\"operator\":{\"type\":\"and\"},\"query\":{\"params\":[]},\"reducer\":{\"params\":[],\"type\":\"avg\"},\"type\":\"query\"}],\"datasource\":{\"type\":\"__expr__\",\"uid\":\"__expr__\"},\"expression\":\"1 == 1\",\"hide\":false,\"intervalMs\":1000,\"maxDataPoints\":43200,\"refId\":\"A\",\"type\":\"math\"}"
	var modelObj map[string]interface{}
	_ = json.Unmarshal([]byte(modelJson), &modelObj)
	return modelObj
}

func getTestFolderUid() string {
}

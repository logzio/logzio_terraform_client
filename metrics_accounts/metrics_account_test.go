package metrics_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/metrics_accounts"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func setupMetricsAccountsTest() (*metrics_accounts.MetricsAccountClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := metrics_accounts.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupMetricsAccountsIntegrationTest() (*metrics_accounts.MetricsAccountClient, string, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, "", err
	}

	email, err := test_utils.GetLogzioEmail()
	if err != nil {
		return nil, "", err
	}

	underTest, err := metrics_accounts.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, email, err
}

func getCreateOrUpdateMetricsAccount(email string) metrics_accounts.CreateOrUpdateMetricsAccount {
	metricsAccount := metrics_accounts.CreateOrUpdateMetricsAccount{
		Email:                 email,
		AccountName:           "tf_client_test",
		PlanUts:               new(int32),
		AuthorizedAccountsIds: []int32{},
	}

	*metricsAccount.PlanUts = 100
	return metricsAccount
}

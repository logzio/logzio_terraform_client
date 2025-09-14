package sub_accounts_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"

	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/logzio/logzio_terraform_client/test_utils"
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

func setupSubAccountsTest() (*sub_accounts.SubAccountClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := sub_accounts.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupSubAccountsIntegrationTestHandler(apiToken string) (*sub_accounts.SubAccountClient, string, error) {
	email, err := test_utils.GetLogzioEmail()
	if err != nil {
		return nil, "", err
	}

	underTest, err := sub_accounts.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, email, err
}

func setupSubAccountsIntegrationTest() (*sub_accounts.SubAccountClient, string, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, "", err
	}

	return setupSubAccountsIntegrationTestHandler(apiToken)
}

func setupSubAccountsWarmIntegrationTest() (*sub_accounts.SubAccountClient, string, error) {
	apiToken, err := test_utils.GetWarmApiToken()
	if err != nil {
		return nil, "", err
	}

	return setupSubAccountsIntegrationTestHandler(apiToken)
}

func setupSubAccountConsumptionIntegrationTest() (*sub_accounts.SubAccountClient, string, error) {
	apiToken, err := test_utils.GetConsumptionApiToken()
	if err != nil {
		return nil, "", err
	}

	return setupSubAccountsIntegrationTestHandler(apiToken)
}

func getCreateOrUpdateSubAccount(email string) sub_accounts.CreateOrUpdateSubAccount {
	subAccount := sub_accounts.CreateOrUpdateSubAccount{
		Email:                  email,
		AccountName:            "tf_client_test",
		MaxDailyGB:             new(float32),
		RetentionDays:          1,
		Searchable:             strconv.FormatBool(false),
		Accessible:             strconv.FormatBool(true),
		SharingObjectsAccounts: []int32{},
		DocSizeSetting:         strconv.FormatBool(false),
		UtilizationSettings: sub_accounts.AccountUtilizationSettingsCreateOrUpdate{
			FrequencyMinutes:   3,
			UtilizationEnabled: strconv.FormatBool(true),
		},
	}

	*subAccount.MaxDailyGB = 1
	return subAccount
}

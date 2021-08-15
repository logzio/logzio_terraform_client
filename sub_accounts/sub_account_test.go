package sub_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/sub_accounts"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func setupSubAccountsTest() (*sub_accounts.SubAccountClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := sub_accounts.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupSubAccountsIntegrationTest() (*sub_accounts.SubAccountClient, string, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, "", err
	}

	email, err := test_utils.GetLogzioEmail()
	if err != nil {
		return nil, "", err
	}

	underTest, err := sub_accounts.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, email, err
}

func getCreatrOrUpdateSubAccount(email string) sub_accounts.CreateOrUpdateSubAccount {
	return sub_accounts.CreateOrUpdateSubAccount{
		Email:                  email,
		AccountName:            "tf_client_test",
		MaxDailyGB:             1,
		RetentionDays:          1,
		Searchable:             strconv.FormatBool(false),
		Accessible:             strconv.FormatBool(true),
		SharingObjectsAccounts: []int32{},
		DocSizeSetting:         strconv.FormatBool(false),
	}
}

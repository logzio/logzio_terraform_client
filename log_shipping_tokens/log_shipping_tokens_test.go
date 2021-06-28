package log_shipping_tokens_test

import (
	"github.com/logzio/logzio_terraform_client/log_shipping_tokens"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"strconv"
)

func setupLogShippingTokensIntegrationTest() (*log_shipping_tokens.LogShippingTokensClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := log_shipping_tokens.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}

func getCreateLogShippingToken() log_shipping_tokens.CreateLogShippingToken {
	return log_shipping_tokens.CreateLogShippingToken{
		Name:    "client_integration_test",
		Enabled: strconv.FormatBool(true),
	}
}
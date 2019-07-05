package alerts_test

import (
	"github.com/jonboydell/logzio_client/alerts"
	"github.com/jonboydell/logzio_client/test_utils"
)

func setupAlertsTest() (*alerts.Alerts, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := alerts.New(apiToken)
	return underTest, nil
}

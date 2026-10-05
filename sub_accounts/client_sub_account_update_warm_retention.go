package sub_accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	updateWarmRetentionServiceUrl      = subAccountServiceEndpoint + "/%d/retention-details"
	updateWarmRetentionServiceMethod   = http.MethodPut
	updateWarmRetentionServiceSuccess  = http.StatusOK
	updateWarmRetentionServiceNotFound = http.StatusNotFound
)

// UpdateWarmRetention sets the warm tier retention, in days, of one account: the main account of the API token or one
// of its sub accounts. It returns the hot and warm retention of the main account and all its sub accounts after the
// update, an error otherwise.
// The main account must already have warm tier and keep at least 1 day of it; a sub account can be set from 0 (no
// warm tier) up to the main account's warm retention. The API rejects anything else with 400.
func (c *SubAccountClient) UpdateWarmRetention(accountId int64, updateWarmRetention UpdateWarmRetention) ([]AccountRetentionDetails, error) {
	err := validateUpdateWarmRetention(updateWarmRetention)
	if err != nil {
		return nil, err
	}

	updateWarmRetentionJson, err := json.Marshal(updateWarmRetention)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateWarmRetentionServiceMethod,
		Url:          fmt.Sprintf(updateWarmRetentionServiceUrl, c.BaseUrl, accountId),
		Body:         updateWarmRetentionJson,
		SuccessCodes: []int{updateWarmRetentionServiceSuccess},
		NotFoundCode: updateWarmRetentionServiceNotFound,
		ResourceId:   accountId,
		ApiAction:    operationUpdateWarmRetention,
		ResourceName: subAccountResourceName,
	})

	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("failed with missing retention details")
	}

	var retentionDetails []AccountRetentionDetails
	err = json.Unmarshal(res, &retentionDetails)
	if err != nil {
		return nil, err
	}

	return retentionDetails, nil
}

func validateUpdateWarmRetention(updateWarmRetention UpdateWarmRetention) error {
	if updateWarmRetention.SnapSearchRetentionDays < 0 {
		return fmt.Errorf("snapSearchRetentionDays should be non-negative")
	}
	return nil
}

package security_rules

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	retrieveFilteredSecurityRulesServiceUrl            = securityRulesServiceEndpoint + "/search"
	retrieveFilteredSecurityRulesServiceMethod  string = http.MethodPost
	retrieveFilteredSecurityRulesMethodSuccess  int    = http.StatusOK
	retrieveFilteredSecurityRulesMethodNotFound int    = http.StatusNotFound
)

// RetrieveFilteredSecurityRules returns security rules in accordance to the filter that was provided, an error otherwise
func (c *SecurityRulesClient) RetrieveFilteredSecurityRules(filter RetrieveFiltered) (*FilteredSecurityRules, error) {
	err := validateRetrieveFiltered(filter)
	if err != nil {
		return nil, err
	}

	filterLogsJson, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   retrieveFilteredSecurityRulesServiceMethod,
		Url:          fmt.Sprintf(retrieveFilteredSecurityRulesServiceUrl, c.BaseUrl),
		Body:         filterLogsJson,
		SuccessCodes: []int{retrieveFilteredSecurityRulesMethodSuccess},
		NotFoundCode: retrieveFilteredSecurityRulesMethodNotFound,
		ResourceId:   nil,
		ApiAction:    securityRuleOperationRetrieveFiltered,
		ResourceName: securityRuleResourceName,
	})

	if err != nil {
		return nil, err
	}

	var rules FilteredSecurityRules
	err = json.Unmarshal(res, &rules)
	if err != nil {
		return nil, err
	}

	return &rules, nil
}

func validateRetrieveFiltered(filter RetrieveFiltered) error {
	validSorts := GetValidFilteredSorts()
	if len(filter.Sort.SortByField) > 0 {
		if !logzio_client.Contains(validSorts, filter.Sort.SortByField) {
			return fmt.Errorf("invalid sort by field. should be one of: %s", validSorts)
		}
	}

	return nil
}

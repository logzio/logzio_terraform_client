package security_rules

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateSecurityRuleServiceUrl     = securityRulesServiceEndpoint + "/%d"
	updateSecurityRuleServiceMethod  = http.MethodPut
	updateSecurityRuleMethodSuccess  = http.StatusOK
	updateSecurityRuleMethodNotFound = http.StatusNotFound
)

// UpdateSecurityRule updates an existing security rule, based on the supplied security rule identifier, using the parameters of the specified rule
// Returns the updated security rule if successful, an error otherwise
func (c *SecurityRulesClient) UpdateSecurityRule(ruleId int32, rule CreateUpdateSecurityRule) (*SecurityRule, error) {
	err := validateCreateUpdateSecurityRule(rule)
	if err != nil {
		return nil, err
	}

	updateRuleJson, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateSecurityRuleServiceMethod,
		Url:          fmt.Sprintf(updateSecurityRuleServiceUrl, c.BaseUrl, ruleId),
		Body:         updateRuleJson,
		SuccessCodes: []int{updateSecurityRuleMethodSuccess},
		NotFoundCode: updateSecurityRuleMethodNotFound,
		ResourceId:   ruleId,
		ApiAction:    securityRuleOperationUpdate,
		ResourceName: securityRuleResourceName,
	})

	if err != nil {
		return nil, err
	}

	var target SecurityRule
	err = json.Unmarshal(res, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}

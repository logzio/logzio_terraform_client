package security_rules

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	deleteSecurityRuleServiceMethod      = http.MethodDelete
	deleteSecurityRuleServiceUrl         = securityRulesServiceEndpoint + "/%d"
	deleteSecurityRuleMethodSuccess  int = http.StatusOK
	deleteSecurityRuleMethodNotFound int = http.StatusNotFound
)

// DeleteSecurityRule deletes security rule specified by its unique id, returns an error if a problem occurred
func (c *SecurityRulesClient) DeleteSecurityRule(ruleId int32) (*SecurityRule, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   deleteSecurityRuleServiceMethod,
		Url:          fmt.Sprintf(deleteSecurityRuleServiceUrl, c.BaseUrl, ruleId),
		Body:         nil,
		SuccessCodes: []int{deleteSecurityRuleMethodSuccess},
		NotFoundCode: deleteSecurityRuleMethodNotFound,
		ResourceId:   ruleId,
		ApiAction:    securityRuleOperationDelete,
		ResourceName: securityRuleResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal SecurityRule
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

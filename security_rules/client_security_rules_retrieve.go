package security_rules

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getSecurityRuleServiceUrl            = securityRulesServiceEndpoint + "/%d"
	getSecurityRuleServiceMethod  string = http.MethodGet
	getSecurityRuleMethodSuccess  int    = http.StatusOK
	getSecurityRuleMethodNotFound int    = http.StatusNotFound
)

// RetrieveSecurityRule returns a security rule given its unique identifier, an error otherwise
func (c *SecurityRulesClient) RetrieveSecurityRule(ruleId int32) (*SecurityRule, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getSecurityRuleServiceMethod,
		Url:          fmt.Sprintf(getSecurityRuleServiceUrl, c.BaseUrl, ruleId),
		Body:         nil,
		SuccessCodes: []int{getSecurityRuleMethodSuccess},
		NotFoundCode: getSecurityRuleMethodNotFound,
		ResourceId:   ruleId,
		ApiAction:    securityRuleOperationRetrieve,
		ResourceName: securityRuleResourceName,
	})

	if err != nil {
		return nil, err
	}

	var rule SecurityRule
	err = json.Unmarshal(res, &rule)
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

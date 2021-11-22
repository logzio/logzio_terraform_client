package security_rules

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	disableSecurityRuleServiceUrl             = securityRulesServiceEndpoint + "/%d/disable"
	disableSecurityRuleServiceMethod   string = http.MethodPost
	disableSecurityRuleMethodNoContent int    = http.StatusNoContent
	disableSecurityRuleMethodNotFound  int    = http.StatusNotFound
)

// DisableSecurityRule disables a security rule given unique identifier, an error otherwise
func (c *SecurityRulesClient) DisableSecurityRule(ruleId int32) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   disableSecurityRuleServiceMethod,
		Url:          fmt.Sprintf(disableSecurityRuleServiceUrl, c.BaseUrl, ruleId),
		Body:         nil,
		SuccessCodes: []int{disableSecurityRuleMethodNoContent},
		NotFoundCode: disableSecurityRuleMethodNotFound,
		ResourceId:   ruleId,
		ApiAction:    securityRuleOperationDisable,
		ResourceName: securityRuleResourceName,
	})

	return err
}

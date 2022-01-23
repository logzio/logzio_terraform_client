package security_rules

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	enableSecurityRuleServiceUrl             = securityRulesServiceEndpoint + "/%d/enable"
	enableSecurityRuleServiceMethod   string = http.MethodPost
	enableSecurityRuleMethodNoContent int    = http.StatusNoContent
	enableSecurityRuleMethodNotFound  int    = http.StatusNotFound
)

// EnableSecurityRule enables a security rule given unique identifier, an error otherwise
func (c *SecurityRulesClient) EnableSecurityRule(ruleId int32) error {
	_, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   enableSecurityRuleServiceMethod,
		Url:          fmt.Sprintf(enableSecurityRuleServiceUrl, c.BaseUrl, ruleId),
		Body:         nil,
		SuccessCodes: []int{enableSecurityRuleMethodNoContent},
		NotFoundCode: enableSecurityRuleMethodNotFound,
		ResourceId:   ruleId,
		ApiAction:    securityRuleOperationEnable,
		ResourceName: securityRuleResourceName,
	})

	return err
}

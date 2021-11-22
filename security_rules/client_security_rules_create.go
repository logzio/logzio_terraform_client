package security_rules

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createSecurityRuleServiceUrl            = securityRulesServiceEndpoint
	createSecurityRuleServiceMethod         = http.MethodPost
	createSecurityRuleServiceCreated        = http.StatusCreated
	createSecurityRuleServiceStatusNotFound = http.StatusNotFound
)

func (c *SecurityRulesClient) CreateSecurityRule(createRule CreateUpdateSecurityRule) (*SecurityRule, error) {
	err := validateCreateUpdateSecurityRule(createRule)
	if err != nil {
		return nil, err
	}

	createRuleJson, err := json.Marshal(createRule)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createSecurityRuleServiceMethod,
		Url:          fmt.Sprintf(createSecurityRuleServiceUrl, c.BaseUrl),
		Body:         createRuleJson,
		SuccessCodes: []int{createSecurityRuleServiceCreated},
		NotFoundCode: createSecurityRuleServiceStatusNotFound,
		ResourceId:   nil,
		ApiAction:    securityRuleOperationCreate,
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

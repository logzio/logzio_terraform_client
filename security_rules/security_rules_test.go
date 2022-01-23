package security_rules_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/security_rules"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

const (
	securityRulesApiBasePath = "/v2/security/rules"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupSecurityRulesTest() (*security_rules.SecurityRulesClient, func(), error) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, err := security_rules.New(apiToken, server.URL)
	if err != nil {
		return nil, nil, err
	}

	return underTest, func() {
		server.Close()
	}, nil
}

func setupSecurityRulesIntegrationTest() (*security_rules.SecurityRulesClient, error) {
	apiToken, err := test_utils.GetLogzioSecurityApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := security_rules.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}

func getCreateUpdateRule() security_rules.CreateUpdateSecurityRule {
	recipients := security_rules.RuleRecipients{
		Emails:                  []string{"some@email.test"},
		NotificationEndpointIds: nil,
	}

	output := security_rules.RuleOutput{
		Recipients:                   recipients,
		SuppressNotificationsMinutes: 60,
		Type:                         security_rules.SecurityRulesOutputTypeJson,
	}

	must := getMapFromJson("{\"match_phrase\":{\"Field\":{\"query\":\"value\"}}}")
	mustNot := getMapFromJson("{\"match_phrase\":{\"otherField\":{\"otherQuery\":\"otherValue\"}}}")

	boolFilter := security_rules.BoolFilter{Bool: security_rules.FilterLists{
		Must:    []map[string]interface{}{must},
		MustNot: []map[string]interface{}{mustNot},
	}}

	queryAllAccounts := new(bool)
	*queryAllAccounts = true

	query := security_rules.RuleQuery{
		Query:                    "type:apache_access",
		Filters:                  boolFilter,
		ShouldQueryOnAllAccounts: queryAllAccounts,
	}

	trigger := security_rules.RuleTrigger{
		Operator: security_rules.SecurityRuleOperatorEquals,
		SeverityThresholdTiers: map[string]float32{
			security_rules.SecurityRuleSeverityThresholdLow:  5,
			security_rules.SecurityRuleSeverityThresholdHigh: 10,
		},
	}

	column := security_rules.ColumnConfig{
		FieldName: "some_name",
		Sort:      security_rules.SecurityRuleOutputSortAsc,
	}

	subRuleOutput := security_rules.SubRuleOutput{Columns: []security_rules.ColumnConfig{column}}

	subComponent := security_rules.SubRule{
		QueryDefinition: query,
		Trigger:         trigger,
		Output:          subRuleOutput,
	}

	enabled := new(bool)
	*enabled = true

	createRule := security_rules.CreateUpdateSecurityRule{
		Title:                  "some_title",
		Description:            "some description",
		Tags:                   []string{"some", "tags"},
		Output:                 output,
		SearchTimeFrameMinutes: 20,
		SubComponents:          []security_rules.SubRule{subComponent},
		Enabled:                enabled,
	}

	return createRule
}

func getMapFromJson(jsonStr string) map[string]interface{} {
	var retVal map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &retVal)
	return retVal
}

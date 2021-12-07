# Security rules

Compatible with Logz.io's [security rules API](https://docs.logz.io/api/#tag/Security-rules).

Security rules help you connect the dots between your data sources and events that could indicate a security threat or breach.

**Note** that for this resource you'll need to use your **security api token**.

To create a security rule:
```go
client, err := security_rules.New(securityApiToken, apiServerAddress)
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
                Operator:              security_rules.SecurityRuleOperatorEquals,
                SeverityThresholdTiers: map[string]float32{
                security_rules.SecurityRuleSeverityThresholdLow: 5,
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
rule, err := client.CreateSecurityRule(createRule)
```

|function|func name|
|---|---|
| create security rule | `func (c *SecurityRulesClient) CreateSecurityRule(createRule CreateUpdateSecurityRule) (*SecurityRule, error)` |
| delete security rule | `func (c *SecurityRulesClient) DeleteSecurityRule(ruleId int32) (*SecurityRule, error)` |
| disable security rule | `func (c *SecurityRulesClient) DisableSecurityRule(ruleId int32) error` |
| enable security rule | `func (c *SecurityRulesClient) EnableSecurityRule(ruleId int32) error` |
| retrieve security rule | `func (c *SecurityRulesClient) RetrieveSecurityRule(ruleId int32) (*SecurityRule, error)` |
| retrieve filtered security rules | `func (c *SecurityRulesClient) RetrieveFilteredSecurityRules(filter RetrieveFiltered) (*FilteredSecurityRules, error)` |
| update security rule | `func (c *SecurityRulesClient) UpdateSecurityRule(ruleId int32, rule CreateUpdateSecurityRule) (*SecurityRule, error)` |
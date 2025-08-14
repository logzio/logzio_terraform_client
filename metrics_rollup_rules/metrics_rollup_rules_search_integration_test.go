package metrics_rollup_rules_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/metrics_rollup_rules"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationMetricsRollupRules_SearchRollupRules(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		req, err := getCreateRollupRule()
		assert.NoError(t, err)
		created, err := underTest.CreateRollupRule(req)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer underTest.DeleteRollupRule(created.Id)
			time.Sleep(2 * time.Second)

			searchReq, err := getSearchRollupRulesRequest()
			assert.NoError(t, err)
			results, err := underTest.SearchRollupRules(searchReq)
			if assert.NoError(t, err) && assert.NotNil(t, results) {
				assert.GreaterOrEqual(t, len(results), 1)
			}
		}
	}
}

func TestIntegrationMetricsRollupRules_SearchRollupRulesEmptyFilter(t *testing.T) {
	underTest, err := setupMetricsRollupRulesIntegrationTest()
	if assert.NoError(t, err) {
		searchReq := metrics_rollup_rules.SearchRollupRulesRequest{
			Filter: &metrics_rollup_rules.SearchFilter{},
			Pagination: &metrics_rollup_rules.Pagination{
				PageNumber: 0,
				PageSize:   10,
			},
		}

		results, err := underTest.SearchRollupRules(searchReq)
		if assert.NoError(t, err) && assert.NotNil(t, results) {
			assert.GreaterOrEqual(t, len(results), 0)
		}
	}
}

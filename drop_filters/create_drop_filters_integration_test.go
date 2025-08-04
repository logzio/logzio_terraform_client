package drop_filters_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationDropFilter_CreateDropFilter(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			defer underTest.DeleteDropFilter(dropFilter.Id)
			assert.NotEmpty(t, dropFilter.Id)
			assert.NotEmpty(t, dropFilter.LogType)
			assert.NotEmpty(t, dropFilter.FieldCondition)
		}
	}
}

func TestIntegrationDropFilter_CreateDropFilterWithThreshold(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilterWithThreshold()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			defer underTest.DeleteDropFilter(dropFilter.Id)
			assert.NotEmpty(t, dropFilter.Id)
			assert.NotEmpty(t, dropFilter.LogType)
			assert.NotEmpty(t, dropFilter.FieldCondition)
			assert.Equal(t, float64(10), dropFilter.ThresholdInGB)
		}
	}
}

func TestIntegrationDropFilters_CreateDropFilterNoLogType(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		createDropFilter.LogType = ""
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			defer underTest.DeleteDropFilter(dropFilter.Id)
		}
	}
}

func TestIntegrationDropFilters_CreateDropFilterNoConditions(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		createDropFilter.FieldConditions = nil
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

func TestIntegrationDropFilters_CreateDropFilterNoConditionFieldName(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		createDropFilter.FieldConditions[0].FieldName = ""
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

func TestIntegrationDropFilters_CreateDropFilterNoConditionValue(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		createDropFilter.FieldConditions[0].Value = nil
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

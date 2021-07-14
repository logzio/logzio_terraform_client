package drop_filters_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationDropFilters_ActivateDropFilter(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteDropFilter(dropFilter.Id)
			err = underTest.ActivateOrDeactivateDropFilter(dropFilter.Id, false)
			if assert.NoError(t, err) {
				time.Sleep(2 * time.Second)
				err = underTest.ActivateOrDeactivateDropFilter(dropFilter.Id, true)
				assert.NoError(t, err)
			}
		}
	}
}

func TestIntegrationDropFilters_ActivateDropFilterNotFound(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		id := "some-invalid-id"
		err = underTest.ActivateOrDeactivateDropFilter(id, true)
		assert.Error(t, err)
	}
}

func TestIntegrationDropFilters_ActivateDropFilterNoId(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.ActivateOrDeactivateDropFilter("", true)
		assert.Error(t, err)
	}
}
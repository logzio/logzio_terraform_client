package drop_filters_test

import (
	"github.com/logzio/logzio_terraform_client/drop_filters"
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
			_, err = underTest.ActivateOrDeactivateDropFilter(*dropFilter, false)
			if assert.NoError(t, err) {
				time.Sleep(2 * time.Second)
				activated, err := underTest.ActivateOrDeactivateDropFilter(*dropFilter, true)
				assert.NoError(t, err)
				assert.NotNil(t, activated)
				assert.True(t, activated.Active)
			}
		}
	}
}

func TestIntegrationDropFilters_ActivateDropFilterNotFound(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		dropFilter := drop_filters.DropFilter{
			Id: "some-invalid-id",
		}

		activate, err := underTest.ActivateOrDeactivateDropFilter(dropFilter, true)
		assert.Error(t, err)
		assert.Nil(t, activate)
	}
}

func TestIntegrationDropFilters_ActivateDropFilterNoId(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		dropFilter := drop_filters.DropFilter{}

		activate, err := underTest.ActivateOrDeactivateDropFilter(dropFilter, true)
		assert.Error(t, err)
		assert.Nil(t, activate)
	}
}
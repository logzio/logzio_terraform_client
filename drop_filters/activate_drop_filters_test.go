package drop_filters_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/drop_filters"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDropFilters_ActivateDropFilter(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("activate_drop_filter.json"))
		})

		toActivate := drop_filters.DropFilter{Id: "some-drop-filter-id"}
		dropFilter, err := underTest.ActivateOrDeactivateDropFilter(toActivate, true)
		assert.NoError(t, err)
		assert.NotNil(t, dropFilter)
		assert.NotEmpty(t, dropFilter.Id)
		assert.True(t, dropFilter.Active)
	}
}

func TestDropFilters_ActivateDropFilterAPIFailed(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("activate_drop_filter_failed.txt"))
		})

		toActivate := drop_filters.DropFilter{Id: "some-drop-filter-id"}
		dropFilter, err := underTest.ActivateOrDeactivateDropFilter(toActivate, true)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

func TestDropFilters_ActivateDropFilterNotFound(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("activate_drop_filter_not_found.txt"))
		})

		toActivate := drop_filters.DropFilter{Id: "some-drop-filter-id-not-exist"}
		dropFilter, err := underTest.ActivateOrDeactivateDropFilter(toActivate, true)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}
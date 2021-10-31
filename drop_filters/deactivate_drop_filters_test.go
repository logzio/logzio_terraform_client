package drop_filters_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDropFilters_DeactivateDropFilter(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("deactivate_drop_filter.json"))
		})

		id := "some-drop-filter-id"
		dropFilter, err := underTest.DeactivateDropFilter(id)
		assert.NoError(t, err)
		assert.NotNil(t, dropFilter)
		assert.False(t, dropFilter.Active)
	}
}

func TestDropFilters_DeactivateDropFilterAPIFailed(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("deactivate_drop_filter_failed.txt"))
		})

		id := "some-drop-filter-id"
		dropFilter, err := underTest.DeactivateDropFilter(id)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

func TestDropFilters_DeactivateDropFilterNotFound(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("deactivate_drop_filter_not_found.txt"))
		})

		id := "some-drop-filter-id-not-exist"
		dropFilter, err := underTest.DeactivateDropFilter(id)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
		assert.Contains(t, err.Error(), "failed with missing drop filter")
	}
}

package drop_filters_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropFilters_RetrieveDropFilters(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("retrieve_drop_filter.json"))
		})

		dropFilters, err := underTest.RetrieveDropFilters()
		assert.NoError(t, err)
		assert.NotNil(t, dropFilters)
		assert.NotEmpty(t, dropFilters)
		assert.Equal(t, 2, len(dropFilters))

		// Test first drop filter
		assert.Equal(t, "some-drop-filter-id", dropFilters[0].Id)
		assert.True(t, dropFilters[0].Active)
		assert.Equal(t, 10.5, dropFilters[0].ThresholdInGB)

		// Test second drop filter
		assert.Equal(t, "another-drop-filter-id", dropFilters[1].Id)
		assert.True(t, dropFilters[1].Active)
		assert.Equal(t, 5.0, dropFilters[1].ThresholdInGB)
	}
}

func TestDropFilters_RetrieveDropFiltersAPIFail(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("retrieve_drop_filter_failed.txt"))
		})

		dropFilters, err := underTest.RetrieveDropFilters()
		assert.Error(t, err)
		assert.Nil(t, dropFilters)
	}
}

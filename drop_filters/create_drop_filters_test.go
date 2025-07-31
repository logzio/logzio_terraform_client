package drop_filters_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/drop_filters"
	"github.com/stretchr/testify/assert"
)

func TestDropFilters_CreateDropFilter(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target drop_filters.CreateDropFilter
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotNil(t, target.FieldConditions)
			for _, condition := range target.FieldConditions {
				assert.NotZero(t, len(condition.FieldName))
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_drop_filter.json"))
		})

		createDropFilter := getCreateDropFilter()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.NoError(t, err)
		assert.NotNil(t, dropFilter)
		assert.True(t, dropFilter.Active)
		assert.NotEmpty(t, dropFilter.Id)
	}
}

func TestDropFilters_CreateDropFilterAPIFail(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("create_drop_filter_failed.txt"))
		})
	}

	createDropFilter := getCreateDropFilter()
	dropFilter, err := underTest.CreateDropFilter(createDropFilter)
	assert.Error(t, err)
	assert.Nil(t, dropFilter)
}

func TestDropFilters_CreateDropFilterNoCondition(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()
	assert.NoError(t, err)

	createDropFilter := getCreateDropFilter()
	createDropFilter.FieldConditions = nil
	dropFilter, err := underTest.CreateDropFilter(createDropFilter)
	assert.Error(t, err)
	assert.Nil(t, dropFilter)
}

func TestDropFilters_CreateDropFilterNoConditionFieldName(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()
	assert.NoError(t, err)

	createDropFilter := getCreateDropFilter()
	createDropFilter.FieldConditions[0].FieldName = ""
	dropFilter, err := underTest.CreateDropFilter(createDropFilter)
	assert.Error(t, err)
	assert.Nil(t, dropFilter)
}

func TestDropFilters_CreateDropFilterNoConditionValue(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()
	assert.NoError(t, err)

	createDropFilter := getCreateDropFilter()
	createDropFilter.FieldConditions[0].Value = nil
	dropFilter, err := underTest.CreateDropFilter(createDropFilter)
	assert.Error(t, err)
	assert.Nil(t, dropFilter)
}

func TestDropFilters_CreateDropFilterWithThreshold(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target drop_filters.CreateDropFilter
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotNil(t, target.FieldConditions)
			assert.Equal(t, 10.5, target.ThresholdInGB)

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_drop_filter.json"))
		})

		createDropFilter := getCreateDropFilterWithThreshold()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.NoError(t, err)
		assert.NotNil(t, dropFilter)
		assert.True(t, dropFilter.Active)
		assert.NotEmpty(t, dropFilter.Id)
		assert.Equal(t, 10.5, dropFilter.ThresholdInGB)
	}
}

func TestDropFilters_CreateDropFilterWithZeroThreshold(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target drop_filters.CreateDropFilter
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotNil(t, target.FieldConditions)
			assert.Equal(t, 0.0, target.ThresholdInGB)

			w.Header().Set("Content-Type", "application/json")
			// Create a fixture response with zero threshold
			response := `{
				"id": "some-drop-filter-id",
				"active": true,
				"logType": "some_type",
				"fieldConditions": [
					{
						"fieldName": "some_field",
						"value": "ok"
					}
				],
				"thresholdInGB": 0.0
			}`
			fmt.Fprint(w, response)
		})

		createDropFilter := getCreateDropFilterWithZeroThreshold()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.NoError(t, err)
		assert.NotNil(t, dropFilter)
		assert.True(t, dropFilter.Active)
		assert.NotEmpty(t, dropFilter.Id)
		assert.Equal(t, 0.0, dropFilter.ThresholdInGB)
	}
}

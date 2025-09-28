package unified_dashboards_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnifiedDashboards_ListDashboards(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses/api/v1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("list_dashboards.json"))
		})

		dashboards, err := underTest.ListDashboards()
		assert.NoError(t, err)
		assert.Len(t, dashboards, 2)
		assert.Equal(t, "dashboard-1", dashboards[0].Uid)
		assert.Equal(t, "System Overview", dashboards[0].Doc["title"])
		assert.Equal(t, "System monitoring dashboard", dashboards[0].Doc["description"])
		assert.Equal(t, 1, dashboards[0].Version)
		assert.Equal(t, "user1", dashboards[0].CreatedBy)
	}
}

func TestUnifiedDashboards_ListDashboardsAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses/api/v1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.ListDashboards()
		assert.Error(t, err)
	}
}

func TestUnifiedDashboards_ListDashboardsNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses/api/v1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.ListDashboards()
		assert.Error(t, err)
	}
}

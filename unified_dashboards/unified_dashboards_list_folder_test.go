package unified_dashboards_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnifiedDashboards_ListFolderDashboards(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("list_folder_dashboards.json"))
		})

		dashboards, err := underTest.ListFolderDashboards("project-1")
		assert.NoError(t, err)
		if assert.Len(t, dashboards, 2) {
			// The fixture gives uid and id different values, as the live API
			// does for any dashboard past its first revision.
			assert.Equal(t, "dashboard-1", dashboards[0].Uid)
			assert.Equal(t, "row-9f2c1a44", dashboards[0].Id)
			assert.Equal(t, "project-1", dashboards[0].ProjectId)
			assert.Equal(t, "Dashboard", dashboards[0].Doc["kind"])
			assert.Equal(t, 3, dashboards[0].Version)
			assert.False(t, dashboards[0].IsPrivate)

			assert.Equal(t, "dashboard-3", dashboards[1].Uid)
			assert.Equal(t, "row-0b7de915", dashboards[1].Id)
			assert.True(t, dashboards[1].IsPrivate)
		}
	}
}

// An unknown folder is not a 404 on this route — the gateway answers 200 with
// an empty array, the same as an existing but empty folder.
func TestUnifiedDashboards_ListFolderDashboardsEmpty(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/empty-folder/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "[]")
		})

		dashboards, err := underTest.ListFolderDashboards("empty-folder")
		assert.NoError(t, err)
		assert.Empty(t, dashboards)
	}
}

func TestUnifiedDashboards_ListFolderDashboardsAPIFail(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/project-1/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("api_error.txt"))
		})

		_, err = underTest.ListFolderDashboards("project-1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status code 500")
	}
}

func TestUnifiedDashboards_ListFolderDashboardsNotFound(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/perses-public/api/v1/projects/missing/dashboards", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("not_found.txt"))
		})

		_, err = underTest.ListFolderDashboards("missing")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing unified dashboard")
	}
}

func TestUnifiedDashboards_ListFolderDashboardsValidation(t *testing.T) {
	underTest, err, teardown := setupUnifiedDashboardsTest()
	defer teardown()

	if assert.NoError(t, err) {
		_, err = underTest.ListFolderDashboards("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "folderId must be set")
	}
}

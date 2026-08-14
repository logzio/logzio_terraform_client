package unified_projects_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_ListProjects(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-list-" + time.Now().Format("20060102150405")

		created, err := underTest.CreateProject(unified_projects.CreateProjectRequest{Name: projectName})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				if err := underTest.DeleteProject(created.Id); err != nil {
					t.Logf("cleanup: failed to delete project %s: %v", created.Id, err)
				}
			}()

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			projects, err := underTest.ListProjects(false)
			if assert.NoError(t, err) {
				found := false
				for _, item := range projects {
					if item.Project.Id == created.Id {
						found = true
						assert.Equal(t, projectName, item.Project.Name)
						break
					}
				}
				assert.True(t, found, "Created project should appear in list")
			}

			// The withDashboards variant must return real, populated entries too.
			projectsWithDashboards, err := underTest.ListProjects(true)
			if assert.NoError(t, err) {
				found := false
				for _, item := range projectsWithDashboards {
					if item.Project.Id == created.Id {
						found = true
						assert.Equal(t, projectName, item.Project.Name)
						assert.Empty(t, item.Dashboards, "freshly created project should have no dashboards")
						break
					}
				}
				assert.True(t, found, "Created project should appear in withDashboards list")
			}
		}
	}
}

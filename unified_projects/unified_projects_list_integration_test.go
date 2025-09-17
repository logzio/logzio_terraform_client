package unified_projects_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_ListProjects(t *testing.T) {
	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-list-" + time.Now().Format("20060102150405")

		// First create a project
		createReq := unified_projects.CreateProjectRequest{
			Name: projectName,
		}

		created, err := underTest.CreateProject(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				// Clean up created project
				underTest.DeleteProject(created.Id)
			}()

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// List projects and verify our project appears
			projects, err := underTest.ListProjects(false)
			if assert.NoError(t, err) {
				found := false
				for _, project := range projects {
					if project.Project.Name == projectName {
						found = true
						assert.Equal(t, created.Id, project.Project.Id)
						break
					}
				}
				assert.True(t, found, "Created project should appear in list")
			}

			// Test list with dashboards flag
			projectsWithDashboards, err := underTest.ListProjects(true)
			if assert.NoError(t, err) {
				assert.GreaterOrEqual(t, len(projectsWithDashboards), 1)
			}
		}
	}
}

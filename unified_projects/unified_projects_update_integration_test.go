package unified_projects_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_UpdateProject(t *testing.T) {
	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-update-" + time.Now().Format("20060102150405")

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

			// Update the project
			updated, err := underTest.UpdateProject(projectName, unified_projects.UpdateProjectRequest{
				DisplayName: "Updated " + projectName,
				Description: "Updated integration test description",
			})
			if assert.NoError(t, err) && assert.NotNil(t, updated) {
				assert.NotEmpty(t, updated.Id)
				assert.Equal(t, projectName, updated.Name)
			}
		}
	}
}

package unified_projects_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_GetProject(t *testing.T) {
	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-get-" + time.Now().Format("20060102150405")

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

			// Now get the project by name
			project, err := underTest.GetProject(projectName)
			if assert.NoError(t, err) && assert.NotNil(t, project) {
				assert.Equal(t, projectName, project.Name)
				assert.Equal(t, created.Id, project.Id)
				assert.NotEmpty(t, project.DisplayName)
				assert.NotNil(t, project.CreatedAt)
			}
		}
	}
}

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
			updateReq := unified_projects.Project{
				Kind: "Project",
				Metadata: unified_projects.ProjectMetadata{
					Name: projectName,
				},
				Spec: unified_projects.ProjectSpec{
					Display: unified_projects.ProjectDisplay{
						Name:        "Updated " + projectName,
						Description: "Updated integration test description",
					},
				},
			}

			updated, err := underTest.UpdateProject(projectName, updateReq)
			if assert.NoError(t, err) && assert.NotNil(t, updated) {
				assert.Equal(t, projectName, updated.Metadata.Name)
				assert.Equal(t, "Updated "+projectName, updated.Spec.Display.Name)
				assert.Equal(t, "Updated integration test description", updated.Spec.Display.Description)
				assert.NotEmpty(t, updated.Metadata.UpdatedAt)
			}
		}
	}
}

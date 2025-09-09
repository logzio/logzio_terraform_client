package unified_projects_test

import (
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_DeleteProject(t *testing.T) {
	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-delete-" + time.Now().Format("20060102150405")

		// First create a project
		createReq := unified_projects.CreateProjectRequest{
			Name: projectName,
		}

		created, err := underTest.CreateProject(createReq)
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Delete the project
			err = underTest.DeleteProject(created.Id)
			assert.NoError(t, err)

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Verify the project no longer exists by trying to get it
			_, err = underTest.GetProject(projectName)
			assert.Error(t, err, "Getting deleted project should return an error")
		}
	}
}

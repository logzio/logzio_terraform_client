package unified_projects_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_CreateProject(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-" + time.Now().Format("20060102150405")

		req := unified_projects.CreateProjectRequest{
			Name: projectName,
		}

		project, err := underTest.CreateProject(req)
		if assert.NoError(t, err) && assert.NotNil(t, project) {
			defer func() {
				// Clean up created project
				underTest.DeleteProject(project.Id)
			}()

			assert.Equal(t, projectName, project.Name)
			assert.NotEmpty(t, project.Id)
			assert.NotEmpty(t, project.CreatedAt)
		}
	}
}

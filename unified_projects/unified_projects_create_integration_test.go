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

		project, err := underTest.CreateProject(unified_projects.CreateProjectRequest{Name: projectName})
		if assert.NoError(t, err) && assert.NotNil(t, project) {
			defer func() {
				if err := underTest.DeleteProject(project.Id); err != nil {
					t.Logf("cleanup: failed to delete project %s: %v", project.Id, err)
				}
			}()

			assert.NotEmpty(t, project.Id)
			assert.Equal(t, projectName, project.Name)
			assert.Equal(t, "Project", project.Doc["kind"])
			assert.NotEmpty(t, project.CreatedAt)
		}
	}
}

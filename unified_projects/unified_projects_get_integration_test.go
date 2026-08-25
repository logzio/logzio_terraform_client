package unified_projects_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_GetProject(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-get-" + time.Now().Format("20060102150405")

		created, err := underTest.CreateProject(unified_projects.CreateProjectRequest{Name: projectName})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				if err := underTest.DeleteProject(created.Id); err != nil {
					t.Logf("cleanup: failed to delete project %s: %v", created.Id, err)
				}
			}()

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// The API addresses projects by id.
			project, err := underTest.GetProject(created.Id)
			if assert.NoError(t, err) && assert.NotNil(t, project) {
				assert.Equal(t, created.Id, project.Id)
				assert.Equal(t, projectName, project.Name)
				assert.Equal(t, "Project", project.Doc["kind"])
			}
		}
	}
}

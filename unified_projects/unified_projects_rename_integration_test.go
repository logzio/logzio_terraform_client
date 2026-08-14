package unified_projects_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_RenameProject(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-rename-" + time.Now().Format("20060102150405")

		created, err := underTest.CreateProject(unified_projects.CreateProjectRequest{Name: projectName})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				if err := underTest.DeleteProject(created.Id); err != nil {
					t.Logf("cleanup: failed to delete project %s: %v", created.Id, err)
				}
			}()

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			renamed, err := underTest.RenameProject(created.Id, projectName+"-renamed")
			if assert.NoError(t, err) && assert.NotNil(t, renamed) {
				assert.Equal(t, created.Id, renamed.Id)
				t.Logf("rename response name: %q (requested %q)", renamed.Name, projectName+"-renamed")
			}

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			retrieved, err := underTest.GetProject(created.Id)
			if assert.NoError(t, err) && assert.NotNil(t, retrieved) {
				t.Logf("post-rename project name: %q", retrieved.Name)
			}
		}
	}
}

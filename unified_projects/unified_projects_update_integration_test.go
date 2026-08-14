package unified_projects_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_UpdateProject(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-update-" + time.Now().Format("20060102150405")

		created, err := underTest.CreateProject(unified_projects.CreateProjectRequest{Name: projectName})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				if err := underTest.DeleteProject(created.Id); err != nil {
					t.Logf("cleanup: failed to delete project %s: %v", created.Id, err)
				}
			}()

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			updated, err := underTest.UpdateProject(created.Id, unified_projects.UpdateProjectRequest{
				Name:        projectName,
				DisplayName: "Updated " + projectName,
				Description: "Updated integration test description",
			})
			if assert.NoError(t, err) && assert.NotNil(t, updated) {
				assert.Equal(t, created.Id, updated.Id)
				assert.Equal(t, "Updated "+projectName, updated.Name)
			}

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Verify the update persisted.
			retrieved, err := underTest.GetProject(created.Id)
			if assert.NoError(t, err) && assert.NotNil(t, retrieved) {
				assert.Equal(t, "Updated "+projectName, retrieved.Name)
			}
		}
	}
}

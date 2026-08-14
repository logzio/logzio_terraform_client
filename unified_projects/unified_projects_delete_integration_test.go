package unified_projects_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_DeleteProject(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		projectName := "tf-client-it-delete-" + time.Now().Format("20060102150405")

		created, err := underTest.CreateProject(unified_projects.CreateProjectRequest{Name: projectName})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			time.Sleep(2 * time.Second) // Allow for eventual consistency

			err = underTest.DeleteProject(created.Id)
			if !assert.NoError(t, err) {
				// Deletion failed — clean up on the way out so the account stays tidy.
				if err := underTest.DeleteProject(created.Id); err != nil {
					t.Logf("cleanup: failed to delete project %s: %v", created.Id, err)
				}
				return
			}

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// Verify the project no longer exists; the error must be the
			// not-found classification, not some other failure.
			_, err = underTest.GetProject(created.Id)
			if assert.Error(t, err, "Getting deleted project should return an error") {
				assert.Contains(t, err.Error(), "failed with missing unified project")
			}
		}
	}
}

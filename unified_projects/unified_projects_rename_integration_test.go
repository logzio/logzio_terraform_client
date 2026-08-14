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
			defer underTest.DeleteProject(created.Id)

			time.Sleep(2 * time.Second)

			renamed, err := underTest.RenameProject(created.Id, projectName+"-renamed")
			if assert.NoError(t, err) && assert.NotNil(t, renamed) {
				assert.Equal(t, projectName+"-renamed", renamed.Name)
			}
		}
	}
}

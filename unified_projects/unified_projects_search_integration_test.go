package unified_projects_test

import (
	"os"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedProjects_SearchProjects(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	underTest, err := setupUnifiedProjectsIntegrationTest()
	if assert.NoError(t, err) {
		uniqueId := time.Now().Format("20060102150405")
		projectName := "tf-client-it-search-" + uniqueId

		// First create a project with a unique name
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

			// Search for projects using part of the unique name
			results, err := underTest.SearchProjects(unified_projects.SearchProjectsRequest{
				Query: uniqueId,
				Limit: 10,
				Page:  1,
			})
			if assert.NoError(t, err) {
				found := false
				for _, p := range results {
					if p.Name == projectName {
						found = true
						assert.Equal(t, created.Id, p.Id)
						break
					}
				}
				assert.True(t, found, "Created project should appear in search results")
			}
		}
	}
}

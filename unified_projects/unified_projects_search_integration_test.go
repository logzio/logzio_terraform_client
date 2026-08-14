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

		created, err := underTest.CreateProject(unified_projects.CreateProjectRequest{Name: projectName})
		if assert.NoError(t, err) && assert.NotNil(t, created) {
			defer func() {
				if err := underTest.DeleteProject(created.Id); err != nil {
					t.Logf("cleanup: failed to delete project %s: %v", created.Id, err)
				}
			}()

			time.Sleep(2 * time.Second) // Allow for eventual consistency

			// The search endpoint behaves as a paginated listing; page through
			// until the created project shows up (bounded to keep CI predictable).
			found := false
			for page := 1; page <= 20 && !found; page++ {
				resp, err := underTest.SearchProjects(unified_projects.SearchProjectsRequest{
					Query: projectName,
					Limit: 100,
					Page:  page,
				})
				if !assert.NoError(t, err) || !assert.NotNil(t, resp) {
					return
				}
				if len(resp.Results) == 0 {
					break
				}
				for _, item := range resp.Results {
					if item.Project.Id == created.Id {
						found = true
						assert.Equal(t, projectName, item.Project.Name)
						break
					}
				}
			}
			assert.True(t, found, "Created project should appear in search results")
		}
	}
}

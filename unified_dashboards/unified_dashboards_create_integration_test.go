package unified_dashboards_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/logzio/logzio_terraform_client/unified_dashboards"
	"github.com/logzio/logzio_terraform_client/unified_projects"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUnifiedDashboards_CreateDashboard(t *testing.T) {
	projClient, err := setupUnifiedProjectsIntegrationTest()
	if !assert.NoError(t, err) || projClient == nil {
		return
	}
	dashClient, err := setupUnifiedDashboardsIntegrationTest()
	if !assert.NoError(t, err) || dashClient == nil {
		return
	}

	uniqueId := time.Now().Format("20060102150405")
	projName := "tf-client-it-dash-" + uniqueId

	// First create a project
	proj, err := projClient.CreateProject(unified_projects.CreateProjectRequest{Name: projName})
	if assert.NoError(t, err) && assert.NotNil(t, proj) {
		defer projClient.DeleteProject(proj.Id)

		// Debug: Print project details
		fmt.Printf("Created project: ID=%s, Name=%s, DisplayName=%s\n", proj.Id, proj.Name, proj.DisplayName)

		// Wait for eventual consistency
		fmt.Printf("Waiting 5 seconds for eventual consistency...\n")
		time.Sleep(5 * time.Second)

		// Verify project exists by trying to get it multiple ways
		fmt.Printf("Verifying project existence...\n")

		// Try to get by name
		retrievedProj, getErr := projClient.GetProject(proj.Name)
		if getErr != nil {
			fmt.Printf("Failed to retrieve project by name: %v\n", getErr)
		} else {
			fmt.Printf("Retrieved project by name: ID=%s, Name=%s, DisplayName=%s\n", retrievedProj.Id, retrievedProj.Name, retrievedProj.DisplayName)
		}

		// Try to list projects and find ours
		fmt.Printf("Checking if project appears in list...\n")
		projects, listErr := projClient.ListProjects(false)
		if listErr != nil {
			fmt.Printf("Failed to list projects: %v\n", listErr)
		} else {
			found := false
			fmt.Printf("Found %d total projects\n", len(projects))
			for _, p := range projects {
				if p.Project.Id == proj.Id {
					found = true
					fmt.Printf("Found our project in list: ID=%s, Name=%s\n", p.Project.Id, p.Project.Name)
					break
				}
			}
			if !found {
				fmt.Printf("Our project (ID=%s) was NOT found in the project list!\n", proj.Id)
			}
		}

		// Additional wait if project verification failed
		if getErr != nil {
			fmt.Printf("Project verification failed, waiting additional 10 seconds...\n")
			time.Sleep(10 * time.Second)
		}

		// Create a dashboard with retry logic
		createReq := unified_dashboards.CreateDashboardRequest{
			Doc: map[string]interface{}{
				"title":       "IT Dashboard " + uniqueId,
				"description": "Integration test dashboard",
				"panels":      []interface{}{},
				"refresh":     "1m",
			},
		}

		// Try dashboard creation with retry
		var created *unified_dashboards.Dashboard
		var dashErr error
		maxRetries := 5

		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Printf("Attempt %d: Creating dashboard in project ID: %s\n", attempt, proj.Id)
			fmt.Printf("  Using URL pattern: /perses/api/v1/projects/%s/dashboards\n", proj.Id)

			created, dashErr = dashClient.CreateDashboard(proj.Id, createReq)

			if dashErr == nil {
				fmt.Printf("Dashboard created successfully on attempt %d\n", attempt)
				break
			}

			fmt.Printf("Attempt %d failed: %v\n", attempt, dashErr)
			if attempt < maxRetries {
				fmt.Printf("Waiting 5 seconds before retry...\n")
				time.Sleep(5 * time.Second)
			}
		}

		if assert.NoError(t, dashErr) && assert.NotNil(t, created) {
			defer dashClient.DeleteDashboard(proj.Id, created.Uid)

			assert.NotEmpty(t, created.Uid)
			assert.NotNil(t, created.Doc)
			assert.Equal(t, "IT Dashboard "+uniqueId, created.Doc["title"])
			assert.Equal(t, "Integration test dashboard", created.Doc["description"])
			assert.NotEmpty(t, created.CreatedAt)
		} else {
			fmt.Printf("Final dashboard creation failed. This suggests the project ID format or timing is incorrect.\n")
		}
	}
}

// Helper function to setup unified projects integration test
func setupUnifiedProjectsIntegrationTest() (*unified_projects.ProjectsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := unified_projects.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

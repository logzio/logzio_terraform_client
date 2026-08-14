package unified_projects

import (
	"encoding/json"
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	projectsServiceEndpoint = "%s/perses-public/api/v1/projects"
	projectsByIdEndpoint    = "%s/perses-public/api/v1/projects/%s"
	projectsSearchEndpoint  = "%s/perses-public/api/v1/projects/search"
	projectsRenameEndpoint  = "%s/perses-public/api/v1/projects/%s/rename"

	projectResourceName = "unified project"

	createProjectOperation  = "CreateUnifiedProject"
	getProjectOperation     = "GetUnifiedProject"
	listProjectsOperation   = "ListUnifiedProjects"
	updateProjectOperation  = "UpdateUnifiedProject"
	searchProjectsOperation = "SearchUnifiedProjects"
	deleteProjectOperation  = "DeleteUnifiedProject"
	renameProjectOperation  = "RenameUnifiedProject"
)

type ProjectsClient struct {
	*client.Client
}

// Request types
// CreateProjectRequest is not marshalled directly — the client translates it
// into the Perses Project envelope the API requires.
type CreateProjectRequest struct {
	Name        string `json:"-"` // the project's identity (metadata.name); required
	DisplayName string `json:"-"` // optional; defaults to Name
	Description string `json:"-"` // optional
}

// UpdateProjectRequest replaces the project's Perses document on PUT: fields left
// empty are cleared on the server, so always send the full desired state.
type UpdateProjectRequest struct {
	Name        string `json:"-"` // metadata.name — the project's identity; required by the API (see ProjectSummary.MetadataName)
	DisplayName string `json:"-"` // required — the PUT replaces the display block entirely
	Description string `json:"-"` // optional; empty clears any existing description
}

// SearchProjectsRequest is the POST body for the projects search endpoint.
type SearchProjectsRequest struct {
	Query string `json:"query,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Page  int    `json:"page,omitempty"`
}

// projectEnvelope is the Perses-style document the API expects as the create/update body.
type projectEnvelope struct {
	Kind     string          `json:"kind"`
	Metadata projectMetadata `json:"metadata"`
	Spec     projectSpec     `json:"spec"`
}

type projectMetadata struct {
	Name string `json:"name"`
}

type projectSpec struct {
	Display projectDisplay `json:"display"`
}

type projectDisplay struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func newProjectEnvelope(name, displayName, description string) projectEnvelope {
	return projectEnvelope{
		Kind:     "Project",
		Metadata: projectMetadata{Name: name},
		Spec:     projectSpec{Display: projectDisplay{Name: displayName, Description: description}},
	}
}

// Response types
type ProjectSummary struct {
	Id        string                 `json:"id"`
	Name      string                 `json:"name,omitempty"` // the project's display name
	Doc       map[string]interface{} `json:"doc,omitempty"`  // the Perses Project document (kind/metadata/spec)
	CreatedAt string                 `json:"createdAt,omitempty"`
	UpdatedAt string                 `json:"updatedAt,omitempty"`
}

// MetadataName returns the project's Perses identity (doc.metadata.name),
// which is distinct from the display name carried in the Name field. Use it
// as UpdateProjectRequest.Name in read-modify-write flows; passing the
// display name there would silently rewrite the project's identity.
func (p *ProjectSummary) MetadataName() string {
	metadata, ok := p.Doc["metadata"].(map[string]interface{})
	if !ok {
		return ""
	}
	name, _ := metadata["name"].(string)
	return name
}

// ProjectDashboard is the dashboard payload embedded in project list/search
// responses. Id is the dashboard's addressable identifier.
type ProjectDashboard struct {
	Id        string                 `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	ProjectId string                 `json:"projectId,omitempty"`
	Doc       map[string]interface{} `json:"doc,omitempty"`
}

// ProjectListItem is one entry of the list and search responses: the project
// plus the dashboards it contains.
type ProjectListItem struct {
	Project    ProjectSummary     `json:"project"`
	Dashboards []ProjectDashboard `json:"dashboards,omitempty"`
}

type SearchProjectsResponse struct {
	Results    []ProjectListItem      `json:"results"`
	Total      int                    `json:"total,omitempty"`
	Pagination map[string]interface{} `json:"pagination,omitempty"`
}

func New(apiToken, baseUrl string) (*ProjectsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	return &ProjectsClient{
		Client: client.New(apiToken, baseUrl),
	}, nil
}

// Validation helpers
func validateCreateProjectRequest(req CreateProjectRequest) error {
	if len(req.Name) == 0 {
		return fmt.Errorf("name must be set")
	}
	return nil
}

func validateGetProjectRequest(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("id must be set")
	}
	return nil
}

func validateUpdateProjectRequest(id string, req UpdateProjectRequest) error {
	if len(id) == 0 {
		return fmt.Errorf("id must be set")
	}
	if len(req.Name) == 0 {
		return fmt.Errorf("name must be set")
	}
	if len(req.DisplayName) == 0 {
		return fmt.Errorf("displayName must be set")
	}
	return nil
}

func validateDeleteProjectRequest(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("id must be set")
	}
	return nil
}

func validateRenameProjectRequest(id, newName string) error {
	if len(id) == 0 {
		return fmt.Errorf("id must be set")
	}
	if len(newName) == 0 {
		return fmt.Errorf("newName must be set")
	}
	return nil
}

// unmarshalProject decodes an API response into a ProjectSummary and guards
// against a silently mismatched wire shape: encoding/json ignores unknown
// fields, so a wrong shape would otherwise yield a zero-valued summary with
// no error.
func unmarshalProject(operation string, res []byte) (*ProjectSummary, error) {
	var result ProjectSummary
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal response: %w (body: %.200s)", operation, err, res)
	}
	if len(result.Id) == 0 {
		return nil, fmt.Errorf("%s succeeded but the response contained no project id (body: %.200s)", operation, res)
	}
	return &result, nil
}

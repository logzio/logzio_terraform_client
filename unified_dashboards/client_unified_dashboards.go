package unified_dashboards

import (
	"encoding/json"
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	dashboardsListEndpoint   = "%s/perses-public/api/v1/dashboards"
	dashboardByUidEndpoint   = "%s/perses-public/api/v1/projects/%s/dashboards/%s"
	folderDashboardsEndpoint = "%s/perses-public/api/v1/projects/%s/dashboards"
	dashboardsMoveEndpoint   = "%s/perses-public/api/v1/dashboards/move"
	dashboardsSearchEndpoint = "%s/perses-public/api/v1/dashboards/search"

	dashboardResourceName = "unified dashboard"

	createDashboardOperation      = "CreateUnifiedDashboard"
	getDashboardOperation         = "GetUnifiedDashboard"
	listDashboardsOperation       = "ListUnifiedDashboards"
	listFolderDashboardsOperation = "ListUnifiedFolderDashboards"
	updateDashboardOperation      = "UpdateUnifiedDashboard"
	moveDashboardOperation        = "MoveUnifiedDashboard"
	searchDashboardsOperation     = "SearchUnifiedDashboards"
	deleteDashboardOperation      = "DeleteUnifiedDashboard"
)

type DashboardsClient struct {
	*client.Client
}

// Request types.
//
// Doc is the full Perses Dashboard document, e.g.
//
//	{"kind": "Dashboard", "metadata": {"name": "..."}, "spec": {"display": {"name": "..."}, "duration": "1h", "panels": {}, "layouts": []}}
//
// The API requires the Perses envelope (kind/metadata/spec); a bare
// {"title": ...} document is rejected with "Dashboard name is required".
type CreateDashboardRequest struct {
	Doc map[string]any `json:"doc"`
}

type UpdateDashboardRequest struct {
	Doc map[string]any `json:"doc"`
}

// SearchDashboardsRequest is the POST body for the dashboards search endpoint.
// Both fields are optional: an empty request returns every dashboard in the
// account, one page at a time.
type SearchDashboardsRequest struct {
	Filter     *SearchDashboardsFilter     `json:"filter,omitempty"`
	Pagination *SearchDashboardsPagination `json:"pagination,omitempty"`
}

type SearchDashboardsFilter struct {
	// SearchTerm matches against the Perses document's metadata.name — the
	// dashboard's identity — and nothing else. Probed 2026-08-20 with a
	// dashboard whose display name, panel name, datasource reference and
	// PromQL query all contained a distinctive token: searching for any of
	// them returned zero results, while its metadata.name matched. This is
	// not a full-text search over the document.
	SearchTerm string `json:"searchTerm,omitempty"`
	// CreatedBy mirrors unified_projects.SearchProjectsFilter, but the gateway
	// currently ignores it — a bogus account id still returns every dashboard
	// (probed 2026-08-20). Sent for forward compatibility; do not rely on it
	// to filter.
	CreatedBy []int64 `json:"createdBy,omitempty"`
}

// SearchDashboardsPagination is echoed back on the response. The server
// defaults to page 1 with a page size of 20.
type SearchDashboardsPagination struct {
	PageNumber int `json:"pageNumber,omitempty"`
	PageSize   int `json:"pageSize,omitempty"`
}

// MoveDashboardRequest relocates a dashboard between folders. All three
// fields are required by the API (note: the published docs describe a
// different shape; this one is verified against the live gateway).
type MoveDashboardRequest struct {
	DashboardId  string `json:"dashboardId"`  // the dashboard's stable Uid, not its version-row Id
	OldProjectId string `json:"oldProjectId"` // current folder id
	NewProjectId string `json:"newProjectId"` // destination folder id
}

// Response types

// Dashboard carries two identifiers that are not interchangeable:
//
//   - Uid is the stable handle. It survives updates, and every folder-scoped
//     route addresses a dashboard by it (get/update/delete/move all take a uid).
//   - Id identifies the version row behind the current revision, so it changes
//     on every update. No method in this client accepts it; it is mapped for
//     completeness only.
//
// Persist Uid, never Id.
type Dashboard struct {
	Id        string         `json:"id,omitempty"` // version-row id — changes on every update
	Uid       string         `json:"uid"`          // stable identifier — address dashboards by this
	Name      string         `json:"name,omitempty"`
	ProjectId string         `json:"projectId,omitempty"`
	Doc       map[string]any `json:"doc"`
	Version   int            `json:"version,omitempty"`
	CreatedAt string         `json:"createdAt,omitempty"`
	UpdatedAt string         `json:"updatedAt,omitempty"`
	IsPrivate bool           `json:"isPrivate"`
}

// MoveDashboardResponse acknowledges a completed move. It carries the moved
// dashboard's *version-row* Id — not its uid, and not a new identifier.
//
// This is easy to get wrong, because the server assigns a brand-new dashboard
// the same value for id, uid and name; the three only diverge once the
// dashboard has been updated at least once. Probed 2026-08-20 against the live
// gateway with a version-2 dashboard: the response is the current row id and
// differs from the uid. Keep addressing the dashboard by the uid you passed as
// MoveDashboardRequest.DashboardId; this Id is an acknowledgement, not a handle.
type MoveDashboardResponse struct {
	Id string `json:"id"`
}

// SearchDashboardsResponse is the dashboards search result set. Total is the
// number of matches before pagination, so it can be larger than len(Results);
// Pagination echoes back the page the server actually served.
type SearchDashboardsResponse struct {
	Results    []Dashboard                 `json:"results"`
	Total      int                         `json:"total,omitempty"`
	Pagination *SearchDashboardsPagination `json:"pagination,omitempty"`
}

func New(apiToken, baseUrl string) (*DashboardsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("base URL not defined")
	}
	return &DashboardsClient{
		Client: client.New(apiToken, baseUrl),
	}, nil
}

// unmarshalDashboard decodes an API response into a Dashboard and guards
// against a silently mismatched wire shape: encoding/json ignores unknown
// fields, so a wrong shape would otherwise yield a zero-valued Dashboard
// with no error.
func unmarshalDashboard(operation string, res []byte) (*Dashboard, error) {
	var result Dashboard
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal response: %w (body: %.200s)", operation, err, res)
	}
	if len(result.Uid) == 0 {
		return nil, fmt.Errorf("%s succeeded but the response contained no dashboard uid (body: %.200s)", operation, res)
	}
	return &result, nil
}

// Validation helpers
func validateCreateDashboardRequest(folderId string, req CreateDashboardRequest) error {
	if len(folderId) == 0 {
		return fmt.Errorf("folderId must be set")
	}
	if len(req.Doc) == 0 {
		return fmt.Errorf("doc must be set")
	}
	return nil
}

func validateUpdateDashboardRequest(folderId, uid string, req UpdateDashboardRequest) error {
	if len(folderId) == 0 {
		return fmt.Errorf("folderId must be set")
	}
	if len(uid) == 0 {
		return fmt.Errorf("uid must be set")
	}
	if len(req.Doc) == 0 {
		return fmt.Errorf("doc must be set")
	}
	return nil
}

func validateListFolderDashboardsRequest(folderId string) error {
	if len(folderId) == 0 {
		return fmt.Errorf("folderId must be set")
	}
	return nil
}

func validateGetDashboardRequest(folderId, uid string) error {
	if len(folderId) == 0 {
		return fmt.Errorf("folderId must be set")
	}
	if len(uid) == 0 {
		return fmt.Errorf("uid must be set")
	}
	return nil
}

func validateDeleteDashboardRequest(folderId, uid string) error {
	if len(folderId) == 0 {
		return fmt.Errorf("folderId must be set")
	}
	if len(uid) == 0 {
		return fmt.Errorf("uid must be set")
	}
	return nil
}

func validateMoveDashboardRequest(req MoveDashboardRequest) error {
	if len(req.DashboardId) == 0 {
		return fmt.Errorf("dashboardId must be set")
	}
	if len(req.OldProjectId) == 0 {
		return fmt.Errorf("oldProjectId must be set")
	}
	if len(req.NewProjectId) == 0 {
		return fmt.Errorf("newProjectId must be set")
	}
	return nil
}

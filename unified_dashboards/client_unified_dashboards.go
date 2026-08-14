package unified_dashboards

import (
	"encoding/json"
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	dashboardsListEndpoint   = "%s/perses-public/api/v1/dashboards"
	dashboardByUidEndpoint   = "%s/perses-public/api/v1/projects/%s/dashboards/%s"
	dashboardsCreateEndpoint = "%s/perses-public/api/v1/projects/%s/dashboards"
	dashboardsMoveEndpoint   = "%s/perses-public/api/v1/dashboards/move"

	dashboardResourceName = "unified dashboard"

	createDashboardOperation = "CreateUnifiedDashboard"
	getDashboardOperation    = "GetUnifiedDashboard"
	listDashboardsOperation  = "ListUnifiedDashboards"
	updateDashboardOperation = "UpdateUnifiedDashboard"
	moveDashboardOperation   = "MoveUnifiedDashboard"
	deleteDashboardOperation = "DeleteUnifiedDashboard"
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
	Doc map[string]interface{} `json:"doc"`
}

type UpdateDashboardRequest struct {
	Doc map[string]interface{} `json:"doc"`
}

type MoveDashboardRequest struct {
	Uid            string `json:"uid"`
	TargetFolderId string `json:"targetFolderId"`
}

// Response types
type Dashboard struct {
	Id        string                 `json:"id,omitempty"`
	Uid       string                 `json:"uid"`
	Name      string                 `json:"name,omitempty"`
	ProjectId string                 `json:"projectId,omitempty"`
	Doc       map[string]interface{} `json:"doc"`
	Version   int                    `json:"version,omitempty"`
	CreatedAt string                 `json:"createdAt,omitempty"`
	UpdatedAt string                 `json:"updatedAt,omitempty"`
	IsPrivate bool                   `json:"isPrivate"`
}

type MoveDashboardResponse struct {
	Uid string `json:"uid"`
}

func New(apiToken, baseUrl string) (*DashboardsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
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
	if len(req.Uid) == 0 {
		return fmt.Errorf("uid must be set")
	}
	if len(req.TargetFolderId) == 0 {
		return fmt.Errorf("targetFolderId must be set")
	}
	return nil
}

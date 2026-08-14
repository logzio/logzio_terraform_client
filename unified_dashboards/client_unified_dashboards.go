package unified_dashboards

import (
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

// Request types
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
	Uid       string                 `json:"uid"`
	Doc       map[string]interface{} `json:"doc"`
	Version   int                    `json:"version,omitempty"`
	CreatedAt string                 `json:"createdAt,omitempty"`
	UpdatedAt string                 `json:"updatedAt,omitempty"`
	CreatedBy string                 `json:"createdBy,omitempty"`
	UpdatedBy string                 `json:"updatedBy,omitempty"`
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

// Validation helpers
func validateCreateDashboardRequest(folderId string, req CreateDashboardRequest) error {
	if len(folderId) == 0 {
		return fmt.Errorf("folderId must be set")
	}
	if req.Doc == nil || len(req.Doc) == 0 {
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
	if req.Doc == nil || len(req.Doc) == 0 {
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

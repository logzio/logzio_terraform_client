package unified_dashboards

import (
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	dashboardsListEndpoint   = "%s/perses/api/v1/dashboards"
	dashboardByUidEndpoint   = "%s/perses/api/v1/projects/%s/dashboards/%s"
	dashboardsCreateEndpoint = "%s/perses/api/v1/projects/%s/dashboards"
	dashboardsSearchEndpoint = "%s/perses/api/v1/dashboards/search"

	dashboardResourceName = "unified dashboard"

	createDashboardOperation  = "CreateUnifiedDashboard"
	getDashboardOperation     = "GetUnifiedDashboard"
	listDashboardsOperation   = "ListUnifiedDashboards"
	updateDashboardOperation  = "UpdateUnifiedDashboard"
	searchDashboardsOperation = "SearchUnifiedDashboards"
	deleteDashboardOperation  = "DeleteUnifiedDashboard"
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

type SearchDashboardsRequest struct {
	Query   string   `json:"query,omitempty"`
	Tag     []string `json:"tag,omitempty"`
	Starred *bool    `json:"starred,omitempty"`
	Limit   int      `json:"limit,omitempty"`
	Page    int      `json:"page,omitempty"`
	Sort    string   `json:"sort,omitempty"`
	Order   string   `json:"order,omitempty"`
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

func validateSearchDashboardsRequest(req SearchDashboardsRequest) error {
	return nil
}

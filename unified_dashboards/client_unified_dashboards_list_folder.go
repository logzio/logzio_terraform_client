package unified_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	listFolderDashboardsMethod   = http.MethodGet
	listFolderDashboardsSuccess  = http.StatusOK
	listFolderDashboardsNotFound = http.StatusNotFound
)

// ListFolderDashboards returns the dashboards contained in a single folder,
// which is cheaper than filtering the account-wide ListDashboards by ProjectId.
//
// The gateway does not 404 for an unknown folder: an id that matches no folder
// returns 200 with an empty array, exactly like an existing but empty folder
// (probed 2026-08-20). Callers that need to tell the two apart must check the
// folder itself via unified_projects.GetProject.
func (c *DashboardsClient) ListFolderDashboards(folderId string) ([]Dashboard, error) {
	if err := validateListFolderDashboardsRequest(folderId); err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listFolderDashboardsMethod,
		Url:          fmt.Sprintf(folderDashboardsEndpoint, c.BaseUrl, folderId),
		Body:         nil,
		SuccessCodes: []int{listFolderDashboardsSuccess},
		NotFoundCode: listFolderDashboardsNotFound,
		ResourceId:   folderId,
		ApiAction:    listFolderDashboardsOperation,
		ResourceName: dashboardResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result []Dashboard
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal response: %w (body: %.200s)", listFolderDashboardsOperation, err, res)
	}

	return result, nil
}

package unified_dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	getDashboardMethod   = http.MethodGet
	getDashboardSuccess  = http.StatusOK
	getDashboardNotFound = http.StatusNotFound
)

func (c *DashboardsClient) GetDashboard(folderId, uid string, source *string) (*Dashboard, error) {
	if err := validateGetDashboardRequest(folderId, uid); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(dashboardByUidEndpoint, c.BaseUrl, folderId, uid)

	if source != nil && len(*source) > 0 {
		u, err := url.Parse(endpoint)
		if err != nil {
			return nil, err
		}
		q := u.Query()
		q.Set("source", *source)
		u.RawQuery = q.Encode()
		endpoint = u.String()
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getDashboardMethod,
		Url:          endpoint,
		Body:         nil,
		SuccessCodes: []int{getDashboardSuccess},
		NotFoundCode: getDashboardNotFound,
		ResourceId:   uid,
		ApiAction:    getDashboardOperation,
		ResourceName: dashboardResourceName,
	})
	if err != nil {
		return nil, err
	}

	var result Dashboard
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

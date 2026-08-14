package unified_dashboards_test

import (
	"os"
	"testing"
)

func TestIntegrationUnifiedDashboards_MoveDashboard(t *testing.T) {
	if os.Getenv("LOGZIO_API_TOKEN") == "" {
		t.Skip("LOGZIO_API_TOKEN not set")
	}

	// POST /perses-public/api/v1/dashboards/move is documented at
	// api-docs.logz.io but was not deployed on the api.logz.io gateway as of
	// 2026-08-14 (the route 404s). Re-enable this test once the endpoint is
	// live; MoveDashboard itself is covered by unit tests.
	t.Skip("dashboards/move endpoint not deployed on the public gateway (verified 2026-08-14)")
}

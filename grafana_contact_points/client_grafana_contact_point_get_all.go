package grafana_contact_points

import "net/http"

const (
	getAllGrafanaContactPointServiceUrl      = grafanaContactPointServiceEndpoint
	getAllGrafanaContactPointServiceMethod   = http.MethodGet
	getAllGrafanaContactPointServiceSuccess  = http.StatusOK
	getAllGrafanaContactPointServiceNotFound = http.StatusNotFound
)

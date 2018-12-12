package logzio_client

import "net/http"

const getEndpointsServiceUrl string = "%s/v1/endpoints/%d"
const getEndpointsServiceMethod string = http.MethodGet
const getEndpointsMethodSuccess int = 200
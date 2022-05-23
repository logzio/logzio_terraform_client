package kibana_objects

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	exportKibanaObjectServiceUrl            = kibanaObjectsServiceEndpoint + "/export"
	exportKibanaObjectServiceMethod         = http.MethodPost
	exportKibanaObjectServiceSuccess        = http.StatusOK
	exportKibanaObjectServiceCreated        = http.StatusCreated
	exportKibanaObjectServiceNoContent      = http.StatusNoContent
	exportKibanaObjectServiceStatusNotFound = http.StatusNotFound
)

// ExportKibanaObject allows export of the Kibana objects configuration.
// https://docs.logz.io/api/#operation/exportSavedObjects
func (c *KibanaObjectsClient) ExportKibanaObject(exportRequest KibanaObjectExportRequest) (*KibanaObjectExportResponse, error) {
	exportReqJson, err := json.Marshal(exportRequest)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   exportKibanaObjectServiceMethod,
		Url:          fmt.Sprintf(exportKibanaObjectServiceUrl, c.BaseUrl),
		Body:         exportReqJson,
		SuccessCodes: []int{exportKibanaObjectServiceSuccess, exportKibanaObjectServiceCreated, exportKibanaObjectServiceNoContent},
		NotFoundCode: exportKibanaObjectServiceStatusNotFound,
		ResourceId:   nil,
		ApiAction:    exportKibanaObjectOperation,
		ResourceName: kibanaObjectResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal KibanaObjectExportResponse
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

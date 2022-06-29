package kibana_objects

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	importKibanaObjectServiceUrl            = kibanaObjectsServiceEndpoint + "/import"
	importKibanaObjectServiceMethod         = http.MethodPost
	importKibanaObjectServiceSuccess        = http.StatusOK
	importKibanaObjectServiceCreated        = http.StatusCreated
	importKibanaObjectServiceNoContent      = http.StatusNoContent
	importKibanaObjectServiceStatusNotFound = http.StatusNotFound
)

// ImportKibanaObject allows import of the Kibana objects configuration.
// https://docs.logz.io/api/#operation/exportSavedObjects
func (c *KibanaObjectsClient) ImportKibanaObject(importRequest KibanaObjectImportRequest) (*KibanaObjectImportResponse, error) {
	importReqJson, err := json.Marshal(importRequest)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   importKibanaObjectServiceMethod,
		Url:          fmt.Sprintf(importKibanaObjectServiceUrl, c.BaseUrl),
		Body:         importReqJson,
		SuccessCodes: []int{importKibanaObjectServiceSuccess, importKibanaObjectServiceCreated, importKibanaObjectServiceNoContent},
		NotFoundCode: importKibanaObjectServiceStatusNotFound,
		ResourceId:   nil,
		ApiAction:    importKibanaObjectOperation,
		ResourceName: kibanaObjectResourceName,
	})

	if err != nil {
		return nil, err
	}

	var retVal KibanaObjectImportResponse
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

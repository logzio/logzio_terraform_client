package restore_logs

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getRestoreLogsServiceUrl                = restoreLogsServiceEndpoint + "/%d"
	getRestoreLogsServiceMethod      string = http.MethodGet
	getRestoreLogsMethodSuccess      int    = http.StatusOK
	getRestoreLogsLogsMethodNotFound int    = http.StatusNotFound
)

// GetRestoreOperation gets a restore operation id and returns its status or an error
func (c *RestoreClient) GetRestoreOperation(restoreId int32) (*RestoreOperation, error) {
	//req, err := c.buildGetApiRequest(c.ApiToken, restoreId)
	//if err != nil {
	//	return nil, err
	//}
	//httpClient := client.GetHttpClient(req)
	//resp, err := httpClient.Do(req)
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//
	//jsonBytes, _ := ioutil.ReadAll(resp.Body)
	//if !logzio_client.CheckValidStatus(resp, []int{getRestoreLogsMethodSuccess}) {
	//	if resp.StatusCode == getRestoreLogsLogsMethodNotFound {
	//		return nil, fmt.Errorf("API call %s failed with missing restore operation %d, data: %s", getRestoreOperation, restoreId, jsonBytes)
	//	}
	//
	//	return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", getRestoreOperation, resp.StatusCode, jsonBytes)
	//}
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getRestoreLogsServiceMethod,
		Url:          fmt.Sprintf(getRestoreLogsServiceUrl, c.BaseUrl, restoreId),
		Body:         nil,
		SuccessCodes: []int{getRestoreLogsMethodSuccess},
		NotFoundCode: getRestoreLogsLogsMethodNotFound,
		ResourceId:   restoreId,
		ApiAction:    getRestoreOperation,
	})

	if err != nil {
		return nil, err
	}

	var restoreOperation RestoreOperation
	err = json.Unmarshal(res, &restoreOperation)
	if err != nil {
		return nil, err
	}

	return &restoreOperation, nil
}

//func (c *RestoreClient) buildGetApiRequest(apiToken string, restoreId int32) (*http.Request, error) {
//	baseUrl := c.BaseUrl
//	req, err := http.NewRequest(getRestoreLogsServiceMethod, fmt.Sprintf(getRestoreLogsServiceUrl, baseUrl, restoreId), nil)
//	if err != nil {
//		return nil, err
//	}
//	logzio_client.AddHttpHeaders(apiToken, req)
//
//	return req, err
//}

package drop_filters

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createDropFilterServiceUrl      = dropFiltersServiceEndpoint
	createDropFilterServiceMethod   = http.MethodPost
	createDropFilterMethodSuccess   = http.StatusOK
	createDropFilterMethodCreated   = http.StatusCreated
	createDropFilterMethodNoContent = http.StatusNoContent
	createDropFilterStatusNotFound  = http.StatusNotFound
)

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// CreateDropFilter creates a drop filter, returns the created drop filter if successful, an error otherwise
func (c *DropFiltersClient) CreateDropFilter(createDropFilter CreateDropFilter) (*DropFilter, error) {
	err := validateCreateDropFilterRequest(createDropFilter)
	if err != nil {
		return nil, err
	}

	dropFilterJson, err := json.Marshal(createDropFilter)
	if err != nil {
		return nil, err
	}

	//req, err := c.buildCreateApiRequest(c.ApiToken, dropFilterJson)
	//if err != nil {
	//	return nil, err
	//}
	//
	//jsonResponse, err := logzio_client.CreateHttpRequestBytesResponse(req)
	//if err != nil {
	//	return nil, err
	//}
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createDropFilterServiceMethod,
		Url:          fmt.Sprintf(createDropFilterServiceUrl, c.BaseUrl),
		Body:         dropFilterJson,
		SuccessCodes: []int{createDropFilterMethodSuccess, createDropFilterMethodNoContent, createDropFilterMethodCreated},
		NotFoundCode: createDropFilterStatusNotFound,
		ResourceId:   nil,
		ApiAction:    createDropFilterOperation,
	})

	if err != nil {
		return nil, err
	}

	var retVal DropFilter
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil

}

//func (c *DropFiltersClient) buildCreateApiRequest(apiToken string, jsonBytes []byte) (*http.Request, error) {
//	baseUrl := c.BaseUrl
//	req, err := http.NewRequest(createDropFilterServiceMethod, fmt.Sprintf(createDropFilterServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
//	logzio_client.AddHttpHeaders(apiToken, req)
//
//	return req, err
//}

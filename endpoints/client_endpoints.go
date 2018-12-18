package endpoints

import (
	"fmt"
	"github.com/jonboydell/logzio_client/client"
	"strings"
)

const (
	fldEndpointId          string = "id"
	fldEndpointType        string = "endpointType"
	fldEndpointTitle       string = "title"
	fldEndpointDescription string = "description"
	fldEndpointUrl         string = "url"
	endpointTypeSlack      string = "slack"
)

type EndpointType struct {
	Id           int64
	EndpointType string
	Title        string
	Description  string
	Url          string
	Message      string
}

func jsonEndpointToEndpoint(jsonEndpoint map[string]interface{}) EndpointType {
	endpoint := EndpointType{
		Id:           int64(jsonEndpoint[fldEndpointId].(float64)),
		EndpointType: jsonEndpoint[fldEndpointType].(string),
		Title:        jsonEndpoint[fldEndpointTitle].(string),
		Description:  jsonEndpoint[fldEndpointDescription].(string),
	}
	if endpointTypeSlack == strings.ToLower(endpoint.EndpointType) {
		endpoint.Url = jsonEndpoint[fldEndpointUrl].(string)
	}
	return endpoint
}


type Endpoints struct {
	client.Client
}

func New(apiToken string) (*Endpoints, error) {
	var c Endpoints
	c.ApiToken = apiToken
	if len(apiToken) > 0 {
		return &c, nil
	} else {
		return nil, fmt.Errorf("API token not defined")
	}
}
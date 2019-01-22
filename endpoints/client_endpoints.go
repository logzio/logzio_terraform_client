package endpoints

import (
	"fmt"
	"github.com/jonboydell/logzio_client/client"
	"strings"
)

const (
	fldEndpointId            string = "id"
	fldEndpointType          string = "endpointType"
	fldEndpointTitle         string = "title"
	fldEndpointDescription   string = "description"
	fldEndpointUrl           string = "url"
	fldEndpointMethod        string = "method"
	fldEndpointHeaders       string = "headers"
	fldEndpointBodyTemplate  string = "bodyTemplate"
	fldEndpointServiceKey    string = "serviceKey"
	fldEndpointApiToken      string = "apiToken"
	fldEndpointAppKey        string = "appKey"
	fldEndpointApiKey        string = "apiKey"
	fldEndpointRoutingKey    string = "routingKey"
	fldEndpointMessageType   string = "messageType"
	fldEndpointServiceApiKey string = "serviceApiKey"
)
const (
	endpointTypeSlack     string = "slack"
	endpointTypeCustom    string = "custom"
	endpointTypePagerDuty string = "pager-duty"
	endpointTypeBigPanda  string = "big-panda"
	endpointTypeDataDog   string = "data-dog"
	endpointTypeVictorOps string = "victorops"
)

type Endpoint struct {
	Id            int64             // all
	EndpointType  string            // all
	Title         string            // all
	Description   string            // all
	Url           string            // custom & slack
	Method        string            // custom
	Headers       interface{} // custom
	BodyTemplate  string // custom
	Message       string            // n.b. this is a hack to determine if there was an error (despite a 200 being returned)
	ServiceKey    string            // pager-duty
	ApiToken      string            // big-panda
	AppKey        string            // big-panda
	ApiKey        string            // data-dog
	RoutingKey    string            // victorops
	MessageType   string            // victorops
	ServiceApiKey string            // victorops
}

func jsonEndpointToEndpoint(jsonEndpoint map[string]interface{}) Endpoint {
	endpoint := Endpoint{
		Id:           int64(jsonEndpoint[fldEndpointId].(float64)),
		EndpointType: jsonEndpoint[fldEndpointType].(string),
		Title:        jsonEndpoint[fldEndpointTitle].(string),
		Description:  jsonEndpoint[fldEndpointDescription].(string),
	}

	switch strings.ToLower(endpoint.EndpointType) {
	case endpointTypeSlack:
		endpoint.Url = jsonEndpoint[fldEndpointUrl].(string)
	case endpointTypeCustom:
		endpoint.Url = jsonEndpoint[fldEndpointUrl].(string)
		endpoint.BodyTemplate = jsonEndpoint[fldEndpointBodyTemplate].(string)
		endpoint.Headers = jsonEndpoint[fldEndpointHeaders].(map[string]string)
		endpoint.Method = jsonEndpoint[fldEndpointMethod].(string)
	case endpointTypePagerDuty:
		endpoint.ServiceKey = jsonEndpoint[fldEndpointServiceKey].(string)
	case endpointTypeBigPanda:
		endpoint.ApiToken = jsonEndpoint[fldEndpointApiToken].(string)
		endpoint.AppKey = jsonEndpoint[fldEndpointAppKey].(string)
	case endpointTypeDataDog:
		endpoint.ApiKey = jsonEndpoint[fldEndpointApiKey].(string)
	case endpointTypeVictorOps:
		endpoint.RoutingKey = jsonEndpoint[fldEndpointRoutingKey].(string)
		endpoint.MessageType = jsonEndpoint[fldEndpointMessageType].(string)
		endpoint.ServiceApiKey = jsonEndpoint[fldEndpointServiceApiKey].(string)
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

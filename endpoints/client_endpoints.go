package endpoints

import (
	"fmt"
	"github.com/jonboydell/logzio_client/client"
	"strings"
)

const (
	fldEndpointId           string = "id"
	fldEndpointType         string = "endpointType"
	fldEndpointTitle        string = "title"
	fldEndpointDescription  string = "description"
	fldEndpointUrl          string = "url"
	fldEndpointMethod       string = "method"
	fldEndpointHeaders      string = "headers"
	fldEndpointBodyTemplate string = "bodyTemplate"

	endpointTypeSlack     string = "slack"
	endpointTypeCustom    string = "custom"
	endpointTypePagerDuty string = "pager-duty"
	endpointTypeBigPanda  string = "big-panda"
	endpointTypeDataDog   string = "data-dog"
	endpointTypeVictorOps string = "victorops"
)

type EndpointType struct {
	Id            int64                  // all
	EndpointType  string                 // all
	Title         string                 // all
	Description   string                 // all
	Url           string                 // custom & slack
	Method        string                 // custom
	Headers       string                 // custom
	BodyTemplate  map[string]interface{} // custom
	Message       string                 // n.b. this is a hack to determine if there was an error (despite a 200 being returned)
	ServiceKey    string                 // pager-duty
	ApiToken      string                 // big-panda
	AppKey        string                 // big-panda
	ApiKey        string                 // data-dog
	RoutingKey    string                 // victorops
	MessageType   string                 // victorops
	ServiceApiKey string                 // victorops
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
	if endpointTypeCustom == strings.ToLower(endpoint.EndpointType) {
		endpoint.Url = jsonEndpoint[fldEndpointUrl].(string)
		endpoint.BodyTemplate = jsonEndpoint[fldEndpointBodyTemplate].(map[string]interface{})
		endpoint.Headers = jsonEndpoint[fldEndpointHeaders].(string)
		endpoint.Method = jsonEndpoint[fldEndpointMethod].(string)
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

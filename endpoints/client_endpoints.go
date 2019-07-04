package endpoints

import (
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	endpointServiceEndpoint = "%s/v1/endpoints"
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
	EndpointTypeSlack     string = "slack"
	EndpointTypeCustom    string = "custom"
	EndpointTypePagerDuty string = "pager-duty"
	EndpointTypeBigPanda  string = "big-panda"
	EndpointTypeDataDog   string = "data-dog"
	EndpointTypeVictorOps string = "victorops"
)

type Endpoint struct {
	Id            int64             // all
	EndpointType  string            // all
	Title         string            // all
	Description   string            // all
	Url           string            // custom & slack
	Method        string            // custom
	Headers       map[string]string // custom
	BodyTemplate  interface{}       // custom
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
	case EndpointTypeSlack:
		endpoint.Url = jsonEndpoint[fldEndpointUrl].(string)
	case EndpointTypeCustom:
		endpoint.Url = jsonEndpoint[fldEndpointUrl].(string)
		endpoint.BodyTemplate = jsonEndpoint[fldEndpointBodyTemplate]
		headerMap := make(map[string]string)
		headerString := jsonEndpoint[fldEndpointHeaders].(string)
		headers := strings.Split(headerString, ",")
		for _, header := range headers {
			kv := strings.Split(header, "=")
			headerMap[kv[0]] = kv[1]
		}
		endpoint.Headers = headerMap
		endpoint.Method = jsonEndpoint[fldEndpointMethod].(string)
	case EndpointTypePagerDuty:
		endpoint.ServiceKey = jsonEndpoint[fldEndpointServiceKey].(string)
	case EndpointTypeBigPanda:
		endpoint.ApiToken = jsonEndpoint[fldEndpointApiToken].(string)
		endpoint.AppKey = jsonEndpoint[fldEndpointAppKey].(string)
	case EndpointTypeDataDog:
		endpoint.ApiKey = jsonEndpoint[fldEndpointApiKey].(string)
	case EndpointTypeVictorOps:
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

type endpointValidator = func(e Endpoint) (error, bool)
type endpointBuilder = func(a string, t string, e Endpoint) (*http.Request, error)
type endpointChecker = func(b []byte) error

func (c *Endpoints) makeEndpointRequest(endpoint interface{}, validator endpointValidator, builder endpointBuilder, checker endpointChecker) ([]byte, error, bool) {
	e := endpoint.(Endpoint)
	if err, ok := validator(e); !ok {
		return nil, err, false
	}
	req, _ := builder(c.ApiToken, e.EndpointType, e)
	httpClient := client.GetHttpClient(req)
	resp, _ := httpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{200}) {
		return nil, fmt.Errorf(errorCreateEndpointApiCallFailed, resp.StatusCode, jsonBytes), false
	}
	err := checker(jsonBytes)
	if err != nil {
		return nil, err, false
	}
	return jsonBytes, nil, true
}

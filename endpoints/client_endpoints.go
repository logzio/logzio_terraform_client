package endpoints

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
	"strings"
)

const (
	endpointServiceEndpoint = "%s/v1/endpoints"
	deleteEndpointMethod    = "DeleteEndpoint"
	getEndpointMethod       = "GetEndpoint"
	listEndpointMethod      = "ListEndpoint"
	updateEndpointMethod    = "UpdateEndpoint"
	createEndpointMethod    = "CreateEndpoint"
)

// supported endpoint types
const (
	EndpointTypeSlack          string = "slack"
	EndpointTypeCustom         string = "custom"
	EndpointTypePagerDuty      string = "pagerduty"
	EndpointTypeBigPanda       string = "bigpanda"
	EndpointTypeDataDog        string = "datadog"
	EndpointTypeVictorOps      string = "victorops"
	EndpointTypeOpsGenie       string = "opsgenie"
	EndpointTypeServiceNow     string = "servicenow"
	EndpointTypeMicrosoftTeams string = "microsoftteams"
)

type CreateOrUpdateEndpoint struct {
	Title         string      `json:"title"`                   // all
	Description   string      `json:"description"`             // all
	Url           string      `json:"url,omitempty"`           // slack, custom, serviceNow, microsoftTeams
	Method        string      `json:"method,omitempty'"`       // custom
	Headers       string      `json:"headers,omitempty"`       // custom
	BodyTemplate  interface{} `json:"bodyTemplate,omitempty"`  // custom
	ServiceKey    string      `json:"serviceKey,omitempty"`    // pagerDuty
	ApiToken      string      `json:"apiToken,omitempty"`      // bigPanda
	AppKey        string      `json:"appKey,omitempty"`        // bigPanda
	ApiKey        string      `json:"apiKey,omitempty"`        // dataDog, opsGenie
	RoutingKey    string      `json:"routingKey,omitempty"`    // victorOps
	MessageType   string      `json:"messageType,omitempty"`   // victorOps
	ServiceApiKey string      `json:"serviceApiKey,omitempty"` // victorOps
	Username      string      `json:"username,omitempty"`      // serviceNow
	Password      string      `json:"password,omitempty"`      // serviceNow
	Type          string      `json:"-"`                       // only for identification of the endpoint
}

type CreateOrUpdateEndpointResponse struct {
	Id int32 `json:"id"`
}

type Endpoint struct {
	Type          string      `json:"endpointType"`
	Id            int32       `json:"id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	Url           string      `json:"url,omitempty"`           // slack, custom, serviceNow, microsoftTeams
	Method        string      `json:"method,omitempty'"`       // custom
	Headers       string      `json:"headers,omitempty"`       // custom
	BodyTemplate  interface{} `json:"bodyTemplate,omitempty"`  // custom
	ServiceKey    string      `json:"serviceKey,omitempty"`    // pagerDuty
	ApiToken      string      `json:"apiToken,omitempty"`      // bigPanda
	AppKey        string      `json:"appKey,omitempty"`        // bigPanda
	ApiKey        string      `json:"apiKey,omitempty"`        // dataDog, opsGenie
	RoutingKey    string      `json:"routingKey,omitempty"`    // victorOps
	MessageType   string      `json:"messageType,omitempty"`   //victorOps
	ServiceApiKey string      `json:"serviceApiKey,omitempty"` // victorOps
	Username      string      `json:"username,omitempty"`      // serviceNow
	Password      string      `json:"password,omitempty"`      // serviceNow
}

type EndpointsClient struct {
	*client.Client
}

func New(apiToken, baseUrl string) (*EndpointsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &EndpointsClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

func (c *EndpointsClient) getURLByType(t string) string {
	switch strings.ToLower(t) {
	case EndpointTypeSlack:
		return "slack"
	case EndpointTypeCustom:
		return "custom"
	case EndpointTypePagerDuty:
		return "pager-duty"
	case EndpointTypeBigPanda:
		return "big-panda"
	case EndpointTypeDataDog:
		return "data-dog"
	case EndpointTypeVictorOps:
		return "victorops"
	case EndpointTypeOpsGenie:
		return "ops-genie"
	case EndpointTypeServiceNow:
		return "service-now"
	case EndpointTypeMicrosoftTeams:
		return "microsoft-teams"
	default:
		panic(fmt.Sprintf("unsupported endpoint type %s", t))
	}
}

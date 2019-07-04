package endpoints

import (
	"fmt"
	"strings"
)

func validSlackEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0
}

func validCustomEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0 &&
		len(endpointType.Method) > 0
}

func validPagerDutyEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Description) > 0 &&
		len(endpointType.ServiceKey) > 0
}

func validBigPandaEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Description) > 0 &&
		len(endpointType.ApiToken) > 0 && len(endpointType.AppKey) > 0
}

func validDataDogEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Description) > 0 &&
		len(endpointType.ApiKey) > 0
}

func validVictorOpsEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Description) > 0 &&
		len(endpointType.RoutingKey) > 0 && len(endpointType.MessageType) > 0 && len(endpointType.ServiceApiKey) > 0
}

// Validates an endpoint request for correctness given it's type, returns an error and FALSE if validation is failed, true otherwise
func ValidateEndpointRequest(endpoint Endpoint) (error, bool) {
	if strings.EqualFold(EndpointTypeSlack, endpoint.EndpointType) && validSlackEndpoint(endpoint) {
		return nil, true
	}

	if strings.EqualFold(EndpointTypeCustom, endpoint.EndpointType) && validCustomEndpoint(endpoint) {
		return nil, true
	}

	if strings.EqualFold(EndpointTypePagerDuty, endpoint.EndpointType) && validPagerDutyEndpoint(endpoint) {
		return nil, true
	}

	if strings.EqualFold(EndpointTypeBigPanda, endpoint.EndpointType) && validBigPandaEndpoint(endpoint) {
		return nil, true
	}

	if strings.EqualFold(EndpointTypeDataDog, endpoint.EndpointType) && validDataDogEndpoint(endpoint) {
		return nil, true
	}

	if strings.EqualFold(EndpointTypeVictorOps, endpoint.EndpointType) && validVictorOpsEndpoint(endpoint) {
		return nil, true
	}

	return fmt.Errorf(errorInvalidEndpointDefinition, endpoint, endpoint.EndpointType), false
}

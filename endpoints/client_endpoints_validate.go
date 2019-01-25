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

func ValidateEndpointRequest(endpoint Endpoint) error {
	if strings.EqualFold(EndpointTypeSlack, endpoint.EndpointType) && validSlackEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(EndpointTypeCustom, endpoint.EndpointType) && validCustomEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(EndpointTypePagerDuty, endpoint.EndpointType) && validPagerDutyEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(EndpointTypeBigPanda, endpoint.EndpointType) && validBigPandaEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(EndpointTypeDataDog, endpoint.EndpointType) && validDataDogEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(EndpointTypeVictorOps, endpoint.EndpointType) && validVictorOpsEndpoint(endpoint) {
		return nil
	}

	return fmt.Errorf(errorInvalidEndpointDefinition, endpoint, endpoint.EndpointType)
}

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
	if strings.EqualFold(endpointTypeSlack, endpoint.EndpointType) && validSlackEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(endpointTypeCustom, endpoint.EndpointType) && validCustomEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(endpointTypePagerDuty, endpoint.EndpointType) && validPagerDutyEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(endpointTypeBigPanda, endpoint.EndpointType) && validBigPandaEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(endpointTypeDataDog, endpoint.EndpointType) && validDataDogEndpoint(endpoint) {
		return nil
	}

	if strings.EqualFold(endpointTypeVictorOps, endpoint.EndpointType) && validVictorOpsEndpoint(endpoint) {
		return nil
	}

	return fmt.Errorf(errorInvalidEndpointDefinition, endpoint, endpoint.EndpointType)
}

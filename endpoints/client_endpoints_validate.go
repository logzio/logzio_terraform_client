package endpoints

import "fmt"

func validSlackEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0
}

func validCustomEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0 &&
		len(endpointType.Method) > 0 && len(endpointType.Headers) > 0 &&
		len(endpointType.BodyTemplate) > 0
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
	if endpointTypeSlack == endpoint.EndpointType && validSlackEndpoint(endpoint) {
		return nil
	}

	if endpointTypeCustom == endpoint.EndpointType && validCustomEndpoint(endpoint) {
		return nil
	}

	if endpointTypePagerDuty == endpoint.EndpointType && validPagerDutyEndpoint(endpoint) {
		return nil
	}

	if endpointTypeBigPanda == endpoint.EndpointType && validBigPandaEndpoint(endpoint) {
		return nil
	}

	if endpointTypeDataDog == endpoint.EndpointType && validDataDogEndpoint(endpoint) {
		return nil
	}

	if endpointTypeVictorOps == endpoint.EndpointType && validVictorOpsEndpoint(endpoint) {
		return nil
	}

	return fmt.Errorf(errorInvalidEndpointDefinition, endpoint, endpoint.EndpointType)
}

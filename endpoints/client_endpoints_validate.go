pacakge endpoints

func validSlackEndpoint(endpointType EndpointType) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0
}

func validCustomEndpoint(endpointType EndpointType) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0 &&
		len(endpointType.Method) > 0 && len(endpointType.Headers) > 0 &&
		len(endpointType.BodyTemplate) > 0
}

func validPagerDutyEndpoint(endpointType EndpointType) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0 &&
		len(endpointType.ServiceKey) > 0
}

func validBigPandaEndpoint(endpointType EndpointType) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0 &&
		len(endpointType.ApiToken) > 0 && len(endpointType.AppKey) > 0
}

func validDataDogEndpoint(endpointType EndpointType) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0 &&
		len(endpointType.AppKey) > 0
}

func validVictorOpsEndpoint(endpointType EndpointType) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0 &&
		len(endpointType.RoutingKey) > 0 && len(endpointType.MessageType) > 0 && len(endpointType.ServiceApiKey) > 0
}

func ValidateEndpointRequest(endpoint EndpointType, service string) error {
	if service == endpointTypeSlack && validSlackEndpoint(endpoint) {
		return nil
	}

	if service == endpointTypeCustom && validCustomEndpoint(endpoint) {
		return nil
	}

	if service == endpointTypePagerDuty && validPagerDutyEndpoint(endpoint) {
		return nil
	}

	if service == endpointTypeBigPanda && validBigPandaEndpoint(endpoint) {
		return nil
	}

	if service == endpointTypeDataDog && validDataDogEndpoint(endpoint) {
		return nil
	}

	if service == endpointTypeVictorOps && validVictorOpsEndpoint(endpoint) {
		return nil
	}

	return fmt.Errorf(errorInvalidEndpointDefinition, endpoint, service)
}
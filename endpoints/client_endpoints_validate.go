package endpoints

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

func validateCreateOrUpdateEndpointRequest(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.Title) == 0 {
		return fmt.Errorf("title must be set")
	}

	if len(endpoint.Type) == 0 {
		return fmt.Errorf("type must be set")
	}

	if !isValidEndpoint(endpoint.Type) {
		return fmt.Errorf("unsupported endpoint type %s", endpoint.Type)
	}

	switch endpoint.Type {
	case EndpointTypeSlack:
		return validateCreateOrUpdateEndpointRequestSlack(endpoint)
	case EndpointTypeCustom:
		return validateCreateOrUpdateEndpointRequestCustom(endpoint)
	case EndpointTypePagerDuty:
		return validateCreateOrUpdateEndpointRequestPagerDuty(endpoint)
	case EndpointTypeBigPanda:
		return validateCreateOrUpdateEndpointRequestBigPanda(endpoint)
	case EndpointTypeDataDog:
		return validateCreateOrUpdateEndpointRequestDataDog(endpoint)
	case EndpointTypeVictorOps:
		return validateCreateOrUpdateEndpointRequestVictorOps(endpoint)
	case EndpointTypeOpsGenie:
		return validateCreateOrUpdateEndpointRequestOpsGenie(endpoint)
	case EndpointTypeServiceNow:
		return validateCreateOrUpdateEndpointRequestServiceNow(endpoint)
	case EndpointTypeMicrosoftTeams:
		return validateCreateOrUpdateEndpointRequestMicrosoftTeams(endpoint)
	default:
		return fmt.Errorf("unsupported endpoint type %s", endpoint.Type)

	}
}

func validateCreateOrUpdateEndpointRequestSlack(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.Url) == 0 {
		return fmt.Errorf("url must be set for type %s", EndpointTypeSlack)
	}

	return nil
}

func validateCreateOrUpdateEndpointRequestCustom(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.Url) == 0 {
		return fmt.Errorf("url must be set for type %s", EndpointTypeCustom)
	}
	if len(endpoint.Method) == 0 {
		return fmt.Errorf("method must be set for type %s", EndpointTypeCustom)
	}

	validMethods := []string{http.MethodPost, http.MethodPut, http.MethodGet}
	if !logzio_client.Contains(validMethods, endpoint.Method) {
		return fmt.Errorf("method should be one of %s", validMethods)
	}

	return nil
}

func validateCreateOrUpdateEndpointRequestPagerDuty(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.ServiceKey) == 0 {
		return fmt.Errorf("service key must be set for type %s", EndpointTypePagerDuty)
	}

	return nil
}

func validateCreateOrUpdateEndpointRequestBigPanda(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.ApiToken) == 0 {
		return fmt.Errorf("api token must be set for type %s", EndpointTypeBigPanda)
	}
	if len(endpoint.AppKey) == 0 {
		return fmt.Errorf("app key must be set for type %s", EndpointTypeBigPanda)
	}

	return nil
}

func validateCreateOrUpdateEndpointRequestDataDog(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.ApiKey) == 0 {
		return fmt.Errorf("api key must be set for type %s", EndpointTypeDataDog)
	}

	return nil
}

func validateCreateOrUpdateEndpointRequestVictorOps(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.RoutingKey) == 0 {
		return fmt.Errorf("routing key must be set for type %s", EndpointTypeVictorOps)
	}
	if len(endpoint.MessageType) == 0 {
		return fmt.Errorf("message type must be set for type %s", EndpointTypeVictorOps)
	}
	if len(endpoint.ServiceApiKey) == 0 {
		return fmt.Errorf("service api key must be set for type %s", EndpointTypeVictorOps)
	}

	return nil
}

func validateCreateOrUpdateEndpointRequestOpsGenie(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.ApiKey) == 0 {
		return fmt.Errorf("api key must be set for type %s", EndpointTypeOpsGenie)
	}

	return nil
}

func validateCreateOrUpdateEndpointRequestServiceNow(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.Username) == 0 {
		return fmt.Errorf("username must be set for type %s", EndpointTypeServiceNow)
	}

	if len(endpoint.Password) == 0 {
		return fmt.Errorf("password must be set for type %s", EndpointTypeServiceNow)
	}

	if len(endpoint.Url) == 0 {
		return fmt.Errorf("url must be set for type %s", EndpointTypeServiceNow)
	}

	return nil
}

func validateCreateOrUpdateEndpointRequestMicrosoftTeams(endpoint CreateOrUpdateEndpoint) error {
	if len(endpoint.Url) == 0 {
		return fmt.Errorf("url must be set for type %s", EndpointTypeMicrosoftTeams)
	}

	return nil
}

func isValidEndpoint(endpointType string) bool {
	validEndpoints := []string{EndpointTypeSlack, EndpointTypeCustom, EndpointTypePagerDuty, EndpointTypeBigPanda,
		EndpointTypeDataDog, EndpointTypeVictorOps, EndpointTypeOpsGenie, EndpointTypeServiceNow, EndpointTypeMicrosoftTeams}
	for _, endpoint := range validEndpoints {
		if endpoint == endpointType {
			return true
		}
	}

	return false
}

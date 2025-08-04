package drop_filters

import (
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	dropFiltersServiceEndpoint    string = "%s/v1/drop-filters"
	createDropFilterOperation            = "CreateDropFilter"
	activateDropFilterOperation          = "ActivateDropFilter"
	deactivateDropFilterOperation        = "DeactivateDropFilter"
	deleteDropFilterOperation            = "DeleteDropFilter"
	retrieveDropFiltersOperation         = "RetrieveDropFilters"

	dropFilterResourceName = "drop filter"
)

type DropFiltersClient struct {
	*client.Client
}

type CreateDropFilter struct {
	LogType         string                 `json:"logType,omitempty"`
	FieldConditions []FieldConditionObject `json:"fieldConditions"`
	ThresholdInGB   float64                `json:"thresholdInGB,omitempty"`
}

type FieldConditionObject struct {
	FieldName string      `json:"fieldName"`
	Value     interface{} `json:"value"`
}

type DropFilter struct {
	Id             string                 `json:"id"`
	Active         bool                   `json:"active"`
	LogType        string                 `json:"logType"`
	FieldCondition []FieldConditionObject `json:"fieldConditions"`
	ThresholdInGB  float64                `json:"thresholdInGB,omitempty"`
}

func New(apiToken, baseUrl string) (*DropFiltersClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &DropFiltersClient{
		Client: client.New(apiToken, baseUrl),
	}

	return c, nil
}

func validateCreateDropFilterRequest(filter CreateDropFilter) error {
	if len(filter.FieldConditions) == 0 {
		return fmt.Errorf("FieldConditions must be set")
	}

	for _, condition := range filter.FieldConditions {
		if len(condition.FieldName) == 0 {
			return fmt.Errorf("FieldName must be set")
		}

		if condition.Value == nil {
			return fmt.Errorf("Value must be set")
		}
	}

	return nil
}

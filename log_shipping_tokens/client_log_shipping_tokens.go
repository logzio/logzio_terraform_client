package log_shipping_tokens

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	logShippingTokensServiceEndpoint         string = "%s/v1/log-shipping/tokens"
)

const (
	operationGetLogShippingToken = "GetLogShippingToken"
	operationUpdateLogShippingToken = "UpdateLogShippingToken"
	operationDeleteLogShippingToken = "DeleteLogShippingToken"
	operationGetLogShippingTokensLimits = "GetNumberOfAvailableTokens"
	operationRetrieveLogShippingTokens = "RetrieveLogShippingTokens"

	retrieveSortFieldCreatedAtValue = "CREATED_AT"
	retrieveSortFieldNameValue = "NAME"
)

type LogShippingTokensClient struct {
	*client.Client
}

type CreateLogShippingToken struct {
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
}

type LogShippingToken struct {
	Name string `json:"name"`
	Id int32 `json:"id"`
	Token string `json:"token"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	Enabled bool `json:"enabled"`
}

type LogShippingTokensLimits struct {
	MaxAllowedTokens int32 `json:"maxAllowedTokens"` // The number of log shipping tokens this account can have.
	NumOfEnabledTokens int32 `json:"numOfEnabledTokens"` // The number of log shipping tokens currently enabled for this account.
}

type RetrieveLogShippingTokensRequest struct {
	Filter ShippingTokensFilterRequest `json:"filter"`
	Sort []ShippingTokensSortRequest `json:"sort"`
	Pagination ShippingTokensPaginationRequest `json:"pagination,omitempty"`
}

type ShippingTokensFilterRequest struct {
	Enabled bool `json:"enabled"`
}

type ShippingTokensSortRequest struct {
	Field string `json:"field"`
	Descending bool `json:"descending"`
}

type ShippingTokensPaginationRequest struct {
	PageNumber int32 `json:"pageNumber,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
}

type RetrieveLogShippingTokensResponse struct {
	Total int32 `json:"total"`
	Results []LogShippingToken `json:"results"`
	Pagination ShippingTokensPaginationRequest `json:"pagination"`
}

func New(apiToken, baseUrl string) (*LogShippingTokensClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &LogShippingTokensClient{
		Client: client.New(apiToken, baseUrl),
	}

	return c, nil
}

func validateCreateLogShippingTokenRequest(token CreateLogShippingToken) error {
	if len(token.Name) == 0 {
		return fmt.Errorf("name must be set")
	}

	return nil
}

func validateRetrieveLogShippingTokensRequest(retrieveRequest RetrieveLogShippingTokensRequest) error {
	validSortFieldValues :=  []string{retrieveSortFieldCreatedAtValue, retrieveSortFieldNameValue}

	if len(retrieveRequest.Sort) > 0 {
		for _, sort := range retrieveRequest.Sort {
			if !logzio_client.Contains(validSortFieldValues, sort.Field) {
				return fmt.Errorf("sort's Field must be one of %s", validSortFieldValues)
			}
		}
	}

	return nil
}
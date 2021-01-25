package shared_tokens

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	sharedTokenServiceEndpoint = "%s/v1/shared-tokens"
	loggerName = "logzio-client"
	fldSharedTokenId = "id"
	fldSharedTokenName = "name"
	fldToken = "token"
	fldSharedTokenFilters = "filters"
)
type SharedToken struct {
	Id        int32
	Name      string
	Token     string
	FilterIds []int32
}

type SharedTokenClient struct {
	*client.Client
	logger                  hclog.Logger
}

// Creates a new entry point into the shared token functions, accepts the user's logz.io API token and account Id
func New(apiToken string, baseUrl string) (*SharedTokenClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	c := &SharedTokenClient{
		Client: client.New(apiToken, baseUrl),
		logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Debug,
			Name:       loggerName,
			JSONFormat: true,
		}),
	}
	return c, nil
}

func jsonToSharedToken(json map[string]interface{}) SharedToken {
	filters := json[fldSharedTokenFilters].([]interface{})
	var intFilters []int32
	for _, num := range filters {
		intFilters = append(intFilters, int32(num.(float64)))
	}
	sharedToken := SharedToken{
		Id:        int32(json[fldSharedTokenId].(float64)),
		Name:      json[fldSharedTokenName].(string),
		Token:     json[fldToken].(string),
		FilterIds: intFilters,
	}
	return sharedToken
}

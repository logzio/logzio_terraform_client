package shared_tokens

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/logzio/logzio_terraform_client/client"
)

const (
	sharedTokenServiceEndpoint = "%s/v1/shared-tokens/"
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
	sharedToken := SharedToken{
		Id:        json[fldSharedTokenId].(int32),
		Name:      json[fldSharedTokenName].(string),
		Token:     json[fldToken].(string),
		FilterIds: json[fldSharedTokenFilters].([]int32),
	}
	return sharedToken
}

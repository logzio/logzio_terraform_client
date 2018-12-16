package logzio_client

import (
	"fmt"
	"testing"
)

func createValidEndpoint() EndpointType {
	return EndpointType{
		Title: "my endpoint",
		Description: "my description",
		EndpointType: "slack",
		Url: "https://this.is.com/some/webhook",
	}
}

func createInvalidEndpoint() EndpointType {
	return EndpointType{
		Title: "my endpoint",
		Description: "my description",
		EndpointType: "slack",
		Url: "https://someUrl",
	}
}

func TestCreateValidEndpoint(t *testing.T) {
	api_token := getApiToken(t)

	var client *Client
	client = New(api_token)

	var endpoint *EndpointType
	var err error
	endpoints := []int64{}

	endpoint, err = client.createEndpoint(createValidEndpoint())
	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "CreateEndpoint", err)
	}
	endpoints = append(endpoints, endpoint.Id)

	endpointId := fmt.Sprintf("%d", endpoint.Id)

	if len(endpointId) == 0 {
		t.Fatalf("%s should have a length > 0: %v", endpointId, err)
	}

	list, err := client.ListEndpoints()
	if err != nil {
		t.Fatalf("%q should not have raised an error: %v", "ListEndpoints", err)
	}

	if len(list) == 0 {
		t.Fatalf("%q endpoints list should be len > 0 : %d", "ListEndpoints", len(list))
	}

	// clean up any created alerts
	for x := 0; x < len(endpoints); x++ {
		client.DeleteEndpoint(endpoints[x])
	}
}


func TestCreateInvalidEndpoint(t *testing.T) {
	api_token := getApiToken(t)

	var client *Client
	client = New(api_token)

	createInvalidEndpoint := createInvalidEndpoint()

	var err error
	_, err = client.createEndpoint(createInvalidEndpoint)
	if err == nil {
		t.Fatalf("%q should have raised an error: %v", "CreateEndpoint", err)
	}
}
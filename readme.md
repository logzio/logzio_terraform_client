# Logz.io client library

DEVELOP - [![Build Status](https://travis-ci.org/jonboydell/logzio_client.svg?branch=develop)](https://travis-ci.org/jonboydell/logzio_client) [![Coverage Status](https://coveralls.io/repos/github/jonboydell/logzio_client/badge.svg?branch=develop)](https://coveralls.io/github/jonboydell/logzio_client?branch=develop)

MASTER - [![Build Status](https://travis-ci.org/jonboydell/logzio_client.svg?branch=master)](https://travis-ci.org/jonboydell/logzio_client)

Client library for logz.io - supports Alerts (reasonably well) and Endpoints (experimentally) at the moment

##### Logz.io API support

|api  |method|support     |implementation|test coverage|
|-----|------|------------|--------------|-------------|
|alert|create|ready|`alerts.Alerts.CreateAlert` |yes          |
|alert|delete|ready|`alerts.Alerts.DeleteAlert` |yes          |
|alert|list  |ready|`alerts.Alerts.ListAlerts`  |yes          |
|alert|update|ready|`alerts.Alerts.UpdateAlert` |yes          |
|alert|read  |ready|`alerts.Alerts.GetAlert`    |yes          |
|endpoints|create|experimental|`endpoints.Endpoints.CreateEndpoint` |yes          |
|endpoints|delete|experimental|`endpoints.Endpoints.DeleteEndpoint` |yes          |
|endpoints|list  |experimental|`endpoints.Endpoints.ListEndpoints`  |yes          |
|endpoints|update|experimental|`endpoints.Endpoints.UpdateEndpoint` |yes          |
|endpoints|read  |experimental|`endpoints.Endpoints.GetEndpoint`    |yes          |


- experimental - it works but not all permutations of parameters are tested
- ready - it works and a large proportion of permutation of parameter are tested

##### Very basic usage

Currently, the `New` functions exist in the relevant API package (e.g. `alerts` or `endpoints`)

```go
package main

import (
 "github.com/jonboydell/logzio_client/alerts"
)

func main() {
	var x alerts.CreateAlertType
	var y *alerts.Alerts
	
	y, _ = alerts.New("my_api_token")
	y.CreateAlert(x)
}

```

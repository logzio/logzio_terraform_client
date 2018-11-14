# Logz.io client library

[![Build Status](https://travis-ci.org/jonboydell/logzio_client.svg?branch=develop)](https://travis-ci.org/jonboydell/logzio_client)

Client library for logz.io - only supports Alerts (experimentally) at the moment

##### API support

|api  |method|support     |implementation|test coverage|
|-----|------|------------|--------------|-------------|
|alert|create|experimental|`CreateAlert` |yes          |
|alert|delete|experimental|`DeleteAlert` |yes          |
|alert|list  |experimental|`ListAlerts`  |yes          |
|alert|update|none        |              |             |
|alert|read  |none        |`GetAlert`    |             |

##### Basic usage

```go
package main

import "github.com/jonboydell/logzio_client"

client := logzio_client.New(api_token)
alerts, err := client.ListAlerts()
```
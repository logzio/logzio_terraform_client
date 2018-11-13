# Logz.io client library

Client library for logz.io - only supports Alerts (experimentally) at the moment

##### API support

|api  |method|support     |implementation|
|-----|------|------------|--------------|
|alert|create|experimental|`CreateAlert` |
|alert|delete|experimental|`DeleteAlert` |
|alert|list  |experimental|`ListAlerts`  |
|alert|update|none        |              |
|alert|read  |none        |              |

##### Basic usage

```go
package main

import "github.com/jonboydell/logzio_client"

client := logzio_client.New(api_token)
alerts, err := client.ListAlerts()
```
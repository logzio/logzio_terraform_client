# Logz.io Terraform client library

Client library for Logz.io API, see below for supported endpoints.

The primary purpose of this library is to act as the API interface for the logz.io Terraform provider.
To use it, you'll need to [create an API token](https://app.logz.io/#/dashboard/settings/api-tokens) and provide it to the client library along with your logz.io regional [API server address](https://docs.logz.io/user-guide/accounts/account-region.html#regions-and-urls).

The library currently supports the following API endpoints:
* [Alerts V2](https://github.com/logzio/logzio_terraform_client/tree/master/alerts_v2).
* [Users](https://github.com/logzio/logzio_terraform_client/tree/master/users).
* [Sub-accounts](https://github.com/logzio/logzio_terraform_client/tree/master/sub_accounts).
* [Endpoints](https://github.com/logzio/logzio_terraform_client/tree/master/endpoints).
* [Log shipping tokens](https://github.com/logzio/logzio_terraform_client/tree/master/log_shipping_tokens).
* [Drop filters](https://github.com/logzio/logzio_terraform_client/tree/master/drop_filters).
* [Archive logs](https://github.com/logzio/logzio_terraform_client/tree/master/archive_logs).
* [Restore logs](https://github.com/logzio/logzio_terraform_client/tree/master/restore_logs).
* [Authentication groups](https://docs.logz.io/api/#tag/Authentication-groups).
* [Kibana Objects](https://docs.logz.io/api/#tag/Import-or-export-Kibana-objects).
* [S3 Fetcher](https://docs.logz.io/api/#tag/Connect-to-S3-Buckets).
* [Grafana dashboards](https://docs.logz.io/api/#operation/createDashboard).
* [Grafana folders](https://api-docs.logz.io/docs/logz/get-all-folders).
* [Grafana Alert Rules API](https://docs.logz.io/api/#tag/Grafana-alerting-provisioning).
* [Grafana datasource](https://docs.logz.io/api/#operation/getDatasourceByAccount).
* [Grafana Notification Policy](https://api-docs.logz.io/docs/logz/route-get-policy-tree).
* [Grafana Contact Point](https://docs.logz.io/api/#tag/Grafana-contact-points).
* [Metrics Accounts](https://api-docs.logz.io/docs/logz/create-a-new-metrics-account)
* [Metrics drop filters](./drop_metrics/README.md) <!--- This should be replaced with the proper docs link once released. -->
* [Metrics roll-up rules](./metrics_rollup_rules/README.md) <!--- This should be replaced with the proper docs link once released. -->

#### Contributing

1. Clone this repo locally.
2. As this package uses Go modules, make sure you are outside of `$GOPATH` or you have the `GO111MODULE=on` environment variable set. Then run `go get` to pull down the dependencies.
3. Use `logzio_client.CallLogzioApi` when you need to make a Logz.io API call.
4. Use structs to represent the requests/responses body, rather than maps.
5. Sample responses for tests should be under `testdata/fixtures`.

##### Run tests
`go test -v -race ./...`

### Trademark Disclaimer

Terraform is a trademark of HashiCorp, Inc.

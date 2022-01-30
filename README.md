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

#### Contributing

1. Clone this repo locally.
2. As this package uses Go modules, make sure you are outside of `$GOPATH` or you have the `GO111MODULE=on` environment variable set. Then run `go get` to pull down the dependencies.
3. Use `logzio_client.CallLogzioApi` when you need to make a Logz.io API call.
4. Use structs to represent the requests/responses body, rather than maps.
5. Sample responses for tests should be under `testdata/fixtures`.

##### Run tests
`go test -v -race ./...`

### Changelog

- v1.10.1
  - Bug fix - **custom endpoint**: allow empty string for Headers field. 
- v1.10.0
    - Add [Authentication groups API](https://docs.logz.io/api/#tag/Authentication-groups).
    - Add tests to retrieve archive.
    - Improve tests.
<details>
  <summary markdown="span">Exapnd to check old versions </summary>
- v1.9.1
    - Bug fix - adjust "not found" message to all resources.
- v1.9.0
    - Add [Archive logs API](https://docs.logz.io/api/#tag/Archive-logs).
    - Add [Restore logs API](https://docs.logz.io/api/#tag/Restore-logs).
- v1.8.0
    - `sub_accounts`:
        - Add `flexible` & `reservedDailyGB`.
        - **Breaking changes:** refactor resource.
    - `endpoints`:
        - **Breaking changes:** refactor resource.
        - Add new endpoint types (OpsGenie, ServiceNow, Microsoft Teams).
- v1.7.0
    - Add [drop filters API](https://docs.logz.io/api/#tag/Drop-filters).
- v1.6.0
    - Add [log shipping tokens API](https://docs.logz.io/api/#tag/Manage-log-shipping-tokens) compatibility.
- v1.5.3
    - Fix for `sub account`: return token & account id on Create. 
- v1.5.2
    - Fix `custom endpoint` -empty headers bug.
    - Allow empty array for sharing accounts in `sub account`.
- v1.5.1
    - Fix alerts_v2 sort bug.
- v1.5
    - Add alerts v2 compatibility.
- v1.3.2
   - fix client custom endpoint headers bug
   - improve tests 
- v1.3
    - unnecessary resource updates bug fix.
    - support tags in alerts
- v1.2
    - Add subaccount support

</details>

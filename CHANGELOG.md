# Changes by Version

<!-- next version -->
## v1.25.1
- Merge the `Create` and `Update` request schemas for `DropMetric` to simplify the API.
- Remove redundant `FilterType` and filter type constants from `DropMetric` API.

## v1.25.0
- Add [Metrics drop filters API](./drop_metrics/README.md).
- Remove restriction on `LogsType` in S3 Fetcher to allow custom type names.

## v1.24.1
- Add `ThresholdInGB` support for drop filters api
## v1.24.0
- Support Warm Retention settings in `sub_accounts`.
  - Add `snapSearchRetentionDays` to create and update requests
  - Add new fields to sub account data: `IsCapped`, `SharedGB`, `TotalTimeBasedDailyGB`, `IsOwner`.

## v1.23.2
- Grafana Contact Point `teams` now refers to Microsoft Teams Workflows Contact Point. The old `teams` endpoint was deprecated in grafana 10.

## v1.23.1
- Grafana Alerts API `ExecErrState` is no longer configurable as of Grafana 10
- Add new Microsoft Teams Workflows Contact Point

## v1.23.0
- Upgrade to Go 1.24.
- Upgrade Grafana API to match Grafana 10.
  - **Breaking changes:**
    - Grafana Alert API `for` field is now a String.
    - Grafana Contact Point API `provenance` field is always overwritten to `api`, regardless of any other value set.

## v1.22.0
- Validate account existence before updating it, and set an empty account name if the name did not change to prevent API errors.

## v1.21.0
- Add [Metrics Accounts API](https://api-docs.logz.io/docs/logz/create-a-new-metrics-account).

## v1.20.1
- Add limitation, Grafana Alert name cannot contain `/` or `\`.

## v1.20.0
- Add [Grafana Contact Point API](https://api-docs.logz.io/docs/logz/route-get-contactpoints).
- Add Grafana Notification Policy API.

## v1.19.0
- Add [Grafana Datasource API](https://api-docs.logz.io/docs/logz/get-datasource-by-account) partial support for specific endpoint.

## v1.18.0
- Add [Grafana Alert Rules API](https://api-docs.logz.io/docs/logz/get-alert-rules) support.

## v1.17.0
- Add Grafana Folders API.
- Remove deprecated `alerts` (v1).

## v1.16.0
- Add [Grafana Dashboards API](https://api-docs.logz.io/docs/logz/create-dashboard) support.

## v1.15.0
- Add [S3 Fetcher](https://api-docs.logz.io/docs/logz/create-buckets).

## v1.14.0
- `alerts_v2` - support new field `schedule`

## v1.13.1
- Add retry mechanism for requests.

## v1.13.0
- Bug fix - **sub_accounts**: field `ReservedDailyGB` in requests can be 0.

## v1.12.0
- Upgrade to Go 1.18.
- Refactor `users`, adjust to the recent API fields.
- Add field `UserName` to `restore` initiate request, to match recent API fields.

## v1.11.0
- Add [Kibana Objects](https://api-docs.logz.io/docs/logz/import-or-export-kibana-objects).

## v1.10.3
- Bug fix - **sub_accounts**: omit maxDailyGb if needed.

## v1.10.2
- Bug fix - **alerts_v2**: allow sending with columns without sort.

## v1.10.1
- Bug fix - **custom endpoint**: allow empty string for Headers field.

## v1.10.0
- Add [Authentication groups API](https://api-docs.logz.io/docs/logz/authentication-groups).
- Add tests to retrieve archive.
- Improve tests.

## v1.9.1
- Bug fix - adjust "not found" message to all resources.

## v1.9.0
- Add [Archive logs API](https://api-docs.logz.io/docs/logz/archive-logs).
- Add [Restore logs API](https://api-docs.logz.io/docs/logz/restore-logs).

## v1.8.0
- `sub_accounts`:
  - Add `flexible` & `reservedDailyGB`.
  - **Breaking changes:** refactor resource.
- `endpoints`:
  - **Breaking changes:** refactor resource.
  - Add new endpoint types (OpsGenie, ServiceNow, Microsoft Teams).

## v1.7.0
- Add [drop filters API](https://api-docs.logz.io/docs/logz/drop-filters).

## v1.6.0
- Add [log shipping tokens API](https://api-docs.logz.io/docs/logz/manage-log-shipping-tokens) compatibility.

## v1.5.3
- Fix for `sub account`: return token & account id on Create. 

## v1.5.2
- Fix `custom endpoint` -empty headers bug.
- Allow empty array for sharing accounts in `sub account`.

## v1.5.1
- Fix alerts_v2 sort bug.

## v1.5
- Add alerts v2 compatibility.

## v1.3.2
- fix client custom endpoint headers bug
- improve tests

## v1.3
- unnecessary resource updates bug fix.
- support tags in alerts

## v1.2
- Add subaccount support
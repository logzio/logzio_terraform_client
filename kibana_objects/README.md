# Kibana Objects
Compatible with Logz.io's [Kibana objects API](https://api-docs.logz.io/docs/logz/import-or-export-kibana-objects).

To import a new Kibana Object:

```go
client, _ := kibana_objects.New(apiToken, apiServerAddress)
source := `{
        "search": {
          "columns": [
            "message"
          ],
          "sort": [
            "@timestamp",
            "desc"
          ],
          "id": "tf-client-test",
          "title": "tf-client-test",
          "version": 1,
          "kibanaSavedObjectMeta": {
            "searchSourceJSON": "{\"highlight\":{\"pre_tags\":[\"@kibana-highlighted-field@\"],\"post_tags\":[\"@/kibana-highlighted-field@\"],\"fields\":{\"*\":{}},\"fragment_size\":2147483647},\"filter\":[],\"query\":{\"query\":\"type: tf-client-test\",\"language\":\"lucene\"},\"source\":{\"excludes\":[]},\"highlightAll\":true,\"version\":true,\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"
          }
        },
        "type": "search",
        "id": "tf-client-test"
      }`
var sourceObj map[string]interface{}
err := json.Unmarshal([]byte(source), &sourceObj)
importReq := kibana_objects.KibanaObjectImportRequest{
                KibanaVersion: "7.2.1",
                Hits:          []map[string]interface{}{map[string]interface{}{
                    "_index":  "logzioCustomerKibanaIndex7",
                    "_type":   "_doc",
                    "_id":     "search:tf-client-test",
                    "_source": sourceObj,
                }},
            }
importRes, err := client.ImportKibanaObject(importReq)
```

To export a Kibana Object:

```go
client, _ := kibana_objects.New(apiToken, apiServerAddress)
exportRes, _ := client.ExportKibanaObject(kibana_objects.KibanaObjectExportRequest{
	                Type: kibana_objects.ExportTypeSearch})
```

| function             | func name                                                                                                                        |
|----------------------|----------------------------------------------------------------------------------------------------------------------------------|
| export kibana object | `func (c *KibanaObjectsClient) ExportKibanaObject(exportRequest KibanaObjectExportRequest) (*KibanaObjectExportResponse, error)` |
| import kibana object | `func (c *KibanaObjectsClient) ImportKibanaObject(importRequest KibanaObjectImportRequest) (*KibanaObjectImportResponse, error)` |
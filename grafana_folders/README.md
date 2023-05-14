# Grafana Folder

To create a Grafana Folder:

```go
	client, err := grafana_folders.New(apiToken, server.URL)
    grafanaFolder, err := client.CreateGrafanaFolder(grafana_folders.CreateUpdateFolder{
                                                    Uid:   "client_test",
                                                    Title: "client_test_title",
                                                    })
	

```

| function             | func name                                                                                               |
|----------------------|---------------------------------------------------------------------------------------------------------|
| create folder        | `func (c *GrafanaFolderClient) CreateGrafanaFolder(payload CreateUpdateFolder) (*GrafanaFolder, error)` |
| update folder        | `func (c *GrafanaFolderClient) UpdateGrafanaFolder(update CreateUpdateFolder) error`                    |
| delete folder by uid | `func (c *GrafanaFolderClient) DeleteGrafanaFolder(uid string) error`                                   |
| get folder by uid    | `func (c *GrafanaFolderClient) GetGrafanaFolder(uid string) (*GrafanaFolder, error)`                    |
| list folders         | `func (c *GrafanaFolderClient) ListGrafanaFolders() ([]GrafanaFolder, error)`                           |
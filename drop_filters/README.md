# Drop Filters
Compatible with Logz.io's [drop filters API](https://api-docs.logz.io/docs/logz/drop-filters).

Drop filters provide a solution for filtering out logs before they are indexed in your account to help lower costs and reduce account volume.
To create a new drop filter:

```go
client, _ := drop_filters.New(apiToken, apiServerAddress)
dropFilter, err := client.CreateDropFilter(drop_filters.CreateDropFilter{
                    LogType: "some_type",
                    FieldConditions: []drop_filters.FieldConditionObject{{
                        FieldName: "some_field_name",
                        Value:     "some_value",
                        ThresholdInGB: 10,
                    }},
                })
```

## API Functions
|function|func name|
|---|---|
|activate drop filter| `func (c *DropFiltersClient) ActivateDropFilter(dropFilterID string) (*DropFilter, error)` |
|create drop filter| `func (c *DropFiltersClient) CreateDropFilter(createDropFilter CreateDropFilter) (*DropFilter, error)` |
|deactivate drop filter| `func (c *DropFiltersClient) DeactivateDropFilter(dropFilterId string) (*DropFilter, error)` |
|delete drop filter| `func (c *DropFiltersClient) DeleteDropFilter(dropFilterId string) error` |
|retrieve drop filters| `func (c *DropFiltersClient) RetrieveDropFilters() ([]DropFilter, error)` |

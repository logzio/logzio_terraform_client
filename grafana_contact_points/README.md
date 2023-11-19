# Grafana Contact Point

To create a Grafana Contact Point:

```go
    createGrafanaContactPoint := grafana_contact_points.GrafanaContactPoint{
                                    Name: "my-contact-point",
                                    Type: "email",
                                    Settings: map[string]interface{}{
                                            "addresses":   "example1@example.com;example2@example.com",
                                            "singleEmail": false,
                                        },
                                    DisableResolveMessage: false,
                                }	
    client, err := grafana_contact_points.New("some-token", "some-url")
    contactPoint, err := client.CreateGrafanaContactPoint(createGrafanaContactPoint)
```

| function                   | func name                                                                                                                 |
|----------------------------|---------------------------------------------------------------------------------------------------------------------------|
| create contact point       | `func (c *GrafanaContactPointClient) CreateGrafanaContactPoint(payload GrafanaContactPoint) (GrafanaContactPoint, error)` |
| delete contact point       | `func (c *GrafanaContactPointClient) DeleteGrafanaContactPoint(uid string) error`                                         |
| get all contact points     | `func (c *GrafanaContactPointClient) GetAllGrafanaContactPoints() ([]GrafanaContactPoint, error)`                         |
| get contact point by uid   | `func (c *GrafanaContactPointClient) GetGrafanaContactPointByUid(uid string) (GrafanaContactPoint, error)`                |
| get contact points by name | `func (c *GrafanaContactPointClient) GetGrafanaContactPointsByName(name string) ([]GrafanaContactPoint, error) `          |
| update contact point       | `func (c *GrafanaContactPointClient) UpdateContactPoint(contactPoint GrafanaContactPoint) error`                          |
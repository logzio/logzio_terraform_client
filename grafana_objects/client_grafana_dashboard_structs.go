package grafana_objects

type DashboardMeta struct {
	IsStarred bool   `json:"isStarred"`
	Url       string `json:"url"`
	FolderId  int    `json:"folderId"`
	FolderUid string `json:"folderUid"`
	Slug      string `json:"slug"`
}

type DashboardObject struct {
	Id            int                      `json:"id"`
	Uid           string                   `json:"uid"`
	Title         string                   `json:"title"`
	Tags          []string                 `json:"tags"`
	Style         string                   `json:"style"`
	Timezone      string                   `json:"timezone"`
	Editable      bool                     `json:"editable"`
	GraphToolTip  int                      `json:"graphTooltip"`
	Time          TimeRange                `json:"time"`
	Timepicker    Timepicker               `json:"timepicker"`
	Templating    TemplatingList           `json:"templating"`
	Annotations   interface{}              `json:"annotations"`
	Refresh       string                   `json:"refresh"`
	SchemaVersion int                      `json:"schemaVersion"`
	Version       int                      `json:"version"`
	Panels        []map[string]interface{} `json:"panels"`
}

type TimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Timepicker struct {
	Collapse         bool     `json:"collapse"`
	Enable           bool     `json:"enable"`
	Notice           bool     `json:"notice"`
	Now              bool     `json:"now"`
	RefreshIntervals []string `json:"refresh_intervals"`
	Status           string   `json:"status"`
	Type             string   `json:"type"`
}

type TemplatingList struct {
	Enable bool             `json:"enable"`
	List   []TemplatingItem `json:"list"`
}

type TemplatingItem struct {
	AllFormat   string                 `json:"string"`
	Current     map[string]interface{} `json:"current"`
	Datasource  string                 `json:"datasource"`
	IncludeAll  bool                   `json:"includeAll"`
	Multi       bool                   `json:"multi"`
	MultiFormat string                 `json:"multiFormat"`
	Name        string                 `json:"name"`
	Options     []TemplatingOptions    `json:"options"`
	Type        string                 `json:"type"`
}

type TemplatingOptions struct {
	Selected bool   `json:"selected"`
	Text     string `json:"test"`
	Value    string `json:"value"`
}

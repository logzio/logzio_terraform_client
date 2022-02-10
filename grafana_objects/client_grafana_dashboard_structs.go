package grafana_objects

type DashboardMeta struct {
	IsStarred bool   `json:"isStarred"`
	Url       string `json:"url"`
	FolderId  int    `json:"folderId"`
	FolderUid string `json:"folderUid"`
	Slug      string `json:"slug"`
}

type DashboardObject struct {
	Id            int            `json:"id,omitempty"`
	Uid           string         `json:"uid"`
	GnetId        *string        `json:"gnetId"`
	Title         string         `json:"title"`
	Tags          []string       `json:"tags"`
	Style         string         `json:"style"`
	Iteration     int            `json:"iteration"`
	Links         []string       `json:"links"`
	Timezone      string         `json:"timezone"`
	Editable      bool           `json:"editable"`
	GraphToolTip  int            `json:"graphTooltip"`
	Time          TimeRange      `json:"time"`
	Timepicker    *Timepicker    `json:"timepicker,omitempty"`
	Templating    TemplatingList `json:"templating"`
	Annotations   interface{}    `json:"annotations"`
	Refresh       string         `json:"refresh"`
	SchemaVersion int            `json:"schemaVersion"`
	Version       int            `json:"version,omitempty"`
	Panels        []interface{}  `json:"panels"`
}

type TimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Timepicker struct {
	Collapse         bool     `json:"collapse,omitempty"`
	Enable           bool     `json:"enable,omitempty"`
	Notice           bool     `json:"notice,omitempty"`
	Now              bool     `json:"now,omitempty"`
	RefreshIntervals []string `json:"refresh_intervals,omitempty"`
	Status           string   `json:"status,omitempty"`
	Type             string   `json:"type,omitempty"`
}

type TemplatingList struct {
	Enable bool             `json:"enable"`
	List   []TemplatingItem `json:"list"`
}

type TemplatingItem struct {
	AllFormat   *string                `json:"string"`
	Current     map[string]interface{} `json:"current"`
	Datasource  *string                `json:"datasource"`
	Definition  *string                `json:"definition"`
	Description *string                `json:"description"`
	IncludeAll  bool                   `json:"includeAll"`
	AllValue    *string                `json:"allValue"`
	Multi       bool                   `json:"multi"`
	Name        *string                `json:"name"`
	Options     []TemplatingOptions    `json:"options"`
	Type        *string                `json:"type"`
	Query       Query                  `json:"query"`
	Error       *string                `json:"error"`
	Hide        int                    `json:"hide"`
	Label       *string                `json:"label"`
	Refresh     int                    `json:"refresh"`
	Regex       *string                `json:"regex"`
	SkipUrlSync bool                   `json:"skipUrlSync"`
	Sort        int                    `json:"sort"`
	String      *string                `json:"string"`
}

type TemplatingOptions struct {
	Selected bool   `json:"selected"`
	Text     string `json:"text"`
	Value    string `json:"value"`
}

type Query struct {
	Query string `json:"query"`
	RefId string `json:"refId"`
}

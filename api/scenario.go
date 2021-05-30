package api

type NewScenario struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"isActive"`
}

type BriefScenario struct {
	ID                uint     `json:"id"`
	Name              string   `json:"name"`
	Interval          int32    `json:"interval"`
	LastScheduledTime string   `json:"lastScheduledTime"`
	IsActive          *bool    `json:"isActive"`
	DigestNodes       []string `json:"nodes"`
}

type FullScenario struct {
	ID                uint          `json:"id"`
	Name              string        `json:"name"`
	IsActive          *bool         `json:"isActive"`
	LastScheduledTime string        `json:"lastScheduledTime"`
	Interval          int32         `json:"interval"`
	Services          []ServiceNode `json:"services"`
}

type NewScenarioInterval struct {
	Interval int32 `json:"interval"`
}

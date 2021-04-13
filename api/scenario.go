package api

type NewScenario struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"isActive"`
}

type BriefScenario struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	IsActive    *bool    `json:"isActive"`
	DigestNodes []string `json:"nodes"`
}

type FullScenarioGraph struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	IsActive *bool    `json:"isActive"`
	Graph    *Service `json:"graph"`
}

type FullScenario struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	IsActive *bool     `json:"isActive"`
	Interval int       `json:"interval"`
	Services []Service `json:"services"`
}

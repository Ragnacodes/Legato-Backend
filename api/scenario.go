package api

type NewScenario struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"`
}

type BriefScenario struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	IsActive    *bool    `json:"is_active"`
	DigestNodes []string `json:"nodes"`
}

type FullScenarioGraph struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	IsActive *bool    `json:"is_active"`
	Graph    *Service `json:"graph"`
}

type FullScenario struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	IsActive *bool     `json:"is_active"`
	Services []Service `json:"services"`
}

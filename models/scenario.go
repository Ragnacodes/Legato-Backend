package models

type NewScenario struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

package api

type HistoryInfo struct {
	ID         uint   `json:"id"`
	CreatedAt  string `json:"created_at"`
	ScenarioId uint   `json:"scenario_id"`
}

type RecentUserHistory struct {
	Scenario    BriefScenario `json:"scenario"`
	HistoryInfo HistoryInfo   `json:"history"`
}

type ServiceLogInfo struct {
	Id        int `json:"id"`
	Messages  []MessageInfo
	Service   ServiceNode
	CreatedAt string `json:"created_at"`
}

type MessageInfo struct {
	CreatedAt string `json:"created_at"`
	Data      string `json:"context"`
	Type      string `json:"type"`
}

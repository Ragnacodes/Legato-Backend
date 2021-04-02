package models

type Service struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Children []Service   `json:"children"`
	Data     interface{} `json:"data"`
}

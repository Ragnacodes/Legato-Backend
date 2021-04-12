package api

type Service struct {
	Id       uint        `json:"id"`
	ParentId *uint        `json:"parentId"`
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Children []Service   `json:"children"`
	Position Position    `json:"position"`
	Data     interface{} `json:"data"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

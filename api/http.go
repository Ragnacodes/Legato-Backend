package api

type HttpInfo struct {
	Id     uint   `json:"id"`
	Url    string `json:"url"`
	Method string `json:"method"`
}

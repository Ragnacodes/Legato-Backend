package api

type HistoryInfo struct {
	ID            uint     `json:"id"`
	CreatedAt     string   `json:"created_at"`
}


type ServiceLogInfo struct {
	Id	      int  			`json:"id"`
	Messages   []MessageInfo 
	Service	   ServiceNode
	CreatedAt string 		`json:"created_at"`
}

type MessageInfo struct{
	CreatedAt string `json:"created_at"`
	Data	  string `json:"context"`
	Type	  string `json:"type"`
}
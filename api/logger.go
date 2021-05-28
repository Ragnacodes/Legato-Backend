package api

type HistoryInfo struct {
	ID             uint      `json:"id"`
	CreatedAt      string    `json:"created_at"`
}


type ServiceLogInfo struct {
	Status	   int  `json:"id"`
	Messages   []MessageInfo 
	Service	   ServiceNode
}

type MessageInfo struct{
	Data	string `json:"context"`
}
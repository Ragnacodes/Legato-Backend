package logging

import (
	"fmt"
	"log"
	"github.com/alexandrevicenzi/go-sse"
)

type serverEvent struct{
    EventServer *sse.Server
}

var (
    SSE serverEvent
)

func (s serverEvent) Init() {
    // Create SSE server
    s.EventServer = sse.NewServer(nil)
}

func (s serverEvent) SendEvent(data string, scid uint){
    channel := fmt.Sprintf("api/events/%v", scid)
    go s.EventServer.SendMessage(channel, sse.SimpleMessage(data))
    log.Printf(data)
}

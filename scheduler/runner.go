package scheduler

import (
	"context"
	"log"
)

func Listen() {
	c := context.Background()

	err := QueueFactory.StartConsumers(c)
	if err != nil {
		panic(err)
	}

	sig := WaitSignal()
	log.Println(sig.String())

	err = QueueFactory.Close()
	if err != nil {
		panic(err)
	}
}

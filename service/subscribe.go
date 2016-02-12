package service

import (
	"encoding/json"
	"log"
	"time"

	ctrl "../controller"
)

type subscribe struct {
	hub *ctrl.Hub
}

// Constructor.
func NewSubscribe(h *ctrl.Hub) *subscribe {
	return &subscribe{h}
}

// Subscribe to topic.
func (this *subscribe) Subscribe(topics map[string]func(...string)[]byte) {
	for topic, _ := range topics {
		sub := &ctrl.WSMessage{"SUBSCRIBE", topic, ""}
		j, err := json.Marshal(sub)
		if err != nil {
			log.Println(err)
			break
		}
		this.hub.Send([]byte(j))
		time.Sleep(time.Millisecond * 500)
	}
}

package controller

import (
	"encoding/json"
	"log"
	"time"

	"service"
)

type subscribe struct {
	hub *Hub
}

// Constructor.
func NewSubscribe(h *Hub) *subscribe {
	return &subscribe{h}
}

// Subscribe to topic.
func (this *subscribe) Subscribe(topics service.TopicPool) {
	for topic, _ := range topics {
		sub := &WSMessage{"SUBSCRIBE", topic, ""}
		j, err := json.Marshal(sub)
		if err != nil {
			log.Println(err)
			break
		}
		this.hub.Send([]byte(j))
		time.Sleep(time.Millisecond * 500)
	}
}

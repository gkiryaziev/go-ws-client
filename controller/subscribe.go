package controller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gkiryaziev/go-ws-client/service"
)

type subscribe struct {
	hub *Hub
}

// NewSubscribe return new subscribe object.
func NewSubscribe(h *Hub) *subscribe {
	return &subscribe{h}
}

// Subscribe to topic.
func (s *subscribe) Subscribe(topics service.TopicPool) {
	for topic, _ := range topics {
		sub := &WSMessage{"SUBSCRIBE", topic, ""}
		j, err := json.Marshal(sub)
		if err != nil {
			log.Println(err)
			break
		}
		s.hub.Send([]byte(j))
		time.Sleep(time.Millisecond * 500)
	}
}

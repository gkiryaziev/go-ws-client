package controller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gkiryaziev/go-ws-client/service"
)

// Subscribe struct
type Subscribe struct {
	hub *Hub
}

// NewSubscribe return new Subscribe object.
func NewSubscribe(h *Hub) *Subscribe {
	return &Subscribe{h}
}

// Subscribe to topic.
func (s *Subscribe) Subscribe(topics service.TopicPool) {
	for topic := range topics {
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

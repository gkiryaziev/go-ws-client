package controller

import (
	"encoding/json"
	"log"
)

type WSMessage struct {
	Action string `json:"action"`
	Topic  string `json:"topic"`
	Data   string `json:"data"`
}

// Get json message
func GetMessage(topic, data string) []byte {
	pub := &WSMessage{"PUBLISH", topic, data}
	j, err := json.Marshal(pub)
	if err != nil {
		log.Println(err)
		return nil
	}
	return j
}

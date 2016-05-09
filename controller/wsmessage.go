package controller

import (
	"encoding/json"
	"log"
)

// WSMessage struct
type WSMessage struct {
	Action string `json:"action"`
	Topic  string `json:"topic"`
	Data   string `json:"data"`
}

// GetMessage return message as json
func GetMessage(topic, data string) []byte {
	pub := &WSMessage{"PUBLISH", topic, data}
	j, err := json.Marshal(pub)
	if err != nil {
		log.Println(err)
		return nil
	}
	return j
}

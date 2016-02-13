package raspberry

import (
	"encoding/json"
	"log"

	ctrl "../controller"
)

// Get json message
func getMessage(topic, data string) []byte {
	pub := &ctrl.WSMessage{"PUBLISH", topic, data}
	j, err := json.Marshal(pub)
	if err != nil {
		log.Println(err)
		return nil
	}
	return j
}

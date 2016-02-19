package raspberry

import (
	"encoding/json"
	"log"

	ctrl "controller"
	cmd "command"
)

type raspberry struct {
	command *cmd.Command
}

func NewRaspberry() *raspberry {
	return &raspberry{cmd.NewCommand()}
}

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

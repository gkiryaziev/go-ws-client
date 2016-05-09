package raspberry

import (
	"encoding/json"
	"log"

	cmd "github.com/gkiryaziev/go-ws-client/command"
	"github.com/gkiryaziev/go-ws-client/command/memory"
	ctrl "github.com/gkiryaziev/go-ws-client/controller"
)

// SystemMemory return system memory.
func (r *Raspberry) SystemMemory(data string) []byte {

	sysMem := memory.Clean(cmd.Exec("cat", "/proc/meminfo"), "MemTotal:", "MemFree:", "MemAvailable:")
	if sysMem == nil {
		return nil
	}
	jSysMem, err := json.Marshal(sysMem)
	if err != nil {
		log.Println(err)
		return nil
	}

	return ctrl.GetMessage("RPI1_SYS_MEM", string(jSysMem))
}

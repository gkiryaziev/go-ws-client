package raspberry

import (
	"encoding/json"
	"log"

	cmd "command"
	"command/memory"
	ctrl "controller"
)

// Get system memory.
func (this *raspberry) SystemMemory(data string) []byte {

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

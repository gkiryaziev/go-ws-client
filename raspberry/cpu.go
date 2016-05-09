package raspberry

import (
	cmd "github.com/gkiryaziev/go-ws-client/command"
	"github.com/gkiryaziev/go-ws-client/command/cpu"
	ctrl "github.com/gkiryaziev/go-ws-client/controller"
)

// CPUTemp return cpu temp.
func (r *Raspberry) CPUTemp(data string) []byte {
	cpuTemp := cpu.Clean(cmd.Exec("vcgencmd", "measure_temp"), "temp=", "'C")
	if cpuTemp == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_TEMP", cpuTemp)
}

// CPUMemory return cpu memory.
func (r *Raspberry) CPUMemory(data string) []byte {

	cpuMem := cpu.Clean(cmd.Exec("vcgencmd", "get_mem", "arm"), "arm=", "M")
	if cpuMem == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_MEM", cpuMem)
}

// CPUCoreVolt return core volt.
func (r *Raspberry) CPUCoreVolt(data string) []byte {

	cpuCoreVolt := cpu.Clean(cmd.Exec("vcgencmd", "measure_volts", "core"), "volt=", "V")
	if cpuCoreVolt == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_CORE_VOLT", cpuCoreVolt)
}

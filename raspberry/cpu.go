package raspberry

import (
	cmd "github.com/gkiryaziev/go-ws-client/command"
	"github.com/gkiryaziev/go-ws-client/command/cpu"
	ctrl "github.com/gkiryaziev/go-ws-client/controller"
)

// CpuTemp return cpu temp.
func (r *raspberry) CpuTemp(data string) []byte {
	cpuTemp := cpu.Clean(cmd.Exec("vcgencmd", "measure_temp"), "temp=", "'C")
	if cpuTemp == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_TEMP", cpuTemp)
}

// CpuMemory return cpu memory.
func (r *raspberry) CpuMemory(data string) []byte {

	cpuMem := cpu.Clean(cmd.Exec("vcgencmd", "get_mem", "arm"), "arm=", "M")
	if cpuMem == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_MEM", cpuMem)
}

// CpuCoreVolt return core volt.
func (r *raspberry) CpuCoreVolt(data string) []byte {

	cpuCoreVolt := cpu.Clean(cmd.Exec("vcgencmd", "measure_volts", "core"), "volt=", "V")
	if cpuCoreVolt == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_CORE_VOLT", cpuCoreVolt)
}

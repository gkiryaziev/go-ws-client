package raspberry

import (
	cmd "rpi.ws.client/command"
	"rpi.ws.client/command/cpu"
	ctrl "rpi.ws.client/controller"
)

// Get cpu temp.
func (this *raspberry) CpuTemp(data string) []byte {
	cpuTemp := cpu.Clean(cmd.Exec("vcgencmd", "measure_temp"), "temp=", "'C")
	if cpuTemp == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_TEMP", cpuTemp)
}

// Get cpu memory.
func (this *raspberry) CpuMemory(data string) []byte {

	cpuMem := cpu.Clean(cmd.Exec("vcgencmd", "get_mem", "arm"), "arm=", "M")
	if cpuMem == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_MEM", cpuMem)
}

// Get core volt.
func (this *raspberry) CpuCoreVolt(data string) []byte {

	cpuCoreVolt := cpu.Clean(cmd.Exec("vcgencmd", "measure_volts", "core"), "volt=", "V")
	if cpuCoreVolt == "" {
		return nil
	}
	return ctrl.GetMessage("RPI1_CPU_CORE_VOLT", cpuCoreVolt)
}

package raspberry

// Get cpu temp.
func (this *raspberry) CpuTemp(data string) []byte {
	temp := this.command.Command("vcgencmd").Arg("measure_temp").Clean("temp=").Clean("'C").Run()
	if temp == "" {
		return nil
	}
	return getMessage("RPI1_CPU_TEMP", temp)
}

// Get cpu memory.
func (this *raspberry) CpuMemory(data string) []byte {
	mem := this.command.Command("vcgencmd").Arg("get_mem").Arg("arm").Clean("arm=").Clean("M").Run()
	if mem == "" {
		return nil
	}
	return getMessage("RPI1_CPU_MEM", mem)
}
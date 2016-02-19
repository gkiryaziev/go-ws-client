package raspberry

// Get cpu temp.
func (this *raspberry) CpuTemp(data string) []byte {
	temp := this.command.Command("go").Arg("run").Arg("temp.go").Clean("temp=").Clean("'C").Run()
	if temp == "" {
		return nil
	}
	return getMessage("RPI1_CPU_TEMP", temp)
}

// Get cpu memory.
func (this *raspberry) CpuMemory(data string) []byte {
	mem := this.command.Command("go").Arg("run").Arg("mem.go").Clean("arm=").Clean("M").Run()
	if mem == "" {
		return nil
	}
	return getMessage("RPI1_CPU_MEM", mem)
}
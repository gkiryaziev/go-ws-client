package raspberry

// Get cpu temp.
func CpuTemp(data string) []byte {
	pub := getMessage("RPI1_CPU_TEMP", "38.2")
	return pub
}

// Get cpu memory.
func CpuMemory(data string) []byte {
	pub := getMessage("RPI1_CPU_MEM", "960")
	return pub
}

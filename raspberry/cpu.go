package raspberry

// Get cpu temp.
func CpuTemp(s ...string) []byte {
	return []byte("{\"action\" : \"PUBLISH\", \"topic\" : \"RPI1_CPU_TEMP\", \"data\" : \"38.2\"}")
}

// Get cpu memory.
func CpuMemory(s ...string) []byte {
	return []byte("{\"action\" : \"PUBLISH\", \"topic\" : \"RPI1_CPU_MEM\", \"data\" : \"960\"}")
}

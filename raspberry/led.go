package raspberry

// Get or Set LED0.
func Led0(s ...string) []byte {
	return []byte("{\"action\" : \"PUBLISH\", \"topic\" : \"RPI1_LED0\", \"data\" : \"OFF\"}")
}
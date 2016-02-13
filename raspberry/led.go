package raspberry

// Get or Set LED0.
func Led0(data string) []byte {
	pub := getMessage("RPI1_LED0", "OFF")
	return pub
}

package raspberry

// Get or Set LED0.
func (this *raspberry) Led0(data string) []byte {
	pub := getMessage("RPI1_LED0", "OFF")
	return pub
}

package raspberry

import (
	ctrl "controller"
)

// Get or Set LED0.
func (this *raspberry) Led0(data string) []byte {
	pub := ctrl.GetMessage("RPI1_LED0", "OFF")
	return pub
}

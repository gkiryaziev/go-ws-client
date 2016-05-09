package raspberry

import (
	ctrl "github.com/gkiryaziev/go-ws-client/controller"
)

// Led0 get or set LED0.
func (r *Raspberry) Led0(data string) []byte {
	pub := ctrl.GetMessage("RPI1_LED0", "OFF")
	return pub
}

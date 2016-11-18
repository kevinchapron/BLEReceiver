package main

import (
	"fmt"
	"github.com/paypal/gatt"
	"github.com/kevinchapron/BLEReceiver/handlers"
	"github.com/kevinchapron/BLEReceiver"
	"github.com/kevinchapron/BLEReceiver/settings"
)


func main() {
	var device gatt.Device = BLEReceiver.InitDevice()

	device.Handle(
		gatt.PeripheralDiscovered(handlers.Discovering),
		gatt.PeripheralConnected(handlers.Connecting),
		gatt.PeripheralDisconnected(handlers.Disconnecting),
	)

	device.Init(handlers.StateChanged)
	<-settings.Done
	fmt.Printf("End of program\n")
}

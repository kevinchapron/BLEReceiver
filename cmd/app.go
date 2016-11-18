package main

import (
	"github.com/kevinchapron/BLEReceiver"
	"github.com/kevinchapron/BLEReceiver/handlers"
	"github.com/paypal/gatt"
	"fmt"
)

var done = make(chan struct{})

func main() {
	var device gatt.Device = BLEReceiver.InitDevice()

	device.Handle(
		gatt.PeripheralDiscovered(handlers.Discovering),
		gatt.PeripheralConnected(handlers.Connecting),
		gatt.PeripheralDisconnected(handlers.Disconnecting),
	)

	device.Init(handlers.StateChanged)
	<-done
	fmt.Printf("End of program\n")
}

package handlers

import (
	"fmt"
	"github.com/paypal/gatt"
)

func StateChanged(d gatt.Device, s gatt.State){
	fmt.Printf("State changed to %s\n",s.String())

	switch s {
		case gatt.StatePoweredOn:
			fmt.Println("Scanning...")
			d.Scan([]gatt.UUID{}, false)
			return
		default:
			d.StopScanning()
	}
}
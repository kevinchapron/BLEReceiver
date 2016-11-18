package handlers

import (
	"github.com/paypal/gatt"
	"github.com/kevinchapron/BLEReceiver"
	"strings"
	"fmt"
)

func Discovering(p gatt.Peripheral, a *gatt.Advertisement, rssi int){
	id_target := strings.ToUpper(BLEReceiver.GetTargetID())
	id_peripheral := strings.ToUpper(p.ID())
	if id_peripheral != id_target{
		fmt.Printf("Discovered Peripheral \"%s\"\n",id_peripheral)
		return
	}

	p.Device().StopScanning()
	fmt.Printf("\nPeripheral ID : %s, NAME : (%s)\n",id_peripheral,p.Name())
	p.Device().Connect(p)
}
package handlers

import (
	"github.com/paypal/gatt"
	"fmt"
	"github.com/kevinchapron/BLEReceiver/settings"
)

func Disconnecting(p gatt.Peripheral, err error){
	fmt.Printf("Disconnected to %s \n",p.ID())
	close(settings.Done)
}
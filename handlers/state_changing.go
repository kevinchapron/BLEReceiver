package handlers

import (
	"fmt"
	"github.com/paypal/gatt"
)

func StateChanged(d gatt.Device, s gatt.State){
	fmt.Printf("State changed to %s\n",s.String())
}
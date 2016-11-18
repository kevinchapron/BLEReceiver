package handlers

import (
	"github.com/paypal/gatt"
	"fmt"
)

func Disconnecting(p gatt.Peripheral, err error){
	fmt.Printf("Disconnected to %s \n",p.ID())
}
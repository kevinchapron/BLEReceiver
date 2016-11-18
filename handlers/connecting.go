package handlers

import (
	"github.com/paypal/gatt"
	"fmt"
)

func Connecting(p gatt.Peripheral, err error){
	fmt.Printf("Connected to %s \n",p.ID())
}
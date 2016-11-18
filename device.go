package BLEReceiver

import (
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"log"
)

func InitDevice() gatt.Device{
	device,err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil{
		log.Fatalf("Failed to init device, err : %s\n",err)
	}
	return device
}
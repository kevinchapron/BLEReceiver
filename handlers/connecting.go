package handlers

import (
	"github.com/paypal/gatt"
	"fmt"
	"github.com/kevinchapron/BLEReceiver/settings"
	"github.com/kevinchapron/BLEReceiver/managers"
)

func Connecting(p gatt.Peripheral, err error){
	fmt.Printf("Connected to %s \n",p.ID())
	mtu_err := p.SetMTU(500);

	if mtu_err != nil {
		fmt.Printf("Failed to set MTU, err: %s\n", mtu_err)
	}

	list_services,err := p.DiscoverServices(nil);
	for _, service := range list_services{
		msg := "Service: " + service.UUID().String()
		if len(service.Name()) > 0 {
			msg += " (" + service.Name() + ")"
		}
		fmt.Println(msg)


		if(!settings.SERVICES_UUID.Has(service.UUID().String())){
			fmt.Printf("--- Service not listed in \"%s\", skipping ...\n","SERVICES_UUID")
			continue
		}
		managers.ManageService(p, service)
	}
}
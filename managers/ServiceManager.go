package managers

import (
	"github.com/paypal/gatt"
	"fmt"
	"github.com/kevinchapron/BLEReceiver/settings"
)

func ManageService(p gatt.Peripheral, s gatt.Service){
	cs, err := p.DiscoverCharacteristics(nil, s)
	if err != nil {
		fmt.Printf("Failed to discover characteristics, err: %s\n", err)
		return
	}

	for _, c := range(cs){
		msg := "  Characteristic  " + c.UUID().String()
		if len(c.Name()) > 0 {
			msg += " (" + c.Name() + ")"
		}

		fmt.Println(msg)

		if(!settings.CHARACTERISTICS_UUID.Has(c.UUID().String())){
			fmt.Printf("--- Characteristic not listed in \"%s\", skipping ...\n","CHARACTERISTICS_UUID")
			continue
		}
		ManageCharacteristic(p, c)
	}
}
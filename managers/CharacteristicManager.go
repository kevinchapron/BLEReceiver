package managers

import (
	"github.com/paypal/gatt"
	"fmt"
)

func ManageCharacteristic(p gatt.Peripheral, c gatt.Characteristic){

	// Discovery descriptors
	ds, err := p.DiscoverDescriptors(nil, c)
	if err != nil {
		fmt.Printf("Failed to discover descriptors, err: %s\n", err)
	}
	ReadDescriptors(p,ds);

	// Subscribe the characteristic, if possible.
	if (c.Properties() & (gatt.CharNotify | gatt.CharIndicate)) != 0 {
		f := func(c *gatt.Characteristic, b []byte, err error) {
			fmt.Printf("notified: % X | %q\n", b, b)
		}
		if err := p.SetNotifyValue(c, f); err != nil {
			fmt.Printf("Failed to subscribe characteristic, err: %s\n", err)
		}
	}

}
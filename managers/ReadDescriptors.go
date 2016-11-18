package managers

import (
	"github.com/paypal/gatt"
	"fmt"
)

func ReadDescriptors(p gatt.Peripheral, descriptors []*gatt.Descriptor){
	for _, d := range descriptors {
		msg := "  Descriptor      " + d.UUID().String()
		if len(d.Name()) > 0 {
			msg += " (" + d.Name() + ")"
		}
		fmt.Println(msg)

		// Read descriptor (could fail, if it's not readable)
		_, err := p.ReadDescriptor(d)
		if err != nil {
			fmt.Printf("Failed to read descriptor, err: %s\n", err)
			continue
		}
	}
}
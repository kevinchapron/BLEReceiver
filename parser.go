package BLEReceiver

import (
	"os"
	"log"
	"flag"
)

func GetTargetID() string{
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalf("Usage: %s ID_PERIPHERAL\n", os.Args[0])
	}
	return os.Args[1]
}
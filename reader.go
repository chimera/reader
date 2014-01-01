package main

import (
	"log"

	"code.google.com/p/gopass"

	"github.com/chimera/auth"
	"github.com/chimera/door"
)

func main() {

	// Create a new door lock instance.
	d := door.NewDoorLock()

	// Wait for input
	for {

		// Prompt for an RFID code.
		code, err := gopass.GetPass("Please input your RFID code for access: ")
		if err != nil {
			log.Fatal(err)
		}

		// If a code is received, send it to get authenticated.
		if code != "" {

			// TODO: Actually authenticate code here...
			ok, err := auth.IsAuthenticated(code)
			if err != nil {
				log.Println(err)
			}

			// The code was authenticated and no errors were raised, open up the door!
			if ok {
				_, err = d.Unlock()
				if err != nil {
					log.Printf("Error unlocking door! %s", err.Error())
				}
			}

		}
	}
}

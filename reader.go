package main

import (

	"code.google.com/p/gopass"
	"github.com/danawoodman/clog"

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
			clog.Error(err.Error())
		}

		// If a code is received, send it to get authenticated.
		if code != "" {

			// TODO: Actually authenticate code here...
			ok, err := auth.IsAuthenticated(code)
			if err != nil {
				log.Println(err)
				clog.Error(err.Error())
			}

			// The code was authenticated and no errors were raised, open up the door!
			if ok {
				_, err = d.Unlock()
				if err != nil {
				}
				clog.Error(fmt.Sprintf("Error unlocking door! %s", err.Error()))
			}

			clog.Success(fmt.Sprintf("Welcome in %s!", user.Name))
		}
	}
}

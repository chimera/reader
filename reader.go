package main

import (
	"flag"
	"fmt"

	"code.google.com/p/gopass"

	"github.com/chimera/auth"
	"github.com/chimera/door"
	"github.com/danawoodman/clog"
)

var path = flag.String("path", "users.json", "The path to the users JSON file.")

func main() {

	flag.Parse()

	// Open up the users database file
	db, err := auth.New(*path)
	if err != nil {
		panic(err)
	}
	clog.Success(fmt.Sprintf("User database `%s` loaded", *path))

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
			user, err := db.FindUser(code)
			if err != nil {
				clog.Error(err.Error())
				continue
			}

			// The code was authenticated and no errors were raised, open up the door!
			err = d.Unlock()
			if err != nil {
				clog.Error(fmt.Sprintf("Error unlocking door! %s", err.Error()))
				continue
			}
			clog.Success(fmt.Sprintf("Welcome in %s!", user.Name))
		}
	}
}

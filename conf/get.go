package conf

import (
	"log"
	"os/user"
)

func getUserPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir + "/"
}

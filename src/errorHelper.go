package src

import (
	"log"
)

func HandleError(err error, msg string) {
	if msg == "" {
		log.Fatal(err.Error())
	} else {
		log.Fatalf("%s :\n\t%s", msg, err.Error())
	}
}

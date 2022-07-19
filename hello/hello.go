package main

import (
	"fmt"
	"log"

	"github.com/olubiyiontheweb/monitoring_alert_tool/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting message
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
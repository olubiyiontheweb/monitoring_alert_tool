package main

import (
	"fmt"
	"log"

	"github.com/olubiyiontheweb/monitoring_alert_tool/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Tosin", "John"}

	// request a greeting message
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
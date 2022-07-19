package greetings

import (
	"errors"
	"fmt"

	"rsc.io/quote/v4"
)

func Hello(name string) (string, error) {
	// if no name was given, return error
	if name == "" {
		return "", errors.New("no name was given")
	}
	fmt.Println("Hello world")
	fmt.Println(quote.Go())
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! We met!",
	}

	return formats[rand.Intn(len(formats))]
}

func Hello(name string) (string, error) {
	// if no name was given, return error
	if name == "" {
		return "", errors.New("no name was given")
	}
	//fmt.Println("Hello world")
	//fmt.Println(quote.Go())
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
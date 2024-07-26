package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(errorMessage())
	}

	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat())
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

func errorMessage() string {
	errorNote := "nope, you didn't enter a name. please try again"
	return errorNote
}

func randomFormat() string {
	formats := [] string {
		"Hi, %v. Welcome!",
		"Great to see you %v",
		"Hey, %v, what do you know?",
	}
	return formats[rand.Intn(len(formats))]
}
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := [] string {
		"Hi, %v. Welcome!",
		"Great to see you %v",
		"Hey, %v, what do you know?",
	}
	return formats[rand.Intn(len(formats))]
}
package main

import (
	"errors"
	"fmt"
)

// ReturnsMessage expects string
func ReturnsMessage(name string) (string, error) {
	if name == "" {
		return "", errors.New("missing name")
	}

	message := fmt.Sprintf("Hello %s ", name)

	return message, nil
}

func main() {
	message, err := ReturnsMessage("dp27664")

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(message)
}

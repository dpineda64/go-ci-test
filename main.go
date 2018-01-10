package main

import (
	"fmt"
	"os"
)

func getUserName() string {
	username := os.Getenv("USER")

	if username != "" {
		username = os.Getenv("USERNAME")
	}

	return username
}

func main() {
	fmt.Println("Hello world %s", getUserName())
}

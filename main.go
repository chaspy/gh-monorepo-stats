package main

import (
	"fmt"
	"log"
)

func run() error {
	// print hello world
	fmt.Println("Hello World")

	return nil
}

func main() {

	err := run()
	if err != nil {
		log.Fatal(err) //nolint:forbidigo
	}
}
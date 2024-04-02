package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	svcURL := os.Getenv("SVC_URL")
	fmt.Println("SVC_URL: ", svcURL)

	for i, arg := range os.Args[1:] {
		fmt.Println("Arg", i, ": ", arg)
	}
}

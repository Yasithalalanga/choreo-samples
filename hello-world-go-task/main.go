package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	for i, arg := range os.Args[1:] {
		fmt.Println("Arg", i, ": ", arg)
	}

	// Print values from environment variables NORMAL_CONFIG & SECRET_CONFIG
	fmt.Println("Printing values from configuration group")
	fmt.Println("Printing environment variables")
	fmt.Println("NORMAL_CONFIG:", os.Getenv("NORMAL_CONFIG"))
	fmt.Println("SECRET_CONFIG:", os.Getenv("SECRET_CONFIG"))

	fmt.Println("Printing file mounts")
	// Read config.json file in /workspace directory
	config, err := os.ReadFile("/workspace/config.json")
	if err != nil {
		fmt.Printf("Error reading config.json: %v\n", err)
	} else {
		fmt.Printf("config.json: %s\n", string(config))
	}

	secretConfig, err := os.ReadFile("/workspace/sample.crt")
	if err != nil {
		fmt.Printf("Error reading sample.crt: %v\n", err)
	} else {
		fmt.Printf("sample.crt: %s\n", string(secretConfig))
	}
}

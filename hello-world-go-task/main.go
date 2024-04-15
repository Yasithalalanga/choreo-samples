package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	fmt.Println("Hello, World!")

	svcURL := os.Getenv("SVC_URL")
	fmt.Println("SVC_URL: ", svcURL)

	consumerKey := os.Getenv("CONSUMER_KEY")
	fmt.Println("CONSUMER_KEY: ", consumerKey)

	consumerSecret := os.Getenv("CONSUMER_SECRET")
	fmt.Println("CONSUMER_SECRET: ", consumerSecret)

	tokenURL := os.Getenv("TOKEN_URL")
	fmt.Println("TOKEN_URL: ", tokenURL)

	var clientCredsConfig = clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     tokenURL,
	}
	client := clientCredsConfig.Client(context.Background())

	//sample go code snippet
	response, err := client.Get(svcURL + "?name=hello")
	if err != nil {
		log.Printf("Error occurred while calling the service: %v", err)
		return
	}

	defer response.Body.Close() // Don't forget to close the body

	// Read the content
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// handle error
		fmt.Println("Error reading the response body:", err)
		return
	}

	fmt.Println("This is connection call from manual task")

	// Convert the body to string and print
	fmt.Printf("Response from the service: %s\n", string(body))

	for i, arg := range os.Args[1:] {
		fmt.Println("Arg", i, ": ", arg)
	}
}

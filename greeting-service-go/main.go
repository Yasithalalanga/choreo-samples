/*
 * Copyright (c) 2023, WSO2 LLC. (https://www.wso2.com/) All Rights Reserved.
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/oauth2/clientcredentials"
)

func main() {

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/greeter/greet", greet)

	serverPort := 9090
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: serverMux,
	}
	go func() {
		log.Printf("Starting HTTP Greeter on port %d\n", serverPort)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP ListenAndServe error: %v", err)
		}
		log.Println("HTTP server stopped serving new requests.")
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh // Wait for shutdown signal

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Shutting down the server...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Shutdown complete.")
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Stranger"
	}

	tokenUrl := os.Getenv("TOKEN_URL")
	serviceUrl := os.Getenv("SERVICE_URL")
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	// clientId, clientSecret and tokenUrl represent variables to which respective environment variables were read
	var clientCredsConfig = clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     tokenUrl,
	}
	client := clientCredsConfig.Client(context.Background())

	//sample go code snippet
	response, err := client.Get(serviceUrl + "?name=hello")
	if err != nil {
		log.Printf("Error occurred while calling the service: %v", err)
		http.Error(w, "Error occurred while calling the service", http.StatusInternalServerError)
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

	// Convert the body to string and print
	fmt.Fprintf(w, "Response from the service: %s\n", string(body))

	fmt.Fprintf(w, "Your service URL is %s, consumer key is %s, consumer secret is %s and token URL is %s\n", name, serviceUrl, consumerKey, consumerSecret, tokenUrl)
}

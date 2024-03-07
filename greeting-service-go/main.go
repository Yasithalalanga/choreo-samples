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
	fmt.Fprintf(w, "Hello, %s!\n", name)

	// Get all the environment variables
	tokenUrl := os.Getenv("TOKEN_URL")
	serviceUrl := os.Getenv("SERVICE_URL")
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	fmt.Printf("All the connection details: %s, %s, %s, %s\n", tokenUrl, serviceUrl, consumerKey, consumerSecret)

	// Get token calling the token endpoint
	config := &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     tokenUrl,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		fmt.Printf("Error getting token: %v\n", err)
		return
	}

	accessToken := token.AccessToken
	fmt.Printf("Access Token: %s\n", accessToken)

	// Append endpoint path to the service URL
	serviceUrl = serviceUrl + "/greeter/greet"

	// Call the service endpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", serviceUrl, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", accessToken)},
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error calling service: %v\n", err)
		return
	}

	defer resp.Body.Close()
	fmt.Printf("Response from service: %v\n", resp)
}

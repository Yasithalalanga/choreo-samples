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

	// Read environment variables
	currency := os.Getenv("CURRENCY")
	apiURL := os.Getenv("STRIPE_API_URL")
	secretKey := os.Getenv("STRIPE_SECRET_KEY")

	fmt.Fprintln(w, "\n\nRead as Environment variables:")
	fmt.Fprintln(w, "Currency:", currency)
	fmt.Fprintln(w, "Stripe API URL:", apiURL)
	fmt.Fprintln(w, "Stripe Secret Key:", secretKey)

	// Read from files
	fmt.Fprintln(w, "\n\nRead from files:")
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Fprintln(w, "Error reading config file: ", err)
		return
	}
	defer configFile.Close()

	buf := make([]byte, 1024)
	n, err := configFile.Read(buf)
	if err != nil {
		fmt.Fprintln(w, "Error reading config file: ", err)
		return
	}

	fmt.Fprintln(w, "Config file content: ", string(buf[:n]))

	certFile, err := os.Open("cert.pem")
	if err != nil {
		fmt.Fprintln(w, "Error reading cert file: ", err)
		return
	}
	defer certFile.Close()

	buf = make([]byte, 1024)
	n, err = certFile.Read(buf)
	if err != nil {
		fmt.Fprintln(w, "Error reading cert file: ", err)
		return
	}

	fmt.Fprintln(w, "Cert file content: ", string(buf[:n]))

	// Redis Sample
	isRedisEnabled := os.Getenv("REDIS_ENABLED")
	if isRedisEnabled == "true" {
		// Connect to Redis
		fmt.Fprintln(w, "\n\nRedis is enabled. Connecting to Redis...")
		redisHost := os.Getenv("REDIS_HOST")
		redisPort := os.Getenv("REDIS_PORT")
		fmt.Fprintf(w, "Redis Host: %s\n", redisHost)
		fmt.Fprintf(w, "Redis Port: %s\n", redisPort)

		// Read from files
		redisConfigFile, err := os.Open("etc/redis/redis.conf")
		if err != nil {
			fmt.Fprintln(w, "Error reading redis config file: ", err)
			return
		}

		defer redisConfigFile.Close()

		buf = make([]byte, 1024)
		content, err := redisConfigFile.Read(buf)
		if err != nil {
			fmt.Fprintln(w, "Error reading redis config file: ", err)
			return
		}

		fmt.Fprintln(w, "\nRedis Config file content:\n\n", string(buf[:content]))
	}

}

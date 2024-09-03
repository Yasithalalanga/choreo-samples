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
	// Create two servers with different ports and handlers
	startServer(9091)
	startServer(9092)

	// Wait for shutdown signal
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh

	// Shutdown both servers
	log.Println("Shutting down servers...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	shutdownServer(shutdownCtx, 9091)
	shutdownServer(shutdownCtx, 9092)
	log.Println("Shutdown complete.")
}

func startServer(port int) {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/greeter/greet", func(w http.ResponseWriter, r *http.Request) {
		greet(w, r, port)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: serverMux,
	}

	go func() {
		log.Printf("Starting HTTP Greeter on port %d\n", port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP ListenAndServe error on port %d: %v", port, err)
		}
		log.Printf("HTTP server on port %d stopped serving new requests.\n", port)
	}()
}

func shutdownServer(ctx context.Context, port int) {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error on port %d: %v", port, err)
	}
	log.Printf("Server on port %d shut down successfully.\n", port)
}

func greet(w http.ResponseWriter, r *http.Request, port int) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Stranger"
	}
	fmt.Fprintf(w, "Hello, %s! You are connected to port %d.\n", name, port)
}

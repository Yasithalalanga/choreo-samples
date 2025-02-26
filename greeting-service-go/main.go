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
	log.Println("Initializing service and loading configurations...")
	loadConfigurations()

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

func loadConfigurations() {
	// Read environment variables
	log.Println("Loaded Environment Variables:")
	log.Printf("CURRENCY: %s", os.Getenv("CURRENCY"))
	log.Printf("STRIPE_API_URL: %s", os.Getenv("STRIPE_API_URL"))
	log.Printf("STRIPE_SECRET_KEY: %s", os.Getenv("STRIPE_SECRET_KEY"))

	// Read from files
	log.Println("\nReading Configuration Files:")
	readAndPrintFile("config.json")
	readAndPrintFile("cert.pem")

	// Redis Sample
	if os.Getenv("REDIS_ENABLED") == "true" {
		log.Println("\nRedis is enabled. Loading Redis configurations...")
		log.Printf("Redis Host: %s", os.Getenv("REDIS_HOST"))
		log.Printf("Redis Port: %s", os.Getenv("REDIS_PORT"))
		readAndPrintFile("etc/redis/redis.conf")
	}
}

func readAndPrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error reading %s: %v", filename, err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		log.Printf("Error reading %s: %v", filename, err)
		return
	}

	log.Printf("Contents of %s:\n%s\n", filename, string(buf[:n]))
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Stranger"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

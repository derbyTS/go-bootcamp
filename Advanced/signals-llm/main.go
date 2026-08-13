package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Add a handler that simulates 5 seconds of work
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Server] Started processing a 5-second request...")
		time.Sleep(5 * time.Second)
		fmt.Fprintln(w, "Work finished successfully!")
		fmt.Println("[Server] Finished processing request.")
	})

	server := &http.Server{Addr: ":8080"}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Server listening on :8080...")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	sig := <-stop
	fmt.Printf("\nReceived %v signal. Starting 10s graceful shutdown...\n", sig)

	// Context gives active requests up to 10s to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Forced shutdown error: %v\n", err)
	}

	fmt.Println("Server stopped cleanly.")
}

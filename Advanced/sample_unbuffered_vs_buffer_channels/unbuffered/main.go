// Unbuffered Channel: The HTTP Health Check (Strict Gateway)
// The Scenario
// You are building an API gateway or a microservice. Before processing a massive user request, your service must fetch a fresh security token or perform a vital health check against a database. If the health check fails or times out, you must abort the user's request immediately.
//
// Why Unbuffered?
// You need guaranteed synchronization and immediate cancellation. If the background worker fails to get the token, the main request thread needs to know right now so it can return an HTTP 500 error to the user. There is no room for "buffering" old tokens or delaying the failure signal.

package main

import (
	"context"
	"fmt"
	"time"
)

func fetchSecurityToken(ctx context.Context, resultChan chan string, errChan chan error) {
	// Simulate an API call to an Auth Provider (e.g., Auth0, Keycloak)
	time.Sleep(150 * time.Millisecond)

	// If everything is fine:
	resultChan <- "super-secret-token-123"

	// If it failed, we would do: errChan <- errors.New("auth failed")
}

func handleUserRequest() {
	// We want to give the auth check a strict 100ms deadline
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Unbuffered channels: We need an instant, blocking hand-off
	resultChan := make(chan string)
	errChan := make(chan error)

	go fetchSecurityToken(ctx, resultChan, errChan)

	// Wait to see who wins the race: the token delivery or the timeout clock
	select {
	case token := <-resultChan:
		fmt.Printf("HTTP 200: Request processed successfully using token: %s\n", token)
	case err := <-errChan:
		fmt.Printf("HTTP 401: Unauthorized: %v\n", err)
	case <-ctx.Done():
		fmt.Println("HTTP 504: Gateway Timeout! Auth took too long.")
	}
}

func main() {
	handleUserRequest()
}

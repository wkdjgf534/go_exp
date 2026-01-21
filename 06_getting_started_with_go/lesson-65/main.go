package main

import (
	"context"
	"fmt"
	"time"
)

func SimulateLongRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled, task stopping...")
			return
		default:
			fmt.Println("Simulating work...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure cancellation happens when main exits

	// Launch the long-running task in a goroutine
	go SimulateLongRunningTask(ctx)

	// Simulate some main thread work for 2 seconds
	time.Sleep(2 * time.Second)

	// Cancel the context after some time
	fmt.Println("Cancelling context...")
	cancel()

	// Wait for the goroutine to finish (although it should be cancelled by now)
	time.Sleep(1 * time.Second)
}

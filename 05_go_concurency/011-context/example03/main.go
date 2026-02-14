package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: context canceled with error: %v, Exiting...\n", id, context.Cause(ctx))
			return
		default:
			fmt.Printf("worker %d: working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	rootContext := context.Background()
	ctx, cancel := context.WithTimeout(rootContext, 5*time.Second)
	defer cancel()

	//time.Sleep(1 * time.Second)
	//cancel()
	//worker 2: context canceled with error: context canceled, Exiting...
	//worker 1: context canceled with error: context canceled, Exiting...

	go worker(ctx, 1)
	go worker(ctx, 2)

	fmt.Scanln()
}

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
	timeoutCtx, _ := context.WithTimeout(rootContext, 6*time.Second)
	deadlineCtx, _ := context.WithDeadline(timeoutCtx, time.Now().Add(3*time.Second))
	valueCtx := context.WithValue(timeoutCtx, "key", "the value from context")

	go worker(timeoutCtx, 1)
	go worker(deadlineCtx, 2)

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Printf("anonymous function canceled with value: %s...\n", ctx.Value("key"))
	}(valueCtx)

	fmt.Scanln()
}

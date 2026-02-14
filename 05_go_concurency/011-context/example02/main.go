package main

import (
	"context"
	"errors"
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
	ctx, cancel := context.WithCancelCause(rootContext)

	go worker(ctx, 1)
	go worker(ctx, 2)

	go func(canc context.CancelCauseFunc) {
		time.Sleep(5 * time.Second)
		canc(errors.New("error x"))
	}(cancel)

	fmt.Scanln()
}

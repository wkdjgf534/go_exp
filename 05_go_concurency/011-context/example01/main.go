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
			fmt.Printf("worker %d: context canceled with error: %s, Exiting...\n", id, ctx.Err().Error())
			return
		default:
			fmt.Printf("worker %d: working...\n", id)
			time.Sleep(1 * time.Second)
		}

	}
}

func main() {
	rootContext := context.Background()

	ctx, cancel := context.WithCancel(rootContext)
	defer cancel()

	go worker(ctx, 1)
	go worker(ctx, 2)

	go func(cancel context.CancelFunc) {
		time.Sleep(5 * time.Second)
		cancel()
	}(cancel)

	fmt.Scanln()
}

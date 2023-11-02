package main

import (
	"context"
	"fmt"
)

func main() {
	c := make(chan string, 1)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	close(c)
	select {
	case <-ctx.Done():
		fmt.Println("done")
	case c <- "hello":
		fmt.Println("send hello")
	}
}

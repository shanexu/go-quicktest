package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	rl := rate.NewLimiter(100, 2)
	prev := time.Now()
	for i := 0; i < 10; i++ {
		_ = rl.Wait(context.Background())
		now := time.Now()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}

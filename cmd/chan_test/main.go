package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	msgs := make(chan string, 1024*10)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
	forloop:
		for {
			select {
			case msg := <-msgs:
				fmt.Println("send: ", msg)
				time.Sleep(100 * time.Millisecond)
			case <-ctx.Done():
				close(msgs)
				break forloop
			}
		}
		for msg := range msgs {
			fmt.Println("drain: ", msg)
		}
		wg.Done()
	}()

	go func() {
		ticker := time.NewTicker(time.Millisecond * 50)
	forloop:
		for {
			select {
			case t := <-ticker.C:
				msgs <- t.String()
			case <-ctx.Done():
				break forloop
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

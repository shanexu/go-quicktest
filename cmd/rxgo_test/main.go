package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	ch := make(chan rxgo.Item, 100)
	ob := rxgo.FromChannel(ch).Debounce(rxgo.WithDuration(time.Second * 5))
	go func() {
		ch := ob.Observe()
		for i := range ch {
			fmt.Println(time.Now().Format(time.RFC3339), "->", i.V)
		}
	}()

	go func() {
		tick := time.NewTicker(time.Second * 10)
		defer tick.Stop()
		i := 0
		for n := range tick.C {
			i++
			ch <- rxgo.Of(i)
			fmt.Println(n.Format(time.RFC3339), "<-", i)
		}
	}()

	select {}
}

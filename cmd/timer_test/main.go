package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/desertbit/timer"
)

func main() {
	useStd := flag.Bool("use_std", false, "use_std")
	flag.Parse()
	done := make(chan struct{})
	var t iTimer
	var c <-chan time.Time

	if *useStd {
		tmp := time.NewTimer(0)
		c = tmp.C
		t = tmp
	} else {
		tmp := timer.NewTimer(0)
		c = tmp.C
		t = tmp
	}

	time.Sleep(time.Second)
	t.Reset(time.Second * 3)
	go func() {
		for {
			select {
			case x := <-c:
				fmt.Println(x)
			case <-done:
				return
			}
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "":
			done <- struct{}{}
			return
		}
	}
}

type iTimer interface {
	Reset(d time.Duration) bool
}

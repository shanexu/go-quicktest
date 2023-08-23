package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("hello world")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()
	conn, err := (&net.Dialer{}).DialContext(ctx, "tcp", "172.16.87.104:12000")
	fmt.Println(err)
	if err != nil {
		return
	}
	time.Sleep(time.Second * 1)
	fmt.Println(conn.Write([]byte("hello world")))
}

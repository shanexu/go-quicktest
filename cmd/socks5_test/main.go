package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"reflect"

	"golang.org/x/net/proxy"
)

type contextDialer interface {
	Dial(network, addr string) (c net.Conn, err error)
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

func main() {
	proxyUrl, _ := url.Parse("socks5://172.16.87.104:11000")
	dialer, _ := proxy.FromURL(proxyUrl, nil)
	cd, _ := dialer.(contextDialer)
	conn, _ := cd.Dial("tcp", "172.29.150.210:9099")
	fmt.Println(reflect.ValueOf(conn).Elem().FieldByName("Conn"))
	fmt.Println(reflect.TypeOf(conn))
	fmt.Println(conn)
}

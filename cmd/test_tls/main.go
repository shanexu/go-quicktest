package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "baidu.com:443")
	if err != nil {
		panic(err)
	}
	tlsConn := tls.Client(conn, &tls.Config{
		ServerName: "baidu.com",
	})
	err = tlsConn.Handshake()
	if err != nil {
		panic(err)
	}

	_, err = tlsConn.Write([]byte("GET / HTTP/1.1\r\n" +
		"Host: baidu.com\r\n" +
		"User-Agent: curl/8.4.0\r\n" +
		"Accept: */*\r\n" +
		"\r\n",
	))
	if err != nil {
		panic(err)
	}
	go func() {
		time.Sleep(time.Second * 1)
		conn.Close()
	}()
	buf := make([]byte, 1024, 1024)
	for {
		n, err := tlsConn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
	}
}

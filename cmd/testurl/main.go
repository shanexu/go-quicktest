package main

import (
	"fmt"
	"net/url"
)

func main() {
	var u url.URL
	u.Host = "127.0.0.1:8080"
	u.Scheme = "http"
	u.Path = "/v1/test"
	q := u.Query()
	q.Add("hello", "world")
	u.RawQuery = q.Encode()
	fmt.Printf("%s\n", u.String())
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	header := make(http.Header)
	fmt.Println(http.CanonicalHeaderKey("__rid__"))
	fmt.Println(http.CanonicalHeaderKey("r-id"))
	header.Set("__rId__", "test")
	fmt.Println(header)
}

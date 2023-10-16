package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("hello world!")
	http.ListenAndServe(":8082", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("receive request", time.Now())
		fmt.Println("request url", request.URL)
		var timeout time.Duration
		var err error
		sleep := request.URL.Query().Get("sleep")
		if sleep != "" {
			timeout, err = time.ParseDuration(sleep)
			if err != nil {
				writer.WriteHeader(400)
				writer.Header().Add("content-type", "plain/text")
				writer.Write([]byte(fmt.Sprintf("parse timeout failed: %s", err)))
				return
			}
		}
		t := time.NewTimer(timeout)
		select {
		case <-t.C:
			result := map[string]string{
				"hello": "world",
			}
			b, _ := json.Marshal(result)
			writer.Header().Add("content-type", "application/json")
			writer.Write(b)
		case <-request.Context().Done():
			fmt.Printf("request done: %v\n", request.Context().Err())
		}
	}))
}

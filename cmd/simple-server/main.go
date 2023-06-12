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
		time.Sleep(time.Second * 10)
		result := map[string]string{
			"hello": "world",
		}
		b, _ := json.Marshal(result)
		writer.Header().Add("content-type", "application/json")
		writer.Write(b)
	}))
}

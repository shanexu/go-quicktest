package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		http.HandleFunc("/foo", func(rw http.ResponseWriter, r *http.Request) {
			s := r.URL.Query().Get("sleep")
			if s == "" {
				s = "0"
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte(err.Error()))
			}
			time.Sleep(time.Second * time.Duration(i))
			rw.Write([]byte("bar"))
		})

		http.ListenAndServe(":8080", nil)
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second)
		client := &http.Client{
			Timeout: time.Second * 5,
		}
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:8080/foo?sleep=4", nil)
		if err != nil {
			log.Println(err)
			return
		}
		rep, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		b, err := ioutil.ReadAll(rep.Body)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(b))
	}()

	wg.Wait()
}

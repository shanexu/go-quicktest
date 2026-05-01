package main

import (
	"fmt"
	"os"
)

func main() {
	buf := make([]byte, 16, 16)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
		fmt.Printf("%x\n", buf[:n])
	}
}

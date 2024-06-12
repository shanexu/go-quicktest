package main

import (
	"fmt"
)

func main() {
	str := "hello"
	defer fmt.Println(str)
	fmt.Println("main")
	str = "world"
	defer fmt.Println(str)
}

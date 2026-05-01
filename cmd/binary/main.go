package main

import "fmt"

func main() {
	var x int8 = -1
	var y uint8 = uint8(x)
	fmt.Printf("%d\n", y)
}

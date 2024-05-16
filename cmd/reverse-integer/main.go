package main

import "fmt"

func main() {
	fmt.Println(Reverse(100))
	fmt.Println(Reverse(101))
	fmt.Println(Reverse(123))
	fmt.Println(Reverse(1))
	fmt.Println(Reverse(1234))
}

func Reverse(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n = n / 10
	}
	return r
}

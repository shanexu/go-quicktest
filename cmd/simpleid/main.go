package main

import (
	"fmt"

	"github.com/shanexu/go-quicktest/internal/simpleid"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(simpleid.NextID())
	}
}

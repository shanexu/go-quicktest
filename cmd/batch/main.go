package main

import (
	"fmt"
	"strconv"
)

func main() {
	batch := 100
	var tokens []string
	for i := 0; i < 120; i++ {
		tokens = append(tokens, strconv.Itoa(i))
	}
	for i, j, l := 0, batch, len(tokens); i < l; i, j = i+batch, j+batch {
		if j > l {
			j = l
		}
		sub := tokens[i:j]
		fmt.Println(sub)
	}

}

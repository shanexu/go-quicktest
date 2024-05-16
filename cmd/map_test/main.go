package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string)
	values := m["hello"]
	values = append(values, "world")
	m["hello"] = values
	fmt.Println(values)
}

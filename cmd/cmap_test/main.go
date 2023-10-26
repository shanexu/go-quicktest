package main

import (
	"fmt"
	"github.com/lrita/cmap"
)

func main() {
	var m cmap.Map[string, string]
	m.Store("hello", "world")
	fmt.Println(m.Count(), m)

	var cm cmap.Cmap
	cm.Store("hello", "world")
	fmt.Println(cm.Count(), cm)
}

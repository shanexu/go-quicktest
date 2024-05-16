package main

import (
	"testing"
)

func FuzzDragDrop(f *testing.F) {
	f.Add(5, 4)
	f.Fuzz(func(t *testing.T, i int, s string) {
		if i > 5 {
			t.Error("i > 5")
		}
		//if i == 5 && s == "x" {
		//	t.Error("5x")
		//}
	})
}

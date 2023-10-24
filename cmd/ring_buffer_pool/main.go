package main

import (
	"fmt"

	"github.com/panjf2000/gnet/v2/pkg/pool/ringbuffer"
)

func main() {
	rb := ringbuffer.Get()
	rb.WriteString("hello")
	fmt.Println(rb.Buffered())
	fmt.Println(rb.Len())
	fmt.Println(rb.IsFull())
	fmt.Println(rb.Cap())

	for i := 0; i < 204; i++ {
		rb.WriteString("hello")
	}

	fmt.Println(rb.Buffered())
	fmt.Println(rb.Len())
	fmt.Println(rb.IsFull())
	fmt.Println(rb.Cap())
}

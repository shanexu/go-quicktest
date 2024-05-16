package main

import (
	"fmt"

	"github.com/panjf2000/gnet/v2/pkg/pool/ringbuffer"
)

func main() {
	rb := ringbuffer.Get()
	rb.WriteString("hello")
	fmt.Println("Buffered:", rb.Buffered())
	fmt.Println("Len:", rb.Len())
	fmt.Println("IsFull: ", rb.IsFull())
	fmt.Println("Cap:", rb.Cap())

	fmt.Println()

	for i := 0; i < 204; i++ {
		rb.WriteString("hello")
	}

	fmt.Println("Buffered:", rb.Buffered())
	fmt.Println("Len:", rb.Len())
	fmt.Println("IsFull:", rb.IsFull())
	fmt.Println("Cap:", rb.Cap())

	fmt.Println()

	for i := 0; i < 205; i++ {
		rb.WriteString("hello")
	}

	fmt.Println("Buffered:", rb.Buffered())
	fmt.Println("Len:", rb.Len())
	fmt.Println("IsFull:", rb.IsFull())
	fmt.Println("Cap:", rb.Cap())

	fmt.Println()

	for i := 0; i < 410; i++ {
		rb.WriteString("hello")
	}

	fmt.Println("Buffered:", rb.Buffered())
	fmt.Println("Len:", rb.Len())
	fmt.Println("IsFull:", rb.IsFull())
	fmt.Println("Cap:", rb.Cap())

	ringbuffer.Put(rb)
}

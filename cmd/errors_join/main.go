package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("hello")
	fmt.Println(errors.Join(err, errors.New("world")))
	fmt.Println(errors.Join(nil, nil) == nil)
}

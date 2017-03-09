package main

import (
	"fmt"
)

type Intset struct {
	words []uint64
}

func main() {
	var s Intset
	s.words = []uint64{1, 10, 0}
	fmt.Println(len(s.words))
	fmt.Println(s.words[1])
	var a uint8 = 255
	fmt.Println(a)
}

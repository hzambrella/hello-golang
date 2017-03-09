package main

import (
	"fmt"
)

func main() {
	var i int
	var st string
	i = 1
	st = ("1111")
	fmt.Println(i, st)
	switch st {
	case "1111":
		if i == 1 {
			fmt.Println("hz")
		} else if i == 0 {
			fmt.Println("hao")
		}
	}
}

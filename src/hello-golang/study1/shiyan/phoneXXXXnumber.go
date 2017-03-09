package main

import (
	"fmt"
	"os"
)

func main() {
	number := []string{}
	mobile := os.Args[1:]
	for _, v := range mobile {
		//		for range v {
		//	number = append(number, v[0:2]+"XXXX"+v[7:])//:15XXXX2392
		number = append(number, v[0:3]+"XXXX"+v[7:])
		//		}

	}
	fmt.Println(number)
}

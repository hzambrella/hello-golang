package main

import (
	"fmt"
)

func main() {
	f, err := haha()
	if err != nil {
		fmt.Println(err)
		fmt.Println(f)
	}
	fmt.Println(f)
}
func haha() (int, error) {
	return 0, nil
}

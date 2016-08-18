package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(os.Stderr, err)
			continue
		}
		fmt.Println(string(data))
		fmt.Println(strings.Split(string(data), "\n"))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		countline(os.Stdin, counts)
	} else {
		for _, arg := range file {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, err)
				continue
			}
			countline(f, counts)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Println(n, line)
			}
		}
	}
}
func countline(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

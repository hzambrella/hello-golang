package main

import (
	"fmt"
	"sync"
	"time"
)

var c int = 0
var sem chan int
var N int = 10
var count int = 10
var l sync.Mutex

func main() {
	sem := make(chan int, N)
	for j := 0; j < 10000000; j++ {
		if count > 0 {
			go func() {
				sem <- 0
				add()
			}()
			count--
		} else {

			<-sem
			time.Sleep(2e4)
			c--
			count++
		}
	}
}
func add() {
	c++
	fmt.Println("c:", c)
	//	fmt.Println("count:", count)
}

/*
func addself(count int) int {
	l.Lock()
	defer l.Unlock()
	count++
	return count
}

func subself(count int) int {
	l.Lock()
	defer l.Unlock()
	count--
	return count
}*/

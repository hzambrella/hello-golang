package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan bool)

	go func() {
		time.Sleep(10e9)
		timeout <- true
	}()

	go func() {
		time.Sleep(2e9)
		fmt.Println("haha ,i close channel \"ch\"")
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("recieve ch")
	case <-timeout:
		fmt.Println("Timeout")
	}

	fmt.Println("finish")
}

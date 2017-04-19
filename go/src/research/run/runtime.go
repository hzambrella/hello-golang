package runtime

import "fmt"
import "time"

var c int=0
func run() {
	ch := make(chan int, 1024)
	timeout := make(chan bool, 1)
	
	go func() {
		for{
			time.Sleep(1e9) // 等待1秒钟
			timeout <- true
		}
	}()
	
	for {
		select {
			case ch <- 0:
			case ch <- 1:
			case <-timeout:
				<-ch
				break
		}
		//
		c++
		fmt.Println(c)
		//fmt.Println("Value received:", i)
	}
}

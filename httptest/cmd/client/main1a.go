package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"runtime"
)

var quit chan int
var c int = 0

//var l sync.RWMutex

func main() {
	for n := 0; n < 10000; n++ {
		quit := make([]chan int, 100)
		for i := 0; i < 100; i++ {
			quit[i] = make(chan int)
			go modify(quit[i])

		}
		for i := 0; i < 100; i++ {
			<-quit[i]
		}
	}
	fmt.Println("test finished")
}

func modify(quit chan int) {

	host := "http://localhost:8080"
	t1 := time.Now()
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})
	t2 := time.Now()
	if err != nil {
		panic(err)
	}
	tq1 := time.Now()
	quit <- 1
	tq2 := time.Now()
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	c++
	cp := runtime.NumGoroutine()
	fmt.Println("request number:", c)
	fmt.Println("go number:", cp)
	fmt.Println(string(result))
	t3 := time.Now()
	fmt.Println("q1:", tq1.Sub(t1))
	fmt.Println("q2:", tq2.Sub(t1))
	fmt.Println("1:", t2.Sub(t1))
	fmt.Println("2:", t3.Sub(t1))
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

var quit chan int
var c int = 0

//var l sync.RWMutex

func main() {
	t1 := time.Now()
	quit := make(chan int)
	for i := 0; i < 100; i++ {
		quit = make(chan int)
		go modify(quit, t1)
	}

	for j := 0; j < 1; j++ {
		fmt.Println("len(quit)", len(quit))
		<-quit

	}
	fmt.Println("test finished")
}

func modify(quit chan int, t1 time.Time) {

	quit <- 1
	host := "http://localhost:8080"
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	t2 := time.Now()
	t3 := t2.Sub(t1)
	c++
	cp := runtime.NumGoroutine()
	fmt.Println("request number:", c)
	fmt.Println("go number:", cp)
	fmt.Println("use time:", t3)
	fmt.Println(string(result))
}

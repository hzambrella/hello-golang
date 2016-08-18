package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

var c int = 0 //统计请求次数
var quit chan int
var count int
var N int = 49

func main() {
	t1 := time.Now()
	count = 49
	quit := make(chan int, N)

	for j := 0; j < 100000000000000000; j++ {
		go func() {
			quit <- 1
			modify(t1)
		}()
	}
}

func modify(t1 time.Time) {
	c++
	fmt.Println("request number:", c)
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
	cp := runtime.NumGoroutine()
	fmt.Println("go number:", cp)
	fmt.Println("use time:", t3)
	fmt.Println(string(result))
}

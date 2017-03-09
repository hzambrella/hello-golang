package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

var c int = 0           //统计请求次数
var finish chan int     //堵住main直到完成final次请求
var Final int = 1000000 //设定总请求数
var N int = 50          //并发数
var quit chan int

func main() {
	timebegan := time.Now()
	quit = make(chan int, N)
	finish = make(chan int)
	for i := 0; i < N; i++ {
		go connect()
	}
	<-finish
	finaltime := time.Now()
	fmt.Println("total time:", finaltime.Sub(timebegan))
}

func connect() {
	for {
		if Final <= 0 {
			fmt.Println("finish")
			finish <- 1
			break
		}
		tm1 := time.Now()
		modify()
		select {
		default:
			<-quit
			Final--
			tm2 := time.Now()
			fmt.Println("connect time: ", tm2.Sub(tm1))
		}
	}
}
func modify() {
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
	c++
	cp := runtime.NumGoroutine()
	fmt.Printf("第%d次请求\n", c)
	fmt.Println("go number:", cp)
	fmt.Println(string(result))
	quit <- 1
}

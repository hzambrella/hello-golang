package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

var requestnum int = 0  //统计请求次数
var finish chan int     //堵住main直到完成final次请求
var Final int = 1000000 //设定总请求数
var N int = 50          //并发数

func main() {
	timebegan := time.Now()
	finish = make(chan int)
	for i := 0; i < N; i++ {
		go connect()
	}
	<-finish
	totaltime := time.Now()
	fmt.Println("total time:", totaltime.Sub(timebegan))
}

func connect() {
	for {
		if Final == 0 {
			fmt.Println("test finish")
			finish <- 1
			break
		}
		tm1 := time.Now()
		modify()
		Final--
		tm2 := time.Now()
		fmt.Println("connect time: ", tm2.Sub(tm1))
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
	requestnum++
	gomum := runtime.NumGoroutine()
	fmt.Printf("第%d次请求\n", requestnum)
	fmt.Println("go number:", gomum)
	fmt.Println(string(result))
}

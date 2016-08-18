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
var t1 time.Time        //程序运行开始时间
var N int = 100         //并发数
var quit chan int

func main() {
	t1 = time.Now()
	quit = make(chan int, N)
	for i := 0; i < N; i++ {
		go connect()
	}
	finish <- 1
}

func connect() {
	for {
		if Final == 0 {
			fmt.Println("finish")
			<-finish
			break
		}
		modify(t1)
		select {
		default:
			<-quit
		}
	}
}
func modify(t1 time.Time) {
	host := "http://localhost:8080"
	tm1 := time.Now()
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
	c++
	cp := runtime.NumGoroutine()
	fmt.Printf("第%d次请求\n", c)
	fmt.Println("go number:", cp)
	fmt.Println("此次连接时间:", t2.Sub(tm1))
	fmt.Println("time passed:", t2.Sub(t1))
	fmt.Println(string(result))
	quit <- 1
}

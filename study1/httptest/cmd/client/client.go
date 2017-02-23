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
var finish chan int
var Final int
var t1 time.Time

func main() {
	t1 = time.Now()
	Final = 1000000 //总循环次数

	for i := 0; i < 50; i++ {
		go line(i)
	}

	finish <- 1
	if Final == 0 {
		<-finish
	}
}

func line(i int) {
	for {

		if Final == 0 {
			fmt.Println("finish")
			break
		}
		modify(t1)
		Final -= 1
	}
}

func modify(t1 time.Time) {
	host := "http://localhost:8080"
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()
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
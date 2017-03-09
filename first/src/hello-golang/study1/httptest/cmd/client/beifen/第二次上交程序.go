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
var quit chan bool
var count int

func main() {
	t1 := time.Now()
	final := 10000000000000
	count = 49
	quit := make(chan bool)
	for {
		if final == 0 {
			fmt.Println("finish")
			break
		}
		if count > 0 {
			count -= 1
			go func() {
				modify(quit, t1)
			}()
		} else {
			<-quit
			final -= 1
			count += 1
		}
	}
}
func modify(quit chan bool, t1 time.Time) {
	host := "http://localhost:8080"
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})
	if err != nil {
		panic(err)
	}
	quit <- true
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

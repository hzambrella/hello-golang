package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var re int = 0
var quit = make([]chan int, 10)

func main() {
	time.Sleep(time.Second)
	t1 := time.Now()
	for n := 0; n < 10000000; n++ {
		for i := 0; i < 10; i++ {
			quit[i] = make(chan int)
			go modify(quit[i])
		}
		time.Sleep(time.Second)
		for j := 0; j < len(quit); j++ {
			<-quit[j]
		}
	}
	t2 := time.Now()
	t3 := t2.Sub(t1)
	fmt.Println("time used", t3)
}
func modify(quit chan int) {
	h := "http://localhost:8080"
	resp, err := http.PostForm(h+"/test/modify", url.Values{"key": {"test1"}, "data": {"aaa"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	re++
	fmt.Println(string(result))

	fmt.Println("已经请求：", re)

	quit <- 0
}

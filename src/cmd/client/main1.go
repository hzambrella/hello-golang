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
	t1 := time.Now()
	for n := 0; n < 10000; n++ {
		quit := make([]chan int, 100)
		for i := 0; i < 100; i++ {
			quit[i] = make(chan int)
			go modify(quit[i], t1)

		}

		for j, _ := range quit {
			fmt.Println("lenquit:", len(quit))
			<-quit[j]
		}
	}
	fmt.Println("test finished")
}

func modify(quit chan int, t1 time.Time) {

	host := "http://localhost:8080"
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})

	if err != nil {
		panic(err)
	}
	quit <- 1
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

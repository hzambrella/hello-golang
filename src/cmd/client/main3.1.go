package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"sync"
	"time"
)

var quit chan int
var c int = 0
var count int
var connect int = 0
var l sync.RWMutex
var dd int = 0

func main() {
	t1 := time.Now()
	count = 50
	quit := make(chan int, 50)

	for n := 0; n < 100000; n++ {

		count = 50 - connect
		for i := 0; i < count; i++ {
			go modify(quit, t1)
		}
		count := connect
		if count > 0 {
			for i := 0; i < count; i++ {
				<-quit
				connect -= 1
			}
		}
	}

	fmt.Println("test finished")
}

func modify(quit chan int, t1 time.Time) {
	connect += 1
	fmt.Println("connect:", connect)
	host := "http://localhost:8080"
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})

	if err != nil {
		panic(err)
	}
	fmt.Println("here")
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

/*func add() {
	//	l.Lock()
	//	defer l.Unlock()
	connect += 1
}*/

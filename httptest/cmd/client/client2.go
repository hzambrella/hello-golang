package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

var quit chan *http.Response
var finish chan int
var resp *http.Response
var tp *http.Response
var final int = 1000000
var c int = 0
var N int = 50
var lack int

func ab() {
	time.Sleep(3e3)
	if lack < N {
		w, t := get()
		res(w, t)
		final--
		revert()
	}
}
func main() {
	finish = make(chan int)
	lack = N
	pool()
	for i := 0; i < N && final > 0; i++ {
		go func() {
			for {
				if final < 0 {
					<-finish
					break
				}
				ab()
			}
		}()
	}
	finish <- 1
	fmt.Println("test finish")
}

//定义连接池
func pool() {
	quit = make(chan *http.Response, N)
	for i := 0; i < N; i++ {
		tp = modify()
		quit <- tp
		lack--
	}
}

//取连接
func get() (*http.Response, time.Time) {
	id := <-quit
	lack++
	t1 := time.Now()
	return id, t1
}

//放回连接
func revert() {
	x := modify()
	lack--
	quit <- x
}

//创建连接
func modify() *http.Response {
	host := "http://localhost:8080"
	//	tm1 := time.Now()
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})
	if err != nil {
		panic(err)
	}
	//	tm2 := time.Now()
	//	fmt.Println("postform  time:", tm2.Sub(tm1))
	return resp
}

//使用连接
func res(t *http.Response, t1 time.Time) {
	defer t.Body.Close()
	result, err := httputil.DumpResponse(t, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	c++
	fmt.Println("request number: ", c)
	cp := runtime.NumGoroutine()
	fmt.Println("go number:", cp)
	t2 := time.Now()
	fmt.Println("use  time of connect: ", t2.Sub(t1))
}

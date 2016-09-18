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
var N int = 90 //40  90

func ab() {
	w := get()
	res(w)
	final--
}

func main() {
	pool()
	for i := 0; i < 10; i++ {
		go func() {
			for {
				if final == 0 {
					break
				}
				ab()
			}
		}()
		time.Sleep(4e4)
	}
	finish <- 1
	if final == 0 {
		<-finish
	}
}

func pool() {
	quit = make(chan *http.Response, N)
	for i := 0; i < N; i++ {
		tp = modify()
		quit <- tp
	}
}

func get() *http.Response {
	id := <-quit
	defer release()
	return id
}

func release() {
	x := modify()
	quit <- x
}

func modify() *http.Response {
	host := "http://localhost:8080"
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})
	if err != nil {
		panic(err)
	}
	return resp
}

func res(t *http.Response) {
	defer t.Body.Close() //!!!!!
	result, err := httputil.DumpResponse(t, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	c++
	fmt.Println("request number", c)
	cp := runtime.NumGoroutine()
	fmt.Println("go number:", cp)
}

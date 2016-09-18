package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	//		"time"
)

var Quit chan *http.Response
var resp *http.Response
var tp *http.Response
var final int = 100000
var c int = 0
var N int = 49

func main() {
	pool()
	for i := 0; i < final; i++ {
		w := get()
		res(w)
	}
}

func pool() {
	Quit = make(chan *http.Response, N)
	for i := 0; i < N; i++ {
		tp = modify()
		Quit <- tp
	}
}

func get() *http.Response {
	id := <-Quit
	defer release()
	return id
}

func release() {
	x := modify()
	Quit <- x
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
	//	defer t.Body.Close()
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

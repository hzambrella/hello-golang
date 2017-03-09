package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

type tps struct {
	rs *http.Response
	t  time.Time
}

var quit chan tps
var finish chan int
var final int = 1000000
var c int = 0
var N int = 50
var lack int

func ab() {
	if lack < N {
		get()
		final--
	} else {
		time.Sleep(100e10)
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

//创建N个连接
func pool() {
	quit = make(chan tps, N)
	for i := 0; i < N; i++ {
		tp, t1 := modify()
		quit <- tps{rs: tp, t: t1}
		lack--
	}
}

//释放一个连接
func get() {
	tps2 := <-quit
	tp2 := tps2.rs
	t2 := tps2.t
	lack++
	res(tp2, t2)
	revert(tp2)
}

//补充连接
func revert(tp2 *http.Response) {
	lack--
	quit <- tps{rs: tp2, t: time.Now()}
}

//创建连接
func modify() (*http.Response, time.Time) {
	host := "http://localhost:8080"
	tm1 := time.Now()
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})
	if err != nil {
		panic(err)
	}
	//	tm2 := time.Now()
	//	fmt.Println("postform  time:", tm2.Sub(tm1))
	return resp, tm1
}

//连接信息

func res(t *http.Response, t1 time.Time) {
	result, err := httputil.DumpResponse(t, true)
	if err != nil {
		panic(err)
	}
	t2 := time.Now()
	cp := runtime.NumGoroutine()
	c++
	fmt.Println(string(result))
	fmt.Println("request number: ", c)
	fmt.Println("go number:", cp)
	fmt.Println("use  time of connect: ", t2.Sub(t1))
	fmt.Println("lack:", lack)
}

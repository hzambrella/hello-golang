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
	final := 10000000000000 //总循环次数
	count = 50              //信号量。初始值为50。为100的话检测EST-C连接数为200左右。
	quit := make(chan bool)
	for {
		if final == 0 {
			fmt.Println("finish")
			break
		}
		if count > 0 { //第一次for循环开始，count大于0时，会放入1个goroutine。
			count -= 1 //放入一个，count减1。count小于零时就不放入了。
			go func() {
				modify(quit, t1)
			}()
		} else { //如果count为0，并发池满了。释放一个信道。
			<-quit
			final -= 1
			count += 1 //count+1。然后count可能又大于0了，下一个for循环加一个go程。
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

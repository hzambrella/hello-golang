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

var c int = 0

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// 创建并发控制池, 设置并发数量与总数量
func NewPool(cap, total int) *Pool {
	p := &Pool{
		queue: make(chan int, cap),
		wg:    new(sync.WaitGroup),
	}
	p.wg.Add(total)
	return p
}

// 向并发队列中添加一个
func (p *Pool) AddOne() {
	p.queue <- 1
}

// 并发队列中释放一个, 并从总数量中减去一个
func (p *Pool) DelOne() {
	<-p.queue
	p.wg.Done()
}
func main() {
	t1 := time.Now()
	pool := NewPool(20, 10000) // 初始化一个容量为20的并发控制池

	go func() {

		pool.AddOne() // 向并发控制池中添加一个, 一旦池满则此处阻塞
		modify(t1)

		pool.DelOne() // 从并发控制池中释放一个, 之后其他被阻塞的可以进入池中
	}()

	pool.wg.Wait() // 等待所有下载全部完成
}
func modify(t1 time.Time) {

	host := "http://localhost:8080"
	resp, err := http.PostForm(host+"/test/modify", url.Values{"key": {"test1"}, "data": {"---10×10---"}})

	if err != nil {
		panic(err)
	}
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
	//	fmt.Println("modify count", count)

}

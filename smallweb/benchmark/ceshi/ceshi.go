package ceshi

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var totalall int = 10000
var total int = 100
var mut sync.Mutex

type tongji struct {
	nowNum int
	mu     sync.Mutex
}

func (a *tongji) Count() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.nowNum = a.nowNum + 1
	fmt.Println(a.nowNum)
}

func (a tongji) Get() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.nowNum
}

func ceshi() {
	// 一定要make!!!
	close := make(chan string)
	tj := &tongji{nowNum: 0, mu: mut}
	for i := 0; i < 100; i++ {
		index := i
		fmt.Println(index, " in it")
		go func() {
			for {
				//				time.Sleep(1e9)
				fmt.Println(index, ":do")
				if tj.Get() >= totalall {
					fmt.Println(index, ":end go")
					close <- "end"
					break
				}
				err := BenchSmallWebPost()
				if err != nil {
					close <- err.Error()
					break
				}
				tj.Count()
			}
		}()
	}

	mess := <-close
	fmt.Println(mess)
}

var HostIp = GetIp()

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		ip := strings.Split(addr.String(), "/")[0]
		code := strings.Split(ip, ".")
		switch code[0] {
		case "10", "127":
			continue
		default:
			return ip
		}
	}
	panic(addrs)
}

func BenchSmallWebGet() error {
	link := "http://" + HostIp + ":31111/test/html2"
	fmt.Println(link)
	_, err := http.Get(link)
	return err
}

func BenchSmallWebPost() error {
	data := make(url.Values, 0)
	data.Set("name", "123")
	data.Set("status", "4")
	link := "http://" + HostIp + ":31111/test/post"
	fmt.Println(link)
	_, err := http.PostForm(link, data)
	return err
}

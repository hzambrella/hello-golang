package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/kr/beanstalk"
)

var (
	TubeName1 string = "channel1"
	TubeName2 string = "channel2"
)

func Producer(fname, tubeName string) {
	if fname == "" || tubeName == "" {
		return
	}

	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true
	fmt.Println(fname, " [Producer] tubeName:", tubeName, " c.Tube.Name:", c.Tube.Name)

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("for %s %d", tubeName, i)
		c.Put([]byte(msg), 30, 0, 120*time.Second)
		fmt.Println(fname, " [Producer] beanstalk put body:", msg)
		//time.Sleep(1 * time.Second)
	}

	c.Close()
	fmt.Println("Producer() end.")
}

func Consumer(fname, tubeName string) {
	if fname == "" || tubeName == "" {
		return
	}

	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true

	fmt.Println(fname, " [Consumer] tubeName:", tubeName, " c.Tube.Name:", c.Tube.Name)

	substr := "timeout"
	for {
		fmt.Println(fname, " [Consumer]///////////////////////// ")
		//从队列中取出
		id, body, err := c.Reserve(1 * time.Second)
		if err != nil {
			if !strings.Contains(err.Error(), substr) {
				fmt.Println(fname, " [Consumer] [", c.Tube.Name, "] err:", err, " id:", id)
			}
			continue
		}
		fmt.Println(fname, " [Consumer] [", c.Tube.Name, "] job:", id, " body:", string(body))

		//从队列中清掉
		err = c.Delete(id)
		if err != nil {
			fmt.Println(fname, " [Consumer] [", c.Tube.Name, "] Delete err:", err, " id:", id)
		} else {
			fmt.Println(fname, " [Consumer] [", c.Tube.Name, "] Successfully deleted. id:", id)
		}
		fmt.Println(fname, " [Consumer]/////////////////////////")
		//time.Sleep(1 * time.Second)
	}
	fmt.Println("Consumer() end. ")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go Producer("PA", TubeName1)
	go Producer("PB", TubeName2)

	go Consumer("CA", TubeName1)
	go Consumer("CB", TubeName2)

	time.Sleep(10 * time.Second)
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)
var(
	// server.go
	/*
	 setHost string= "http://localhost"
	 setPort string=":8080"
	 setPath string="/test/modify"
	 urlValues url.Values=url.Values{"key": {"test1"}, "data": {"哈哈。。。"}}
	 */
	 //n5d
	//192.168.0.128:9600/api/orderback/orshop/1/goods/105
	//setHost string= "http://192.168.0.128"
	 //setPort string=":9600"
//	 setPath string="/api/orderback/orshop/1/goods/105"
	// urlValues url.Values=url.Values{"price":{"1"}}

	// n5d
	setHost string= "http://192.168.0.128"
	 setPort string=":9600"
	setPath string="/auth/login"
	//setHost string= "http://192.168.0.128"
	 AllNum int=1000
	 thread int =100
)

type TestPt interface{
	TestPt(u url.Values)(time.Time)
	TestPtGet()time.Time
}

type TestParams struct{
	Host string
	Port string
	Path string
}

func main(){
	set:=TestParams{Host:setHost,Port:setPort,Path:setPath}
	//fmt.Println(fmt.Sprintf("post:%surl\n data:%s\n",murl,u))
	t1 := time.Now()
	all:=make(chan int,0)
	allNum:=0
	for i:=0;i<thread;i++{
		mtChan:=make(chan time.Time,0)

		m:=&ModifyTest{set,mtChan}
		go func(){
			for {
				//tm1:=m.TestPtPost(urlValues)
				tm1:=m.TestPtGet()
				tm2:=time.Now()
				tm2Subtm1:=tm2.Sub(tm1)
				fmt.Println("time use:",tm2Subtm1)
				allNum++
			}
		}()

		if allNum>AllNum{
			all<-allNum
		}
	}

	allallNum:=<-all
	fmt.Println("all req:",allallNum)
	t2 := time.Now()
	d := t2.Sub(t1)
	fmt.Println(d)
}

type ModifyTest struct{
	TestParams
	connTime chan time.Time
}

func (m *ModifyTest)TestPtPost(u url.Values)time.Time{
	murl:=m.Host+m.Port+m.Path
	tm1:=time.Now()

	//fmt.Println(fmt.Sprintf("post:%surl\n data:%s\n",murl,u))
	resp, err := http.PostForm(murl,u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	return tm1
}

func (m *ModifyTest)TestPtGet()time.Time{
	murl:=m.Host+m.Port+m.Path
	tm1:=time.Now()

	//fmt.Println(fmt.Sprintf("post:%surl\n data:%s\n",murl,u))
	resp, err := http.Get(murl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
//		panic(err)
	}
	if err==nil{
	fmt.Println(string(result))
 }
	return tm1
}

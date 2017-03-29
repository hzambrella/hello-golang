package main

import (
	"fmt"
	"net"
	"net/http"
	"http/template"
	"log"
)

var port="8080"
func main(){
	// handle
	http.HandleFunc("/test",test)

	// ip
	ip,err:=getIp()
	if err!=nil{
		log.Println(err.Error())
	}

	log.Println("listen and serve at"+port+" "+ip)
	http.ListenAndServe(port,nil)
}


func getIp()(string,error){
	addrs,err:=net.InterfaceAddrs()
	if err!=nil{
		log.Println(addrs)
		return "",err
	}

}

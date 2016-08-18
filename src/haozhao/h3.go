package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	url, _ := url.Parse(os.Args[1])
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url.String(), nil)
	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println("------------------------------------------- ")
	fmt.Println(string(dump))
	fmt.Println("------------------------------------------- ")
	response, _ := client.Do(request)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	var buf [512]byte
	reader := response.Body
	for {
		n, _ := reader.Read(buf[0:])
		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)
}

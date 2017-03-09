package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	url := os.Args[1]
	response, _ := http.Get(url)
	b, _ := httputil.DumpResponse(response, false)
	fmt.Println("-------------------------1------------------")
	fmt.Print(string(b))
	fmt.Println("-------------------------1------------------")
	contentTypes := response.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		fmt.Println("-------------------------2------------------")
		fmt.Println("Cannot handle", contentTypes)
		fmt.Println("-------------------------2------------------")
		os.Exit(4)
	}
	var buf [512]byte
	reader := response.Body
	for {
		n, _ := reader.Read(buf[0:])
		fmt.Println("-------------------------3------------------")
		fmt.Print(string(buf[0:n]))
		fmt.Println("-------------------------3------------------")
	}
	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	for _, cType := range contentTypes {
		if strings.Index(cType, "UTF-8") != -1 {
			return true
		}
	}
	return false
}

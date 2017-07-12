package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://localhost:9200/lazy")
	if err != nil {
		fmt.Println(err.Error())
	}
	b, err := httputil.DumpResponse(resp, false)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(b))
}

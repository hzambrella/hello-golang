// golang获取完整的url
/*查阅了r *http.Request对象中的所有属性，没有发现可以直接获取完整的url的方法。于是尝试根据host和请求地址进行拼接。在golang中可以通过r.Host获取hostname，r.RequestURI获取相应的请求地址。

但是还少一个协议的判断，怎么区分是http和https呢？一开始尝试通过r.Proto属性判断，但是发现该属性不管是http，还是https都是返回HTTP/1.1，又寻找了下,发现TLS属性，在https协议下有对应值，在http下为nil。*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	scheme :="http://"
	fmt.Println(r.TLS)
	if r.TLS != nil {
		scheme= "https://"
	}
	fmt.Println(strings.Join([]string{scheme, r.Host, r.RequestURI},""))
}

func main() {
	http.HandleFunc("/index", index)
	log.Println("start at:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
